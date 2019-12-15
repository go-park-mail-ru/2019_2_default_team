package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"kino_backend/models"
	"kino_backend/utilits/errors"
)

type TicketRepository struct {
	database *sqlx.DB
}

func NewTicketRepository(db *sqlx.DB) TicketRepository {
	return TicketRepository{
		database: db,
	}
}

func (TR TicketRepository) CheckTicket(u *models.RegisterTicket) (bool, error) {
	res := []models.Ticket{}
	resOne := models.Ticket{}

	qres, err := TR.database.Queryx(`
		SELECT ticket_id, movie_session_id, seat_id, profile_id FROM ticket_profile
		WHERE movie_session_id = $1 AND seat_id = $2`,
		u.MSID, u.SeatID)
	if err != nil {
		return false, err
	}

	for qres.Next() {
		err = qres.StructScan(&resOne)
		res = append(res, resOne)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	if len(res) != 0 {
		return true, nil
	}

	return false, nil
}

func (TR TicketRepository) CreateNewTicket(u *models.RegisterTicket) (models.Ticket, error) {
	res := models.Ticket{}
	qres := TR.database.QueryRowx(`
		INSERT INTO ticket_profile (profile_id, movie_session_id, seat_id, price)
		VALUES ($1, $2, $3, $4) RETURNING ticket_id, movie_session_id`,
		u.UserID, u.MSID, u.SeatID, u.Price)

	fmt.Println(u.SeatID)

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

	qt := TR.database.QueryRowx(`
		UPDATE seat SET is_taken = true WHERE seat_id = $1`,
		u.SeatID)

	if err := qt.Err(); err != nil {
		pqErr := err.(*pq.Error)
		switch pqErr.Code {
		case "23502":
			return res, errors.ErrNotNullConstraintViolation
		case "23505":
			return res, errors.ErrUniqueConstraintViolation
		}
	}

	return res, nil
}

func (TR TicketRepository) GetTicketProfileByID(id uint) (models.Ticket, error) {
	res := models.Ticket{}
	qres := TR.database.QueryRowx(`
		SELECT ticket_id, profile_id, movie_session_id, seat_id, price FROM ticket_profile
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

func (TR TicketRepository) CheckExistenceOfTicket(n int) (bool, error) {
	res := models.Ticket{}
	qres := TR.database.QueryRowx(`
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
