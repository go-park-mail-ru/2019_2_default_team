package Postgres

import (
	"2019_2_default_team/db"
	"2019_2_default_team/models"
	"context"

	"github.com/jmoiron/sqlx"
)

type TicketRepository struct {
	database *sqlx.DB
}

func NewTicketRepository(db *sqlx.DB) *TicketRepository {
	return &TicketRepository{
		database: db,
	}
}

func (t TicketRepository) GetTicket(ctx context.Context, params *models.RequestTicket) (models.Ticket, error) {
	profile, err := db.GetTicketProfileByID(params.TicketID)
	if err != nil {
		return models.Ticket{}, err
	}

	return profile, nil
}

func (t TicketRepository) PostTicket(ctx context.Context, u *models.RegisterTicket) (models.Ticket, error) {
	newT, err := db.CreateNewTicket(u)
	if err != nil {
		return models.Ticket{}, err
	}

	return newT, nil
}
