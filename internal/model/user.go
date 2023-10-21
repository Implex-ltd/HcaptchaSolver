package model

import "github.com/surrealdb/surrealdb.go"

type User struct {
	surrealdb.Basemodel `table:"user"`

	ID      string `json:"id,omitempty"`
	Balance int    `json:"balance"`

	// Hcaptcha
	SolvedHcaptcha     int `json:"solved_hcaptcha"`
	ThreadUsedHcaptcha int `json:"thread_used_hcaptcha"`
	ThreadMaxHcaptcha  int `json:"thread_max_hcaptcha"`
}
