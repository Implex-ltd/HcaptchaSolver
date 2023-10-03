package hcaptcha

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/Implex-ltd/cleanhttp/cleanhttp"
	"github.com/Implex-ltd/fingerprint-client/fpclient"
	"github.com/Implex-ltd/hcsolver/internal/hcaptcha/fingerprint"
)

const (
	VERSION = "1b812e2"
	LANG    = "fr"
)

func NewHcaptcha(config *Config) (*Hcap, error) {
	fp, err := ApplyFingerprint(config)
	if err != nil {
		return nil, err
	}

	c, err := cleanhttp.NewFastCleanHttpClient(&cleanhttp.Config{
		Proxy:               config.Proxy,
		BrowserFp:           fp,
		Timeout:             15,
		ReadTimeout:         time.Second * 15,
		WriteTimeout:        time.Second * 15,
		MaxIdleConnDuration: time.Minute,
	})
	if err != nil {
		return nil, err
	}

	builder, err := fingerprint.NewFingerprintBuilder(config.UserAgent)
	if err != nil {
		return nil, err
	}

	if _, err := builder.GenerateProfile(); err != nil {
		return nil, err
	}

	return &Hcap{
		Fingerprint:          fp,
		Config:               config,
		Http:                 c,
		Logger:               config.Logger,
		ChallengeFingerprint: builder,
	}, nil
}

func ApplyFingerprint(config *Config) (*fpclient.Fingerprint, error) {
	fp, err := fpclient.LoadFingerprint(&fpclient.LoadingConfig{
		FilePath: "../../assets/chrome.json",
	})

	if err != nil {
		return nil, err
	}

	infos := cleanhttp.ParseUserAgent(config.UserAgent)

	fp.Navigator.UserAgent = config.UserAgent
	fp.Navigator.AppVersion = strings.Split(config.UserAgent, "Mozilla/")[1] // can crash
	fp.Navigator.Platform = infos.OSName

	// get ipinfos
	/*if config.Proxy != "" {
		infos, err := utils.Lookup(Address)

	}*/

	return fp, nil
}

func (c *Hcap) CheckSiteConfig(hsw bool) (*SiteConfig, error) {
	st := time.Now()
	swa := "1"
	if hsw {
		swa = "0"
	}

	body, err := c.Http.Do(cleanhttp.RequestOption{
		Method: "POST",
		Url:    fmt.Sprintf("https://hcaptcha.com/checksiteconfig?v=%s&host=%s&sitekey=%s&sc=1&swa=%s&spst=1", VERSION, c.Config.Domain, c.Config.SiteKey, swa),
		Header: c.HeaderCheckSiteConfig(),
	})
	c.SiteConfigProcessing = time.Since(st)
	if err != nil {
		return nil, err
	}

	var config SiteConfig
	if err := json.Unmarshal(body.Body(), &config); err != nil {
		return nil, err
	}

	if !config.Pass {
		return &config, fmt.Errorf("checksiteconfig pass is false: %v", config)
	}

	if !config.Features.A11YChallenge && c.Config.FreeTextEntry {
		return &config, fmt.Errorf("a11y_challenge is disabled on this website")
	}

	return &config, nil
}

