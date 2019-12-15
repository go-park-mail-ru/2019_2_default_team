package useCase

import (
	"context"
	"kino_backend/models"
	"kino_backend/repository"
)

type TicketsUseCase interface {
	GetTicket(ctx context.Context, params *models.RequestTicket) (models.Ticket, error)
	PostTicket(ctx context.Context, u *models.RegisterTicket) (models.Ticket, error)
	CheckTicket(ctx context.Context, u *models.RegisterTicket) (bool, error)
}

type ticketUseCase struct {
	ticketRepo repository.TicketRepository
}

func NewTicketUseCase(ticketRepo repository.TicketRepository) *ticketUseCase {
	return &ticketUseCase{
		ticketRepo: ticketRepo,
	}
}

func (t ticketUseCase) GetTicket(ctx context.Context, params *models.RequestTicket) (models.Ticket, error) {
	profile, err := t.ticketRepo.GetTicketProfileByID(params.TicketID)
	if err != nil {
		return models.Ticket{}, err
	}

	return profile, nil
}

func (t ticketUseCase) PostTicket(ctx context.Context, u *models.RegisterTicket) (models.Ticket, error) {
	newT, err := t.ticketRepo.CreateNewTicket(u)
	if err != nil {
		return models.Ticket{}, err
	}

	return newT, nil
}

func (t ticketUseCase) CheckTicket(ctx context.Context, u *models.RegisterTicket) (bool, error) {
	res, err := t.ticketRepo.CheckTicket(u)
	if err != nil {
		return true, err
	}

	return res, nil
}
