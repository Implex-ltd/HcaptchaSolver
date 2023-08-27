package main

import (
	"fmt"

	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/config"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/database"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadSettings()
	database.ConnectDB()

	app := fiber.New()
	router.SetupRoutes(app)

	config.Logger.Info("DB Connected and api online")
	if err := app.Listen(fmt.Sprintf(":%d", config.Config.API.Port)); err != nil {
		panic(err)
	}
}