func (c *Hcap) GetChallenge(config *SiteConfig, hsj bool) (*Captcha, error) {
	var pow string
	var err error

	st := time.Now()

	pow = "fail"
	if !hsj {
		pow, err = c.GetHsw(config.C.Req)
		if err != nil {
			return nil, err
		}
	}
	c.AnswerProcessing = time.Since(st)

	pdc, _ := json.Marshal(&Pdc{
		S:   st.UTC().UnixNano() / 1e6,
		N:   0,
		P:   0,
		Gcs: int(time.Since(st).Milliseconds()),
	})

	motion := c.NewMotionData(&Motion{
		IsCheck: false,
	})

	C, _ := json.Marshal(&C{
		Type: config.C.Type,
		Req:  config.C.Req,
	})

	payload := url.Values{}
	for name, value := range map[string]string{
		`v`:          VERSION,
		`sitekey`:    c.Config.SiteKey,
		`host`:       c.Config.Domain,
		`hl`:         LANG,
		`motionData`: motion,
		`pdc`:        string(pdc),
		`n`:          pow,
		`c`:          string(C),
		`pst`:        `false`,
	} {
		payload.Set(name, value)
	}

	if c.Config.Rqdata != "" {
		payload.Set("rqdata", c.Config.Rqdata)
	}

	if c.Config.FreeTextEntry {
		payload.Set(`a11y_tfe`, `true`)
	}

	header := c.HeaderGetCaptcha()
	if c.Config.HcAccessibility != "" {
		header.Add("cookie", fmt.Sprintf("hc_accessibility=%s", c.Config.HcAccessibility))
	}

	t := time.Now()
	body, err := c.Http.Do(cleanhttp.RequestOption{
		Method: "POST",
		Url:    fmt.Sprintf("https://hcaptcha.com/getcaptcha/%s", c.Config.SiteKey),
		Body:   strings.NewReader(payload.Encode()),
		Header: header,
	})
	c.GetProcessing = time.Since(t)
	if err != nil {
		return nil, err
	}

	if body == nil {
		return nil, fmt.Errorf("GetChallenge body is nil")
	}

	if body.StatusCode() == 429 {
		return nil, fmt.Errorf("ip is ratelimited")
	}

	var captcha Captcha
	if err := json.Unmarshal(body.Body(), &captcha); err != nil {
		return nil, err
	}

	return &captcha, nil
}

func (c *Hcap) CheckCaptcha(captcha *Captcha) (*ResponseCheckCaptcha, error) {
	var answers any
	var payload []byte
	var pow string
	var err error

	resultChans := make(chan error)
	st := time.Now()

	go func() {
		answers, err = c.SolveImages(captcha)
		if err != nil {
			resultChans <- err
			return
		}

		c.AnswerProcessing = time.Since(st)
		resultChans <- nil
	}()

	go func() {
		pow, err = c.GetHsw(captcha.C.Req)
		if err != nil {
			resultChans <- err
			return
		}

		c.HswProcessing = time.Since(st)
		resultChans <- nil
	}()

	for i := 0; i < 2; i++ {
		err := <-resultChans
		if err != nil {
			return nil, err
		}
	}

	motion := c.NewMotionData(&Motion{
		IsCheck: true,
		Answers: map[string]string{"x": "true", "y": "true", "z": "true"},
	})

	C, _ := json.Marshal(&C{
		Type: captcha.C.Type,
		Req:  captcha.C.Req,
	})

	payload, err = json.Marshal(&PayloadCheckChallenge{
		V:            VERSION,
		Sitekey:      c.Config.SiteKey,
		Serverdomain: c.Config.Domain,
		JobMode:      captcha.RequestType,
		MotionData:   motion,
		N:            pow,
		C:            string(C),
		Answers:      answers,
	})

	if err != nil {
		return nil, err
	}

	time.Sleep((time.Millisecond * time.Duration(c.Config.TurboSt)) - time.Since(st))

	t := time.Now()
	body, err := c.Http.Do(cleanhttp.RequestOption{
		Url:    fmt.Sprintf("https://hcaptcha.com/checkcaptcha/%s/%s", c.Config.SiteKey, captcha.Key),
		Body:   strings.NewReader(string(payload)),
		Method: "POST",
		Header: c.HeaderCheckCaptcha(),
	})
	c.CheckProcessing = time.Since(t)

	if err != nil {
		return nil, err
	}

	var Resp ResponseCheckCaptcha
	if json.Unmarshal([]byte(body.Body()), &Resp) != nil {
		return nil, fmt.Errorf("checkCaptcha-unmarshal: %+v", err)
	}

	if !Resp.Pass {
		return nil, fmt.Errorf("checkCaptcha: failed to pass: %+v", string(body.Body()))
	}

	return &Resp, nil
}
