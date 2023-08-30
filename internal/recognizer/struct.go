package recognizer

import (
	"sync"

	"github.com/zenthangplus/goccm"
)

var (
	Hashlist   = map[string][]string{}
	Selectlist = map[string][]HashData{}
	HMut       sync.RWMutex

	Ccm = goccm.New(5000)
)

type HashData struct {
	X    int
	Y    int
	Hash string
}

type Recognizer struct {
	TaskType   string
	Question   string
	Target     string
	TaskList   []TaskList
	EntityType string
	Requester  map[string]map[string]string
	OMut, TMut sync.RWMutex
}

type TaskList struct {
	DatapointURI string `json:"datapoint_uri"`
	TaskKey      string `json:"task_key"`
}

type SolveRepsonse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}
