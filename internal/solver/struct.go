package solver

import (
	"sync"
	"time"
)

var (
	pm            = sync.Mutex{}
	hashlistMutex sync.RWMutex
)

type Base64JSON struct {
	Images  map[string]string `json:"images"`
	Target  string            `json:"target"`
	Method  string            `json:"method"`
	Sitekey string            `json:"sitekey"`
	Site    string            `json:"site"`
	Ln      string            `json:"ln"`
}

type NoCapAnswer struct {
	Answer         []any  `json:"answer"`
	ID             string `json:"id"`
	Message        string `json:"message"`
	ProcessingTime string `json:"processing_time"`
	Solution       []int  `json:"solution"`
	Status         string `json:"status"`
	Target         string `json:"target"`
	URL            string `json:"url"`
}

var (
	Hashlist = map[string][]string{}
	mu       sync.Mutex
	Config   = Cfg{}
)

type Result struct {
	Hash  string
	Match bool
	Err   error
	St    time.Duration
	Url   string
	Key   string
}

type TaskList struct {
	DatapointURI string `json:"datapoint_uri"`
	TaskKey      string `json:"task_key"`
}

type BodyNewSolveTask struct {
	TaskType string     `json:"task_type"`
	Question string     `json:"question"`
	TaskList []TaskList `json:"tasklist"`
}

type SolveRepsonse struct {
	Success bool              `json:"success"`
	Data    map[string]string `json:"data"`
}

type Cfg struct {
	Server struct {
		Port int `toml:"port"`
	} `toml:"server"`
	Login struct {
		Output  string `toml:"output"`
		Enabled bool   `toml:"enabled"`
	} `toml:"login"`
}
