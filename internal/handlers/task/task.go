package task

import (
	"fmt"

	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/config"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/database"
	"github.com/Implex-ltd/hcsolver/internal/hcaptcha"
	"github.com/Implex-ltd/hcsolver/internal/model"
	"github.com/surrealdb/surrealdb.go"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	db := database.DB
	task := new(model.Task)

	var taskData BodyNewSolveTask
	err := c.BodyParser(&taskData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Review your input",
			"data":    err.Error(),
			"status":  "error",
		})
	}

	T, err := Newtask(&hcaptcha.Config{
		UserAgent: taskData.UserAgent,
		SiteKey:   taskData.SiteKey,
		Domain:    taskData.Domain,
		Proxy:     taskData.Proxy,
		Logger:    config.Logger,
	})

	if err := T.Create(); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "could not create new task",
			"data":    err.Error(),
		})
	}

	//task.ID = T.ID
	task.Status = T.Status.Status
	task.Token = T.Status.Token

	data, err := db.Create("task", task)
	if err != nil {
		panic(err)
	}

	// Unmarshal data
	createTask := make([]model.Task, 1)
	err = surrealdb.Unmarshal(data, &createTask)
	if err != nil {
		panic(err)
	}

	fmt.Println(createTask)

	go func(task *model.Task) {
		captcha, err := T.Solve()

		if err != nil {
			task.Status = STATUS_ERROR
			task.Success = false
			task.Error = err.Error()

			// Update user
			if _, err = db.Change(createTask[0].ID, task); err != nil {
				panic(err)
			}
			return
		}

		task.Status = STATUS_SOLVED
		task.Success = true
		task.Token = captcha.GeneratedPassUUID
		task.Expiration = captcha.Expiration

		if _, err = db.Change(createTask[0].ID, task); err != nil {
			panic(err)
		}
	}(task)

	return c.Status(200).JSON(fiber.Map{
		"message": "Created task",
		"success": true,
		"data":    data,
	})
}

func GetTask(c *fiber.Ctx) error {
	id := c.Params("taskId")

	data, err := database.DB.Select(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "task Found",
		"data":    data,
	})
}
