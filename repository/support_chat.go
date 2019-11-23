package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"kino_backend/db"
	"kino_backend/models"
	"kino_backend/utilits/errors"
)

type SupportChatRepository struct {
	database *sqlx.DB
}

func NewSupportChatRepository(db *sqlx.DB) SupportChatRepository {
	return SupportChatRepository{
		database: db,
	}
}

func (SC SupportChatRepository) CreateMessage(m *models.Message) (*models.Message, error) {
	var err error
	qres := SC.database.QueryRowx(`
	INSERT INTO message (author_id, to_user, message_text)
	VALUES ($1, $2, $3) RETURNING *`,
		m.Author, m.To, m.Message)
	if err := qres.Err(); err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Code == "23502" {
			return nil, db.ErrNotNullConstraintViolation
		}
	}
	res := &models.Message{}
	err = qres.StructScan(res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (SC SupportChatRepository) GetAllGlobalMessages() (*[]models.Message, error) {

	res := &[]models.Message{}
	err := SC.database.Select(res, `SELECT * FROM message
		WHERE to_user IS NULL
		ORDER BY created`)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (SC SupportChatRepository) GetAllDialogs() ([]models.USERTOSUP, error) {
	res := []models.USERTOSUP{}
	resOne := models.USERTOSUP{}

	qres, err := SC.database.Queryx(`
		SELECT * FROM user_to_sup`)
	if err != nil {
		return res, err
	}

	for qres.Next() {
		err = qres.StructScan(&resOne)
		res = append(res, resOne)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.FilmNotFoundError{}
		}
		return []models.USERTOSUP{}, err
	}

	return res, nil
}

func (SC SupportChatRepository) GetSupStatus(id int) (models.Status, error) {
	res := models.Status{}
	qres := SC.database.QueryRowx(`
		SELECT status FROM support
		WHERE id = $1`,
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

func (SC SupportChatRepository) GetAllSupportDialogs(id int) ([]models.SupportDialog, error) {
	res := []models.SupportDialog{}
	resOne := models.SupportDialog{}

	qres, err := SC.database.Queryx(`
		SELECT uts_id FROM user_to_sup WHERE sup_id = ?`, id)
	if err != nil {
		return res, err
	}

	for qres.Next() {
		err = qres.StructScan(&resOne)
		res = append(res, resOne)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.FilmNotFoundError{}
		}
		return []models.SupportDialog{}, err
	}

	return res, nil
}

func (SC SupportChatRepository) GetClientDialog(id int) ([]models.OneMessage, error) {
	res := []models.OneMessage{}
	resOne := models.OneMessage{}

	qres, err := SC.database.Queryx(`
		SELECT message, time FROM message WHERE author_id = ?`, id)
	if err != nil {
		return res, err
	}

	for qres.Next() {
		err = qres.StructScan(&resOne)
		res = append(res, resOne)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.FilmNotFoundError{}
		}
		return []models.OneMessage{}, err
	}

	return res, nil
}
