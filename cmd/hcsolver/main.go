package main

import (
	"encoding/base64"
	"fmt"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/config"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/database"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/router"
	_ "github.com/Implex-ltd/hcsolver/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/surrealdb/surrealdb.go"
)

type Fingerprint struct {
	ID          string `json:"id,omitempty"`
	Fingerprint string `json:"fp"`
}

func check_fp() {
	req, err := database.FpDB.Select("fp")
	if err != nil {
		panic(err)
	}

	var FingerprintSlice []Fingerprint
	err = surrealdb.Unmarshal(req, &FingerprintSlice)
	if err != nil {
		panic(err)
	}

	valid := 0
	for _, fp := range FingerprintSlice {
		fmt.Println(fp.ID)
		_, err := base64.RawStdEncoding.DecodeString(fp.Fingerprint)
		if err != nil {
			continue
		}
		valid++

	}

	fmt.Println("valid fp:", valid)
	fmt.Println("total fp:", len(FingerprintSlice))
}

func main() {
	config.LoadSettings()
	database.ConnectDB(config.Config.Database.IP, config.Config.Database.Username, config.Config.Database.Password, config.Config.Database.Port)

	app := fiber.New()
	router.SetupRoutes(app)

	config.Logger.Info("DB Connected and api online")
	if err := app.Listen(fmt.Sprintf(":%d", config.Config.API.Port)); err != nil {
		panic(err)
	}
}
