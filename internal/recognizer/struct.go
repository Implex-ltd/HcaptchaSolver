package recognizer

import (
	"net/http"
	"sync"
)

var (
	Hashlist   = new(sync.Map)
	Selectlist = new(sync.Map)
	Answerlist = new(sync.Map)
)

type HashData struct {
	X    int
	Y    int
	Hash string
}

type Recognizer struct {
	Http       *http.Client
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

	// a11y_challenge
	DatapointHash string            `json:"datapoint_hash"`
	DatapointText map[string]string `json:"datapoint_text"`
}

type SolveResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

type AnswerStruct struct {
	Text string `json:"text"`
}
