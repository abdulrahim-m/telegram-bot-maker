package repositories

import (
	"github.com/abdulrahim-m/telegram-bot-maker/internal/models"
	"github.com/jmoiron/sqlx"
)

type SessionsRepo struct {
	BaseRepository[models.Session]
}

func NewSessionsRepo(db sqlx.DB, table string) *SessionsRepo {
	return &SessionsRepo{
		BaseRepository: BaseRepository[models.Session]{
			DB:    db,
			Table: table,
		},
	}
}

func (a *AdminRepo) FindBySessionID(SessionsID string) (*models.Admin, error) {
	return a.FindByField("session_id", SessionsID)
}
