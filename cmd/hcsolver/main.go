package main

import (
	"fmt"

	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/config"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/database"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/router"
	"github.com/Implex-ltd/hcsolver/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	addr, err := utils.Lookup("158.46.169.117")
	if err != nil {
		panic(err)
	}

	fmt.Println(addr)

	config.LoadSettings()
	database.ConnectDB()

	app := fiber.New()
	router.SetupRoutes(app)

	config.Logger.Info("DB Connected and api online")
	if err := app.Listen(fmt.Sprintf(":%d", config.Config.API.Port)); err != nil {
		panic(err)
	}
}
