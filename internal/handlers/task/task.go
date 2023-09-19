package task

import (
	"time"

	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/config"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/database"
	"github.com/Implex-ltd/hcsolver/internal/hcaptcha"
	"github.com/Implex-ltd/hcsolver/internal/model"
	"github.com/surrealdb/surrealdb.go"
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
)

const (
	TYPE_ENTERPRISE = 0
	TYPE_NORMAL     = 1

	SUBMIT = 5000
)

func CreateTask(c *fiber.Ctx) error {
	db := database.DB
	task := new(model.Task)

	var taskData BodyNewSolveTask
	err := c.BodyParser(&taskData)
	if err != nil {
		config.Logger.Error("error-CreateTask", zap.Error(err))
		return c.Status(500).JSON(fiber.Map{
			"message": "Review your input",
			"data":    err.Error(),
			"status":  "error",
		})
	}

	switch taskData.Turbo {
	case true:
		if taskData.TurboSt > 10000 {
			taskData.TurboSt = 10000
		}
	case false:
		taskData.TurboSt = SUBMIT
	}

	T, err := Newtask(&hcaptcha.Config{
		UserAgent:       taskData.UserAgent,
		SiteKey:         taskData.SiteKey,
		Domain:          taskData.Domain,
		Proxy:           taskData.Proxy,
		Logger:          config.Logger,
		TaskType:        taskData.TaskType,
		Invisible:       taskData.Invisible,
		FreeTextEntry:   taskData.FreeTextEntry,
		Turbo:           taskData.Turbo,
		TurboSt:         taskData.TurboSt,
		HcAccessibility: taskData.HcAccessibility,
		Oneblick:        taskData.OneclickOnly,
	})

	if err := T.Create(); err != nil {
		config.Logger.Error("error-T.Create", zap.Error(err))
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "could not create new task",
			"data":    err.Error(),
		})
	}

	task.Status = T.Status.Status
	task.Token = T.Status.Token

	data, err := db.Create("task", task)
	if err != nil {
		config.Logger.Error("error-", zap.Error(err))
		panic(err)
	}

	// Unmarshal data
	createTask := make([]model.Task, 1)
	err = surrealdb.Unmarshal(data, &createTask)
	if err != nil {
		panic(err)
	}

	go func(task *model.Task) {
		captcha, err := T.Solve()

		if err != nil {
			config.Logger.Error("error-solve", zap.Error(err))

			task.Status = STATUS_ERROR
			task.Success = false
			task.Error = err.Error()

			// Update user
			if _, err = db.Change(createTask[0].ID, task); err != nil {
				config.Logger.Error("error-db.Change", zap.Error(err))
				panic(err)
			}

			go func() {
				time.Sleep(120 * time.Second)
				database.DB.Delete(createTask[0].ID)
			}()

			return
		}

		task.Status = STATUS_SOLVED
		task.Success = true
		task.Token = captcha.GeneratedPassUUID
		task.Expiration = captcha.Expiration

		if _, err = db.Change(createTask[0].ID, task); err != nil {
			panic(err)
		}

		go func() {
			time.Sleep(time.Duration(task.Expiration) * time.Second)
			database.DB.Delete(createTask[0].ID)
		}()
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
		config.Logger.Error("error-GetTask", zap.Error(err))
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
