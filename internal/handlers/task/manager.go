package task

import (
	"fmt"
	"time"

	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/config"
	"github.com/Implex-ltd/hcsolver/internal/hcaptcha"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func Newtask(config *hcaptcha.Config) (*HcaptchaTask, error) {
	T := HcaptchaTask{
		Status: TaskStatus{
			Status: STATUS_SOLVING,
		},
		ID:     uuid.NewString(),
		Config: config,
	}

	return &T, nil
}

func (T *HcaptchaTask) Create() error {
	hc, err := hcaptcha.NewHcaptcha(T.Config)
	if err != nil {
		return err
	}

	config.Logger.Info("new task",
		zap.String("useragent", T.Config.UserAgent),
		zap.String("sitekey", T.Config.SiteKey),
		zap.String("domain", T.Config.Domain),
		zap.String("proxy", T.Config.Proxy),
	)

	T.Captcha = hc
	return nil
}

func (T *HcaptchaTask) Solve() (*hcaptcha.ResponseCheckCaptcha, error) {
	st := time.Now()

	site, err := T.Captcha.CheckSiteConfig()
	if err != nil {
		return nil, err
	}

	if !site.Pass {
		return nil, fmt.Errorf("checksiteconfig wont pass")
	}

	fmt.Println(site)

	captcha, err := T.Captcha.GetChallenge(site)
	if err != nil {
		return nil, err
	}

	config.Logger.Info("GetChallenge",
		zap.String("req", captcha.C.Req),
		zap.String("type", captcha.C.Type),
		zap.String("prompt", captcha.RequesterQuestion.En),
	)

	if captcha.RequestType != "image_label_binary" {
		return nil, fmt.Errorf("invalid request-type: %s", captcha.RequestType)
	}

	response, err := T.Captcha.CheckCaptcha(captcha)
	if err != nil {
		return nil, err
	}

	config.Logger.Info("CheckCaptcha",
		zap.Bool("pass", response.Pass),
		zap.String("token", response.GeneratedPassUUID),
		zap.String("req", response.C.Req),
		zap.String("type", response.C.Type),
		zap.Int("expiration", response.Expiration),
	)

	if !response.Pass {
		return nil, fmt.Errorf("captcha wont pass")
	}

	config.Logger.Info("solved",
		zap.String("key", response.GeneratedPassUUID),
		zap.String("prompt", captcha.RequesterQuestion.En),
		zap.Int64("hsw_process", T.Captcha.HswProcessing.Milliseconds()),
		zap.Int64("img_process", T.Captcha.AnswerProcessing.Milliseconds()),
		zap.Int64("task_process", time.Since(st).Milliseconds()),

		zap.String("type", captcha.RequestType),

		zap.String("useragent", T.Config.UserAgent),
		zap.String("sitekey", T.Config.SiteKey),
		zap.String("domain", T.Config.Domain),
		zap.String("proxy", T.Config.Proxy),
	)

	return response, nil
}
