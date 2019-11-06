package Postgres

import(
	"context"
	"github.com/jmoiron/sqlx"
	"kino_backend/models"
	"kino_backend/db"
)

type TicketRepository struct{
	database *sqlx.DB
}

func NewTicketRepository(db *sqlx.DB) *TicketRepository{
	return &TicketRepository{
		database:db,
	}
}

func (t TicketRepository) GetTicket(ctx context.Context, params *models.RequestTicket) (models.Ticket, error) {
	profile, err := db.GetTicketProfileByID(params.TicketID)
	if err != nil {
		return models.Ticket{}, err
	}

	return profile, nil
}


func (t TicketRepository) PostTicket(ctx context.Context, u *models.RegisterTicket) (models.Ticket, error){
	newT, err := db.CreateNewTicket(u)
	if err != nil {
		return models.Ticket{}, err
	}

	return newT, nil
}
