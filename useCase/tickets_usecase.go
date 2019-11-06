package useCase

import (
	"context"
	"kino_backend/models"
)

type TicketsUseCase interface {
	GetTicket(ctx context.Context, params *models.RequestTicket) (models.Ticket, error)
	PostTicket(ctx context.Context, u *models.RegisterTicket) (models.Ticket, error)
}
