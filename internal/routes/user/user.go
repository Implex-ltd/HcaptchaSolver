package userRoutes

import (
	"crypto/sha256"
	"crypto/subtle"

	userHandler "github.com/Implex-ltd/hcsolver/internal/handlers/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

/*
	- Authorization header required
	- /api/internal/user
*/

var (
	apiKey = "Cr4p0nT0pD54sdljn4D"
)

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {
	hashedAPIKey := sha256.Sum256([]byte(apiKey))
	hashedKey := sha256.Sum256([]byte(key))

	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	user.Get("/:userId", userHandler.GetUser)

	internalApi := user.Group("/internal", logger.New())
	internalApi.Use(keyauth.New(keyauth.Config{
		KeyLookup: "header:Authorization",
		Validator: validateAPIKey,
	}))

	internalApi.Post("/new", userHandler.CreateUser)
	internalApi.Post("/refill", userHandler.AddBalance)
	internalApi.Post("/set-bypass", userHandler.SetypassRestricted)
}
