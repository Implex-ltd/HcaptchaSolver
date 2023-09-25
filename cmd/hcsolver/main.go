package main

import (
	"fmt"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/config"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/database"
	//"github.com/Implex-ltd/hcsolver/cmd/hcsolver/database"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/router"
	_ "github.com/Implex-ltd/hcsolver/internal/utils"
	"github.com/gofiber/fiber/v2"

	/*"encoding/base64"
	"github.com/surrealdb/surrealdb.go"
	"log"*/
)

type Fingerprint struct {
	ID          string `json:"id,omitempty"`
	Fingerprint string `json:"fp"`
}

func main() {
	/*
		addr, err := utils.Lookup("158.46.169.117")
		if err != nil {
			panic(err)
		}

		fmt.Println(addr)
	*/

	config.LoadSettings()

	database.ConnectDB(config.Config.Database.IP, config.Config.Database.Username, config.Config.Database.Password, config.Config.Database.Port)

	/*
		req, err := database.FpDB.Select("fp")
		if err != nil {
			panic(err)
		}

		var FingerprintSlice []Fingerprint
		err = surrealdb.Unmarshal(req, &FingerprintSlice)
		if err != nil {
			panic(err)
		}
		for _, fp := range FingerprintSlice {
			fmt.Println(fp.ID)
			if fp.ID == "fp:fc49974c8afd714dab25042617bf5d36" {
				fmt.Println(fp.Fingerprint)
				b, err := base64.RawStdEncoding.DecodeString(fp.Fingerprint)
				if err != nil {
					log.Println(err)
					continue
				}
				log.Println(string(b))
			}

		}
	*/

	app := fiber.New()
	router.SetupRoutes(app)

	config.Logger.Info("DB Connected and api online")
	if err := app.Listen(fmt.Sprintf(":%d", config.Config.API.Port)); err != nil {
		panic(err)
	}
}
