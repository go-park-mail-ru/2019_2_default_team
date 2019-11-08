package repository

import (
	"kino_backend/models"
)

type TicketsRepository interface{
	CreateNewTicket(u *models.RegisterTicket) (models.Ticket, error)
	GetTicketProfileByID(id uint) (models.Ticket, error)
	CheckExistenceOfTicket(n int) (bool, error)
}


