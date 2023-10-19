package model

import "github.com/surrealdb/surrealdb.go"

type Task struct {
	surrealdb.Basemodel `table:"task"`

	ID         string `json:"id,omitempty"`
	Status     int    `json:"status"`
	Token      string `json:"token"`
	Error      string `json:"error"`
	Success    bool   `json:"success"`
	Expiration int    `json:"expiration"`
	UserAgent  string `json:"user_agent"`
}
