package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"kino_backend/models"
	"kino_backend/utilits/errors"
	"strings"
	"time"
)

type UserRepository struct {
	database *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{
		database: db,
	}
}

func (UR UserRepository) GetUserPassword(e string) (models.User, error) {
	res := models.User{}
	qres := UR.database.QueryRowx(`
		SELECT user_id, email, password FROM user_profile
		WHERE email = $1`,
		e)
	if err := qres.Err(); err != nil {
		return res, err
	}
	err := qres.StructScan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.UserNotFoundError{"email"}
		}
		return res, err
	}

	return res, nil
}

func (UR UserRepository) CreateNewUser(u *models.RegisterProfile) (models.Profile, error) {
	res := models.Profile{}
	qres := UR.database.QueryRowx(`
		INSERT INTO user_profile (email, password, nickname, first_name, last_name)
		VALUES ($1, $2, $3, $4, $5) RETURNING user_id, email, nickname`,
		u.Email, u.Password, u.Nickname, u.FirstName, u.LastName)
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

	for _, value := range u.Genres {
		qresg := UR.database.QueryRowx(`
		INSERT INTO user_genres (user_id, genre)
		VALUES ($1, $2)`,
			res.UserID, value.LovelyGenre)

		if err := qresg.Err(); err != nil {
			pqErr := err.(*pq.Error)
			switch pqErr.Code {
			case "23502":
				return res, errors.ErrNotNullConstraintViolation
			case "23505":
				return res, errors.ErrUniqueConstraintViolation
			}
		}
	}

	return res, nil
}

func (UR UserRepository) UpdateUserByID(id uint, u *models.RegisterProfile) error {
	if u.Email == "" && u.Password == "" && u.Nickname == "" {
		return nil
	}

	q := strings.Builder{}
	q.WriteString(`
		UPDATE user_profile
		SET `)
	hasBefore := false
	if u.Email != "" {
		q.WriteString("email = :email")
		hasBefore = true
	}
	if u.Password != "" {
		if hasBefore {
			q.WriteString(", ")
		}
		q.WriteString("password = :password")
		hasBefore = true
	}
	if u.Nickname != "" {
		if hasBefore {
			q.WriteString(", ")
		}
		q.WriteString("nickname = :nickname")
	}
	q.WriteString(`
		WHERE user_id = :user_id`)

	_, err := UR.database.NamedExec(q.String(), &models.Profile{
		User: models.User{
			UserID: id,
			UserPassword: models.UserPassword{
				Email:    u.Email,
				Password: u.Password,
			},
		},
		Nickname: u.Nickname,
	})
	if err != nil {
		return err
	}

	return nil
}

