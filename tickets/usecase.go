package tickets

import (
	"2019_2_default_team/models"
	"context"
)

type UseCase interface {
	GetTicket(ctx context.Context, params *models.RequestTicket) (models.Ticket, error)
	PostTicket(ctx context.Context, u *models.RegisterTicket) (models.Ticket, error)
}
