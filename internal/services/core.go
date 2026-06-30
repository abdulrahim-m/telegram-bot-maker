package services

import (
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
