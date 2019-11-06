package tickets

import (
	"context"
	"kino_backend/models"
)

type UseCase interface {
	GetTicket(ctx context.Context, params *models.RequestTicket) (models.Ticket, error)
	PostTicket(ctx context.Context, u *models.RegisterTicket) (models.Ticket, error)
}
