package task

import (
	"fmt"
	"strings"
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

	SUBMIT = 7000
	MIN    = 1000
)

func checkBody(B BodyNewSolveTask) (errors []string) {
	if B.Domain == "" {
		errors = append(errors, "domain")
	}

	if B.SiteKey == "" {
		errors = append(errors, "site_key")
	}

	if B.TaskType != TYPE_ENTERPRISE && B.TaskType != TYPE_NORMAL {
		errors = append(errors, "task_type")
	}

	if len(errors) != 0 {
		return errors
	}

	return nil
}

func CreateTask(c *fiber.Ctx) error {
	db := database.TaskDB
	task := new(model.Task)

	var taskData BodyNewSolveTask

	if err := c.BodyParser(&taskData); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"data":    "invalid task body",
		})
	}

	if err := checkBody(taskData); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"data":    fmt.Errorf("invalid task body fields: %v", strings.Join(err, ", ")),
		})
	}

	switch taskData.Turbo {
	case true:
		if taskData.TurboSt < 1000 {
			taskData.TurboSt = 1000
		}
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
		OneClick:        taskData.OneclickOnly,
		Rqdata:          taskData.Rqdata,
	})

	if err := T.Create(); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"data":    fmt.Sprintf("could not create new task: %v", err.Error()),
		})
	}

	task.Status = T.Status.Status
	task.Token = T.Status.Token

	data, err := db.Create("task", task)
	if err != nil {
		config.Logger.Error("db-error", zap.Error(err))

		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"data":    fmt.Sprintf("db error (please report to admin): %v", err.Error()),
		})
	}

	createTask := make([]model.Task, 1)
	err = surrealdb.Unmarshal(data, &createTask)
	if err != nil {
		config.Logger.Error("db-error", zap.Error(err))

		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"data":    fmt.Sprintf("db error (please report to admin): %v", err.Error()),
		})
	}

	go func(task *model.Task) {
		captcha, err := T.Solve()

		if err != nil {
			config.Logger.Error("error-solve", zap.Error(err))

			task.Status = STATUS_ERROR
			task.Success = false
			task.Error = err.Error()

			if _, err = db.Change(createTask[0].ID, task); err != nil {
				config.Logger.Error("db-error", zap.Error(err))
				return
			}

			go func() {
				time.Sleep(120 * time.Second)
				database.TaskDB.Delete(createTask[0].ID)
			}()

			return
		}

		task.Status = STATUS_SOLVED
		task.Success = true
		task.Token = captcha.GeneratedPassUUID
		task.Expiration = captcha.Expiration

		if _, err := db.Update(createTask[0].ID, task); err != nil {
			config.Logger.Error("db-error", zap.Error(err))
			return
		}

		go func() {
			time.Sleep(time.Duration(task.Expiration) * time.Second)
			database.TaskDB.Delete(createTask[0].ID)
		}()
	}(task)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func GetTask(c *fiber.Ctx) error {
	id := c.Params("taskId")

	data, err := database.TaskDB.Select(id)
	if err != nil {
		config.Logger.Error("error-GetTask", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
