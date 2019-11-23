package repository

import (
	"kino_backend/models"
)

type SupportChatsRepository interface {
	CreateMessage(m *models.Message) (*models.Message, error)
	GetAllGlobalMessages() (*[]models.Message, error)
}
