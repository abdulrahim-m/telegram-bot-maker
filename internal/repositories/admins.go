package repositories

import "github.com/abdulrahim-m/telegram-bot-maker/internal/models"

type AdminRepo struct {
	BaseRepository[models.Admin]
}
