package main

type BodyNewSolveTask struct {
	Domain    string `json:"domain"`
	SiteKey   string `json:"site_key"`
	UserAgent string `json:"user_agent"`
	Proxy     string `json:"proxy"`
}

type TaskResponse struct {
	Success bool `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	Task    Task    `json:"task"`
	Metrics Metrics `json:"metrics"`
	Token   Token   `json:"token"`
	Err     Err     `json:"err"`
}

type Err struct {
	Retry  int64    `json:"retry"`
	Errors []string `json:"errors"`
}

type Metrics struct {
	StartTime   int64 `json:"start_time"`
	HswProcess  int64 `json:"hsw_process"`
	ImgProcess  int64 `json:"img_process"`
	TTLProcess  int64 `json:"ttl_process"`
	TaskProcess int64 `json:"task_process"`
}

type Task struct {
	TaskType   string `json:"task_type"`
	TaskPrompt string `json:"task_prompt"`
}

type Token struct {
	CAPTCHAKey string `json:"captcha_key"`
	Expiration int64  `json:"expiration"`
}
