package repository

import (
	"context"
	"kino_backend/models"
)

type TicketsRepository interface{
	GetTicket(ctx context.Context, params *models.RequestTicket) (models.Ticket, error)
	PostTicket(ctx context.Context, u *models.RegisterTicket) (models.Ticket, error)
}
