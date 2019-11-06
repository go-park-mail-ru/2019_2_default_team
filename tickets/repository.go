package tickets

import (
	"2019_2_default_team/models"
	"context"
)

type Repository interface {
	GetTicket(ctx context.Context, params *models.RequestTicket) (models.Ticket, error)
	PostTicket(ctx context.Context, u *models.RegisterTicket) (models.Ticket, error)
}
