package tickets_usecase

import (
	"context"
	"kino_backend/models"
	"kino_backend/repository"
)

type TicketUseCase struct{
	ticketRepo repository.TicketsRepository
}

func NewTicketUseCase(ticketRepo repository.TicketsRepository) *TicketUseCase{
	return &TicketUseCase{
		ticketRepo: ticketRepo,
	}
}

func (t TicketUseCase) GetTicket(ctx context.Context, params *models.RequestTicket) (models.Ticket, error){
	return t.ticketRepo.GetTicket(ctx, params)
}

func (t TicketUseCase) PostTicket(ctx context.Context, u *models.RegisterTicket) (models.Ticket, error){
	return t.ticketRepo.PostTicket(ctx, u)
}



