package task

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/surrealdb/surrealdb.go"
	"go.uber.org/zap"

	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/config"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/database"
	"github.com/Implex-ltd/hcsolver/internal/handlers/task/validator"
	"github.com/Implex-ltd/hcsolver/internal/handlers/user"
	"github.com/Implex-ltd/hcsolver/internal/hcaptcha"
	"github.com/Implex-ltd/hcsolver/internal/model"
)

const (
	TYPE_ENTERPRISE = 0
	TYPE_NORMAL     = 1

	SUBMIT = 7000
	MIN    = 1000
)

var (
	SitekeyPattern = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	DomainPattern  = regexp.MustCompile(`^(?i)[a-z0-9-]+(\.[a-z0-9-]+)+$`)
)

func IsValidUUID(input string) bool {
	return SitekeyPattern.MatchString(input)
}

func IsDomainName(input string) bool {
	return DomainPattern.MatchString(input)
}

func checkBody(B BodyNewSolveTask) (errors []string) {
	if B.Domain == "" {
		errors = append(errors, "domain is missing")
	} else {
		if !IsDomainName(B.Domain) {
			errors = append(errors, "domain is invalid")
		}
	}

	if B.SiteKey == "" {
		errors = append(errors, "site_key is missing")
	} else {
		if !IsValidUUID(B.SiteKey) {
			errors = append(errors, "site_key is invalid")
		}
	}

	if B.TaskType != TYPE_ENTERPRISE /* && B.TaskType != TYPE_NORMAL */ {
		errors = append(errors, "task_type is invid")
	}

	if len(B.UserAgent) > 255 {
		errors = append(errors, "user-agent is invalid")
	}

	if B.Proxy == "" {
		errors = append(errors, "please provide proxy")
	} else {
		if len(B.Proxy) > 500 || !strings.HasPrefix(B.Proxy, "http") {
			errors = append(errors, "invalid proxy format, please use http(s)://user:pass@ip:port or http(s)://ip:port")
		}
	}

	if B.Rqdata != "" {
		if len(B.Rqdata) > 15000 {
			errors = append(errors, "rqdata seems too long, please contact support")
		}
	}

	if B.HcAccessibility != "" {
		if len(B.HcAccessibility) > 15000 {
			errors = append(errors, "hc_accessibility seems too long, please contact support")
		}
	}

	if !B.FreeTextEntry {
		errors = append(errors, "please enable a11y_tfe, if your target doesn't support it, we can't solve your captchas for now")
	}

	if len(errors) != 0 {
		return errors
	}

	return nil
}

func CreateTask(c *fiber.Ctx) error {
	authToken, ok := c.GetReqHeaders()["Authorization"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    "you must provide api-key for this endpoint",
		})
	}
	auth := authToken[0]

	db := database.TaskDB
	task := new(model.Task)

	var taskData BodyNewSolveTask

	if err := c.BodyParser(&taskData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    "invalid task body",
		})
	}

	if err := checkBody(taskData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    fmt.Sprintf("invalid task body fields: %v", strings.Join(err, ", ")),
		})
	}

	settings, err := validator.Validate(taskData.SiteKey)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    err.Error(),
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

	if settings != nil {
		if !settings.Enabled {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"data":    "this site-key is disabled, please contact support",
			})
		}

		if settings.AlwaysText && !taskData.FreeTextEntry {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"data":    fmt.Sprintf("bad config 'a11y_tfe' for the current domain '%s', please check required settings using '/api/misc/check/%s'", taskData.Domain, taskData.Domain),
			})
		}

		if taskData.Turbo {
			if taskData.TurboSt < settings.MinSubmitTime || taskData.TurboSt > settings.MaxSubmitTime {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"success": false,
					"data":    fmt.Sprintf("bad config 'turbo_st' for the current domain '%s', please check required settings using '/api/misc/check/%s'", taskData.Domain, taskData.Domain),
				})
			}
		}

		if !taskData.OneclickOnly && settings.OneclickOnly {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"data":    fmt.Sprintf("bad config 'oneclick_only' for the current domain '%s', please check required settings using '/api/misc/check/%s'", taskData.Domain, taskData.Domain),
			})
		}

		if taskData.Domain != settings.Domain {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"data":    "domain doesn't match the site-key",
			})
		}
	}

	authUser, err := user.GetUserByID(auth)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    "invalid api-key",
		})
	}

	selectedUser := new(model.User)
	err = surrealdb.Unmarshal(authUser, &selectedUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"data":    "failed to parse user, please report to admin",
		})
	}

	if selectedUser.Balance <= 0 {
		return c.Status(fiber.StatusTeapot).JSON(fiber.Map{
			"success": false,
			"data":    "no credit left, please refill balance",
		})
	}

	if selectedUser.ThreadUsedHcaptcha >= selectedUser.ThreadMaxHcaptcha {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"success": false,
			"data":    "you reached max thread limit, please buy more slot",
		})
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

	if data == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    "your payload is probably wrong, or a backend issue. please read documentation.",
		})
	}

	createTask := make([]model.Task, 1)

	if err := surrealdb.Unmarshal(data, &createTask); err != nil {
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

		task.Req = captcha.C.Req
		task.Status = STATUS_SOLVED
		task.Success = true
		task.Token = captcha.GeneratedPassUUID
		task.Expiration = captcha.Expiration
		task.UserAgent = T.Captcha.Manager.Manager.Fingerprint.Browser.UserAgent

		if _, err := db.Update(createTask[0].ID, task); err != nil {
			config.Logger.Error("db-error", zap.Error(err))
			return
		}

		go func() {
			database.UserDB.Query(`
BEGIN TRANSACTION;

UPDATE $user SET balance -= $to_add;
UPDATE $user SET solved_hcaptcha += $to_add;

COMMIT TRANSACTION;
			`, map[string]any{
				"user":   auth,
				"to_add": 1,
			})

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

	if !strings.HasPrefix(id, "task:") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    "please provide 'task:id' format",
		})
	}

	data, err := database.TaskDB.Select(id)
	if err != nil {
		config.Logger.Error("error-GetTask", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func GetSitekeySettings(c *fiber.Ctx) error {
	id := c.Params("siteKey")

	if !IsValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    "invalid site-key",
		})
	}

	settings, err := validator.Validate(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    err.Error(),
		})
	}

	if settings == nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": true,
			"data":    "this site-key doesn't have any restrictions",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"data":    settings,
		})
	}
}
