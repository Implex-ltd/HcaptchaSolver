package task

import (
	"fmt"
	"time"

	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/config"
	"github.com/Implex-ltd/hcsolver/internal/hcaptcha"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
		zap.String("ua", T.Config.UserAgent),
		zap.String("key", T.Config.SiteKey),
		zap.String("dns", T.Config.Domain),
		zap.String("proxy", T.Config.Proxy),
		zap.String("rqdata", T.Config.Rqdata),
		zap.Bool("inv", T.Config.Invisible),
		zap.Bool("a11y", T.Config.FreeTextEntry),
		zap.Int("type", T.Config.TaskType),
		zap.Bool("turbo", T.Config.Turbo),
		zap.Int("st", T.Config.TurboSt),
		zap.String("hc_accessibility", T.Config.HcAccessibility),
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
		config.Logger.Info("solved (oneclick)",
			zap.Object("perf", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
				enc.AddInt64("hsw", T.Captcha.HswProcessing.Milliseconds())
				enc.AddInt64("recognition", T.Captcha.AnswerProcessing.Milliseconds())
				enc.AddInt64("task", time.Since(st).Milliseconds())
				enc.AddInt64("cc", T.Captcha.CheckProcessing.Milliseconds())
				enc.AddInt64("gc", T.Captcha.GetProcessing.Milliseconds())
				enc.AddInt64("gs", T.Captcha.SiteConfigProcessing.Milliseconds())
				return nil
			})),
			zap.Object("config", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
				enc.AddString("rqdata", T.Config.Rqdata)
				enc.AddBool("inv", T.Config.Invisible)
				enc.AddBool("a11y", T.Config.FreeTextEntry)
				enc.AddInt("type", T.Config.TaskType)
				enc.AddString("ua", T.Config.UserAgent)
				enc.AddString("key", T.Config.SiteKey)
				enc.AddString("dns", T.Config.Domain)
				enc.AddString("proxy", T.Config.Proxy)
				return nil
			})),
			zap.Object("captcha", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
				enc.AddString("key", captcha.GeneratedPassUUID)
				enc.AddString("prompt", captcha.RequesterQuestion.En)
				enc.AddString("type", captcha.RequestType)
				return nil
			})),
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
		zap.Object("perf", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
			enc.AddInt64("hsw", T.Captcha.HswProcessing.Milliseconds())
			enc.AddInt64("recognition", T.Captcha.AnswerProcessing.Milliseconds())
			enc.AddInt64("task", time.Since(st).Milliseconds())
			enc.AddInt64("cc", T.Captcha.CheckProcessing.Milliseconds())
			enc.AddInt64("gc", T.Captcha.GetProcessing.Milliseconds())
			enc.AddInt64("gs", T.Captcha.SiteConfigProcessing.Milliseconds())
			return nil
		})),
		zap.Object("config", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
			enc.AddString("rqdata", T.Config.Rqdata)
			enc.AddBool("inv", T.Config.Invisible)
			enc.AddBool("a11y", T.Config.FreeTextEntry)
			enc.AddInt("type", T.Config.TaskType)
			enc.AddString("ua", T.Config.UserAgent)
			enc.AddString("key", T.Config.SiteKey)
			enc.AddString("dns", T.Config.Domain)
			enc.AddString("proxy", T.Config.Proxy)
			return nil
		})),
		zap.Object("captcha", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
			enc.AddString("key", response.GeneratedPassUUID)
			enc.AddString("prompt", captcha.RequesterQuestion.En)
			enc.AddString("type", captcha.RequestType)
			return nil
		})),
	)

	return response, nil
}
