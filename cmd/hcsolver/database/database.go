package database

import (
	"fmt"

	"github.com/Implex-ltd/hcsolver/internal/model"
	"github.com/Implex-ltd/hcsolver/internal/utils"
	"github.com/surrealdb/surrealdb.go"
)

var (
	TaskDB *surrealdb.DB
	FpDB   *surrealdb.DB
	UserDB *surrealdb.DB
)

func mergeAll() {
	UserDB.Query(`
	UPDATE user MERGE {
		settings: {
			bypass_restricted_sites: false,
		},
	};	
	`, nil)
}

func showUsers() {
	req, err := UserDB.Select("user")
	if err != nil {
		panic(err)
	}

	var FingerprintSlice []model.User
	err = surrealdb.Unmarshal(req, &FingerprintSlice)
	if err != nil {
		panic(err)
	}

	for _, fp := range FingerprintSlice {
		fmt.Println(fp.BypassRestrictedSites)

	}
}

func doBackup() {
	req, err := UserDB.Select("user")
	if err != nil {
		panic(err)
	}

	var FingerprintSlice []model.User
	err = surrealdb.Unmarshal(req, &FingerprintSlice)
	if err != nil {
		panic(err)
	}

	for _, fp := range FingerprintSlice {
		line := fmt.Sprintf("%s,%d,%d,%v", fp.ID, fp.Balance, fp.SolvedHcaptcha, fp.BypassRestrictedSites)
		fmt.Println(line)

		utils.AppendLine(line, "backup.csv")
	}
}

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

	//mergeAll()

	//showUsers()
	doBackup()
}
