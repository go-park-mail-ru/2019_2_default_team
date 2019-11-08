package repository

import(
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"kino_backend/models"
	"kino_backend/db"
	"kino_backend/utilits/errors"
)

type TicketRepository struct{
	database *sqlx.DB
}

func NewTicketRepository(db *sqlx.DB) TicketRepository{
	return TicketRepository{
		database:db,
	}
}


func CreateNewTicket(u *models.RegisterTicket) (models.Ticket, error) {
	res := models.Ticket{}
	qres := db.Db.QueryRowx(`
		INSERT INTO ticket_profile (user_id, film_id)
		VALUES ($1, $2) RETURNING ticket_id, film_id`,
		u.UserID, u.FilmID)
	if err := qres.Err(); err != nil {
		pqErr := err.(*pq.Error)
		switch pqErr.Code {
		case "23502":
			return res, errors.ErrNotNullConstraintViolation
		case "23505":
			return res, errors.ErrUniqueConstraintViolation
		}
	}
	err := qres.StructScan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func GetTicketProfileByID(id uint) (models.Ticket, error) {
	res := models.Ticket{}
	qres := db.Db.QueryRowx(`
		SELECT ticket_id, user_id, film_id FROM ticket_profile
		WHERE ticket_id = $1`,
		id)
	if err := qres.Err(); err != nil {
		return res, err
	}
	err := qres.StructScan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.TicketNotFoundError{"id"}
		}
		return res, err
	}

	return res, nil
}


func CheckExistenceOfTicket(n int) (bool, error) {
	res := models.Ticket{}
	qres := db.Db.QueryRowx(`
		SELECT FROM ticket_profile
		WHERE ticket_id = $1`,
		n)
	if err := qres.Err(); err != nil {
		return false, err
	}
	err := qres.StructScan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
