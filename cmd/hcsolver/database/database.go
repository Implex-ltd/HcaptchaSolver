package database

import (
	"github.com/surrealdb/surrealdb.go"
)

var (
	DB *surrealdb.DB
)

func ConnectDB() {
	var err error

	DB, err = surrealdb.New("ws://104.194.129.150:8000/rpc")
	if err != nil {
		panic(err)
	}

	if _, err = DB.Signin(map[string]interface{}{
		"user": "root",
		"pass": "vivithlerlmao",
	}); err != nil {
		panic(err)
	}

	if _, err = DB.Use("task", "test"); err != nil {
		panic(err)
	}
}