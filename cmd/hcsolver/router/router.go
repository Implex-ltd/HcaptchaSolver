package router

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/config"

	taskRoutes "github.com/Implex-ltd/hcsolver/internal/routes/task"
	userRoutes "github.com/Implex-ltd/hcsolver/internal/routes/user"

	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	api.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1" || c.IP() == "92.149.51.174"
		},
		Max:        config.Config.Ratelimit.APIMax,
		Expiration: time.Duration(config.Config.Ratelimit.APIExpiration) * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"data":    errors.New("ratelimit exceeded"),
			})
		},
	}))

	taskRoutes.SetupTaskRoutes(api)
	userRoutes.SetupUserRoutes(api)
}
