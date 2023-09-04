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
		zap.String("rqdata", T.Config.Rqdata),
		zap.Bool("invisible", T.Config.Invisible),
		zap.Int("task_type", T.Config.TaskType),
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
	captcha, err := T.Captcha.GetChallenge(site)
	if err != nil {
		return nil, err
	}

	if captcha.GeneratedPassUUID != "" {
		config.Logger.Info("solved (one-click)",
			zap.String("key", captcha.GeneratedPassUUID),
			zap.String("prompt", captcha.RequesterQuestion.En),
			zap.Int64("hsw_process", T.Captcha.HswProcessing.Milliseconds()),
			zap.Int64("img_process", T.Captcha.AnswerProcessing.Milliseconds()),
			zap.Int64("task_process", time.Since(st).Milliseconds()),
			zap.Int64("get_process", T.Captcha.GetProcessing.Milliseconds()),
			zap.Int64("siteconfig_process", T.Captcha.SiteConfigProcessing.Milliseconds()),

			zap.String("prompt_type", captcha.RequestType),
			zap.String("rqdata", T.Config.Rqdata),
			zap.Bool("invisible", T.Config.Invisible),
			zap.Int("task_type", T.Config.TaskType),

			zap.String("useragent", T.Config.UserAgent),
			zap.String("sitekey", T.Config.SiteKey),
			zap.String("domain", T.Config.Domain),
			zap.String("proxy", T.Config.Proxy),
		)

		return &hcaptcha.ResponseCheckCaptcha{
			C:                 captcha.C,
			Pass:              captcha.Pass,
			GeneratedPassUUID: captcha.GeneratedPassUUID,
			Expiration:        captcha.Expiration,
		}, nil
	}

	response, err := T.Captcha.CheckCaptcha(captcha)
	if err != nil {
		return nil, err
	}

	if !response.Pass {
		return nil, fmt.Errorf("captcha wont pass")
	}

	config.Logger.Info("solved",
		zap.String("key", response.GeneratedPassUUID),
		zap.String("prompt", captcha.RequesterQuestion.En),
		zap.Int64("hsw_process", T.Captcha.HswProcessing.Milliseconds()),
		zap.Int64("img_process", T.Captcha.AnswerProcessing.Milliseconds()),
		zap.Int64("task_process", time.Since(st).Milliseconds()),
		zap.Int64("check_process", T.Captcha.CheckProcessing.Milliseconds()),
		zap.Int64("get_process", T.Captcha.GetProcessing.Milliseconds()),
		zap.Int64("siteconfig_process", T.Captcha.SiteConfigProcessing.Milliseconds()),

		zap.String("prompt_type", captcha.RequestType),
		zap.String("rqdata", T.Config.Rqdata),
		zap.Bool("invisible", T.Config.Invisible),
		zap.Int("task_type", T.Config.TaskType),

		zap.String("useragent", T.Config.UserAgent),
		zap.String("sitekey", T.Config.SiteKey),
		zap.String("domain", T.Config.Domain),
		zap.String("proxy", T.Config.Proxy),
	)

	return response, nil
}
