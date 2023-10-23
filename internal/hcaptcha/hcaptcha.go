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
	"github.com/Implex-ltd/hcsolver/internal/utils"

)

const (
	LANG = "fr"
)

func NewHcaptcha(config *Config) (*Hcap, error) {
	builder, err := fingerprint.NewFingerprintBuilder(config.UserAgent, config.Href)
	if err != nil {
		return nil, err
	}

	if _, err := builder.GenerateProfile(); err != nil {
		return nil, err
	}

	config.UserAgent = builder.Manager.Fingerprint.Browser.UserAgent

	fp, err := ApplyFingerprint(config)
	if err != nil {
		return nil, err
	}

	// edit fp to match scrapped one
	fp.Navigator.Languages = builder.Profile.Navigator.Languages
	fp.Navigator.Language = builder.Profile.Navigator.Language
	fp.Navigator.Platform = builder.Profile.Navigator.Platform

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

	hc := &Hcap{
		Fingerprint: fp,
		Config:      config,
		Http:        c,
		Logger:      config.Logger,
		Manager:     builder,
		Sessions:    [][]string{},
	}
	
	if _, err := hc.Manager.GenerateProfile(); err != nil {
		return nil, fmt.Errorf("cant generate fingerprint profile")
	}

	return hc, nil
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
	fp.Navigator.Platform = infos.OSName

	return fp, nil
}

func (c *Hcap) CheckSiteConfig() (*SiteConfig, error) {
	st := time.Now()

	body, err := c.Http.Do(cleanhttp.RequestOption{
		Method: "POST",
		Url:    fmt.Sprintf("https://%shcaptcha.com/checksiteconfig?v=%s&host=%s&sitekey=%s&sc=1&swa=1&spst=1", utils.RandomElementString([]string{"api2.", ""}), fingerprint.VERSION, c.Config.Domain, c.Config.SiteKey),
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

func (c *Hcap) GetChallenge(config *SiteConfig) (*Captcha, error) {
	var pow string
	var err error

	st := time.Now()

	pow, err = c.GetHsw(config.C.Req)
	if err != nil {
		return nil, err
	}

	c.AnswerProcessing = time.Since(st)

	pdc, _ := json.Marshal(&Pdc{
		S:   st.UTC().UnixNano() / 1e6,
		N:   0,
		P:   0,
		Gcs: utils.RandomNumber(40, 110), //int(time.Since(st).Milliseconds()),
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
		`v`:          fingerprint.VERSION,
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

	/*if c.Config.FreeTextEntry {
		payload.Set(`a11y_tfe`, `true`)
	}*/

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
		return nil, fmt.Errorf("ip is ratelimited or site-key is geo-restricted, please switch your proxy")
	}

	var captcha Captcha
	if err := json.Unmarshal(body.Body(), &captcha); err != nil {
		return nil, err
	}

	return &captcha, nil
}

func (c *Hcap) GetChallengeFreeTextEntry(config *SiteConfig) (*Captcha, error) {
	var pow string
	var err error

	chall, err := c.GetChallenge(config)
	if err != nil {
		return nil, err
	}

	c.Sessions = append(c.Sessions, []string{
		chall.Key,
		c.WidgetIDList[0],
	})

	st := time.Now()

	pow, err = c.GetHsw(chall.C.Req)
	if err != nil {
		return nil, err
	}

	c.AnswerProcessing = time.Since(st)

	pdc, _ := json.Marshal(&Pdc{
		S:   st.UTC().UnixNano() / 1e6,
		N:   1,
		P:   1,
		Gcs: utils.RandomNumber(40, 110), //int(time.Since(st).Milliseconds()),
	})

	motion := c.NewMotionData(&Motion{
		IsCheck: false,
	})

	C, _ := json.Marshal(&C{
		Type: chall.C.Type,
		Req:  chall.C.Req,
	})

	TextChall, _ := json.Marshal(&chall)

	payload := url.Values{}
	for name, value := range map[string]string{
		`v`:          fingerprint.VERSION,
		`sitekey`:    c.Config.SiteKey,
		`host`:       c.Config.Domain,
		`hl`:         LANG,
		`action`:     `challenge-refresh`,
		`extraData`:  string(TextChall),
		`motionData`: motion,
		`pdc`:        string(pdc),
		`old_ekey`:   chall.Key,
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
		return nil, fmt.Errorf("ip is ratelimited or site-key is geo-restricted, please switch your proxy")
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
		V:            fingerprint.VERSION,
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
