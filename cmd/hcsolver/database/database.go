package database

import (
	"fmt"

	"github.com/surrealdb/surrealdb.go"
)

var (
	TaskDB *surrealdb.DB
	FpDB   *surrealdb.DB
	UserDB *surrealdb.DB
)

func ConnectDB(Ip, User, Pass string, Port int) {
	var err error

	TaskDB, err = surrealdb.New(fmt.Sprintf("ws://%s:%d/rpc", Ip, Port))
	if err != nil {
		panic(err)
	}

	if _, err = TaskDB.Signin(map[string]interface{}{
		"user": User,
		"pass": Pass,
	}); err != nil {
		panic(err)
	}

	if _, err = TaskDB.Use("task", "test"); err != nil {
		panic(err)
	}

	FpDB, err = surrealdb.New(fmt.Sprintf("ws://%s:%d/rpc", Ip, Port))
	if err != nil {
		panic(err)
	}

	if _, err = FpDB.Signin(map[string]interface{}{
		"user": User,
		"pass": Pass,
	}); err != nil {
		panic(err)
	}

	if _, err = FpDB.Use("fingerprint", "fp"); err != nil {
		panic(err)
	}

	UserDB, err = surrealdb.New(fmt.Sprintf("ws://%s:%d/rpc", Ip, Port))
	if err != nil {
		panic(err)
	}

	if _, err = UserDB.Signin(map[string]interface{}{
		"user": User,
		"pass": Pass,
	}); err != nil {
		panic(err)
	}

	if _, err = UserDB.Use("users", "user"); err != nil {
		panic(err)
	}
}
