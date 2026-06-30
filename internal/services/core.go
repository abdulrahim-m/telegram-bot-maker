package services

import (
	"time"

	"github.com/abdulrahim-m/telegram-bot-maker/internal"
	"github.com/abdulrahim-m/telegram-bot-maker/internal/models"
	"github.com/abdulrahim-m/telegram-bot-maker/internal/repositories"
	"github.com/jmoiron/sqlx"
)

type Core struct {
	AdminRepo repositories.AdminRepo
}

func InitCoreService(db sqlx.DB) *Core {
	return &Core{
		*repositories.NewAdminRepo(db, "admins"),
	}
}

func (c *Core) IsFirstLaunch() (bool, error) {
	admins, err := c.AdminRepo.GetAll()
	if err != nil {
		return false, err
	}

	return len(admins) == 0, nil
}

func (c *Core) InitAdmin(username, password string) error {
	passhash, err := internal.HashPassword(password)
	if err != nil {
		return err
	}

	return c.AdminRepo.Create(&models.Admin{
		Username:  username,
		Password:  passhash,
		Role:      models.RoleCreator,
		CreatedAt: time.Now(),
	})
}