func (UR UserRepository) GetUserProfileByID(id uint) (models.FullProfile, error) {
	res := models.FullProfile{}
	qres := UR.database.QueryRowx(`
		SELECT user_id, email, nickname, avatar, first_name, last_name FROM user_profile
		WHERE user_id = $1`,
		id)
	if err := qres.Err(); err != nil {
		return res, err
	}
	err := qres.StructScan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.UserNotFoundError{"id"}
		}
		return res, err
	}

	resT := []models.TicketProfilePro{}
	resOneT := models.TicketProfilePro{}

	qrestickets, err := UR.database.Queryx(`
		SELECT a.ticket_id, a.movie_session_id, a.seat_id, a.profile_id, a.price, a.start_datetime, b.row, b.seat_number FROM ticket_profile a INNER JOIN seat b ON a.seat_id = b.seat_id
		WHERE profile_id = $1`,
		id)
	if err != nil {
		return res, err
	}

	for qrestickets.Next() {
		err = qrestickets.StructScan(&resOneT)
		resT = append(resT, resOneT)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.TicketNotFoundError{""}
		}
		return models.FullProfile{}, err
	}

	ticketAdd := models.TicketAddInfo{}
	//ticketAdds :=  []models.TicketAddInfo{}
	var ticketMap map[uint]models.TicketAddInfo
	var msIDs []uint

	ticketMap = make(map[uint]models.TicketAddInfo)

	for _, value := range resT {
		msIDs = append(msIDs, value.MSID)
	}

	query, args, err := sqlx.In("SELECT a.ms_id, b.poster_popup, a.hall_name, b.title FROM movie_session a INNER JOIN film_profile b ON a.movie_id = b.film_id WHERE a.ms_id IN (?);", msIDs)

	// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
	query = UR.database.Rebind(query)
	qresfilms, err := UR.database.Queryx(query, args...)
	if err != nil {
		return res, err
	}

	for qresfilms.Next() {
		err = qresfilms.StructScan(&ticketAdd)
		ticketMap[ticketAdd.MsID] = ticketAdd
		//ticketAdds = append(ticketAdds, ticketAdd)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.FilmNotFoundError{"title"}
		}
		return res, err
	}

	for index, value := range resT {
		resT[index].PosterPopup = ticketMap[value.MSID].PosterPopup
		resT[index].Title = ticketMap[value.MSID].Title
		resT[index].HallName = ticketMap[value.MSID].HallName
	}

	for _, value := range resT {
		if value.TicketProfile.Date.After(time.Now().AddDate(0, 0, 3)) {
			res.Tickets = append(res.Tickets, value)
		} else {
			res.TicketsHistory = append(res.TicketsHistory, value)
		}
	}

	resG := []models.Genre{}
	resOneG := models.Genre{}

	qresgenres, err := UR.database.Queryx(`
		SELECT genre FROM user_genres
		WHERE user_id = $1`,
		id)
	if err != nil {
		return res, err
	}

	for qresgenres.Next() {
		err = qresgenres.StructScan(&resOneG)
		resG = append(resG, resOneG)
	}

	res.Genres = resG

	if res.TicketsHistory == nil {
		res.TicketsHistory = []models.TicketProfilePro{}
	}

	if res.Tickets == nil {
		res.Tickets = []models.TicketProfilePro{}
	}

	return res, nil
}

func (UR UserRepository) GetUserProfileByNickname(nickname string) (models.Profile, error) {
	res := models.Profile{}
	qres := UR.database.QueryRowx(`
		SELECT user_id, email, nickname, avatar, first_name, last_name FROM user_profile
		WHERE nickname = $1`,
		nickname)
	if err := qres.Err(); err != nil {
		return res, err
	}
	err := qres.StructScan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.UserNotFoundError{"nickname"}
		}
		return res, err
	}

	return res, nil
}

func (UR UserRepository) CheckExistenceOfEmail(e string) (bool, error) {
	res := models.Profile{}
	qres := UR.database.QueryRowx(`
		SELECT FROM user_profile
		WHERE email = $1`,
		e)
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

func (UR UserRepository) CheckExistenceOfNickname(n string) (bool, error) {
	res := models.Profile{}
	qres := UR.database.QueryRowx(`
		SELECT FROM user_profile
		WHERE nickname = $1`,
		n)
	if err := qres.Err(); err != nil {
		fmt.Println("c1")
		return false, err
	}
	err := qres.StructScan(&res)
	if err != nil {
		fmt.Println("c2")
		if err == sql.ErrNoRows {
			fmt.Println("c3")
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (UR UserRepository) GetCountOfUsers() (int, error) {
	res := 0
	// TODO: optimize it
	qres := UR.database.QueryRowx(`
		SELECT COUNT(*) FROM user_profile`)
	if err := qres.Err(); err != nil {
		return 0, err
	}
	err := qres.Scan(&res)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (UR UserRepository) UploadAvatar(uID uint, path string) error {
	qres, err := UR.database.Exec(`
		UPDATE user_profile
		SET avatar = $2
		WHERE user_id = $1`,
		uID, path)
	if err != nil {
		return err
	}
	res, err := qres.RowsAffected()
	if err != nil {
		return err
	}
	if res == 0 {
		return &errors.UserNotFoundError{"id"}
	}

	return nil
}

func (UR UserRepository) DeleteAvatar(uID uint) error {
	qres, err := UR.database.Exec(`
		UPDATE user_profile
		SET avatar = NULL
		WHERE user_id = $1`,
		uID)
	if err != nil {
		return err
	}
	res, err := qres.RowsAffected()
	if err != nil {
		return err
	}
	if res == 0 {
		return &errors.UserNotFoundError{"id"}
	}

	return nil
}
