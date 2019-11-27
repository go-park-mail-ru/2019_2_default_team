package repository

import (
	"kino_backend/models"
)

type SupportChatsRepository interface {
	CreateMessage(m *models.Message) (*models.Message, error)
	GetAllGlobalMessages() (*[]models.Message, error)
	GetClientDialog(id int) ([]models.OneMessage, error)
	GetAllSupportDialogs(id int) ([]models.SupportDialog, error)
	GetSupStatus(id int) (models.Status, error)
	GetAllDialogs() ([]models.USERTOSUP, error)
}
