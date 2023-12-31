package task

import (
	"github.com/Implex-ltd/hcsolver/internal/hcaptcha"
)

const (
	STATUS_SOLVING = 0
	STATUS_SOLVED  = 1
	STATUS_ERROR   = 2
)

type TaskStatus struct {
	Status int
	Token  string
}

type HcaptchaTask struct {
	Config  *hcaptcha.Config
	Captcha *hcaptcha.Hcap
	Status  TaskStatus
	ID      string
}

type BodyNewSolveTask struct {
	Domain          string `json:"domain"`
	SiteKey         string `json:"site_key"`
	UserAgent       string `json:"user_agent"`
	Proxy           string `json:"proxy"`
	TaskType        int    `json:"task_type"`
	Invisible       bool   `json:"invisible"`
	Rqdata          string `json:"rqdata"`
	FreeTextEntry   bool   `json:"a11y_tfe"`
	Turbo           bool   `json:"turbo"`
	TurboSt         int    `json:"turbo_st"`
	HcAccessibility string `json:"hc_accessibility"`
	OneclickOnly    bool   `json:"oneclick_only"`

	// motion data customisation
	Href string `json:"href"`
	Exec bool   `json:"exec"`
	Dr   string `json:"dr"`
}
