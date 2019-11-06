package usecase

import (
	"2019_2_default_team/models"
	"2019_2_default_team/tickets"
	"context"
)

type TicketUseCase struct {
	ticketRepo tickets.Repository
}

func NewTicketUseCase(ticketRepo tickets.Repository) *TicketUseCase {
	return &TicketUseCase{
		ticketRepo: ticketRepo,
	}
}

func (t TicketUseCase) GetTicket(ctx context.Context, params *models.RequestTicket) (models.Ticket, error) {
	return t.ticketRepo.GetTicket(ctx, params)
}

func (t TicketUseCase) PostTicket(ctx context.Context, u *models.RegisterTicket) (models.Ticket, error) {
	return t.ticketRepo.PostTicket(ctx, u)
}
