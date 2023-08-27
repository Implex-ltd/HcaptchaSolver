package taskRoutes

import (
	taskHandler "github.com/Implex-ltd/hcsolver/internal/handlers/task"
	"github.com/gofiber/fiber/v2"
)

func SetupTaskRoutes(router fiber.Router) {
	user := router.Group("/task")

	user.Post("/new", taskHandler.CreateTask)
	user.Get("/:taskId", taskHandler.GetTask)
}
