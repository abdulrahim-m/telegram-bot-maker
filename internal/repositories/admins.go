package repositories

import (
	"github.com/abdulrahim-m/telegram-bot-maker/internal/models"
	"github.com/jmoiron/sqlx"
)

type AdminRepo struct {
	BaseRepository[models.Admin]
}

func NewAdminRepo(db sqlx.DB, table string) *AdminRepo {
	return &AdminRepo{
		BaseRepository: BaseRepository[models.Admin]{
			DB:    db,
			Table: table,
		},
	}
}
