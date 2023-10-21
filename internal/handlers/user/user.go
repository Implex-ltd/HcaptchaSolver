package user

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/config"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/database"
	"github.com/Implex-ltd/hcsolver/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/surrealdb/surrealdb.go"
	"go.uber.org/zap"
)

var (
	mut = sync.Mutex{}
)

/*
Example:

	if err := EditUser(user, func(u *model.User) {
			u.Balance += balance
		}); err != nil {
			return err
		}
*/
func EditUser(user string, change func(*model.User)) (out interface{}, err error) {
	mut.Lock()
	defer mut.Unlock()

	data, err := database.UserDB.Select(user)
	if err != nil {
		return nil, err
	}

	selectedUser := new(model.User)
	err = surrealdb.Unmarshal(data, &selectedUser)
	if err != nil {
		return nil, err
	}

	change(selectedUser)

	if out, err := database.UserDB.Update(user, selectedUser); err != nil {
		return out, err
	}

	return out, nil
}

func GetUserByID(user string) (interface{}, error) {
	if !strings.HasPrefix(user, "user:") {
		return nil, errors.New("please provide 'user:id' format")
	}

	data, err := database.UserDB.Select(user)
	if err != nil {
		return false, nil
	}

	return data, nil
}

func CreateUser(c *fiber.Ctx) error {
	db := database.UserDB

	data, err := db.Create("user", model.User{
		Balance:               0.0,
		SolvedHcaptcha:        0,
		ThreadMaxHcaptcha:     100,
		ThreadUsedHcaptcha:    0,
		BypassRestrictedSites: false,
	})

	if err != nil {
		config.Logger.Error("db-error", zap.Error(err))

		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"data":    fmt.Sprintf("db error (please report to admin): %v", err.Error()),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("userId")

	if id == "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"data":    "invalid user ID",
		})
	}

	data, err := GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func AddBalance(c *fiber.Ctx) error {
	var data BodyAddBalance

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    "invalid task body",
		})
	}

	if !strings.HasPrefix(data.User, "user:") || data.Amount <= 0.0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    "invalid task body",
		})
	}

	user, err := EditUser(data.User, func(u *model.User) {
		u.Balance += data.Amount
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func SetypassRestricted(c *fiber.Ctx) error {
	var data BodySetBypass

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    "invalid task body",
		})
	}

	if !strings.HasPrefix(data.User, "user:") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"data":    "invalid task body",
		})
	}

	user, err := EditUser(data.User, func(u *model.User) {
		u.BypassRestrictedSites = data.Enabled
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}
