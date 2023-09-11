package database

import (
	"fmt"

	"github.com/surrealdb/surrealdb.go"
)

var (
	DB *surrealdb.DB
)

func ConnectDB(Ip, User, Pass string, Port int) {
	var err error

	DB, err = surrealdb.New(fmt.Sprintf("ws://%s:%d/rpc", Ip, Port))
	if err != nil {
		panic(err)
	}

	if _, err = DB.Signin(map[string]interface{}{
		"user": User,
		"pass": Pass,
	}); err != nil {
		panic(err)
	}

	if _, err = DB.Use("task", "test"); err != nil {
		panic(err)
	}
}
