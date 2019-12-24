package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"kino_backend/models"
	"kino_backend/utilits/errors"
	"math/rand"
	"strings"
	"time"
)

type FilmRepository struct {
	database *sqlx.DB
}

func NewFilmRepository(db *sqlx.DB) FilmRepository {
	return FilmRepository{
		database: db,
	}
}

func (FR FilmRepository) CreateNewMovieSession(u *models.RegisterMovieSession, seatsNumber int) (models.MovieSession, error) {
	res := models.MovieSession{}
	qres := FR.database.QueryRowx(`
		INSERT INTO movie_session (hall_name, movie_id, start_datetime, type)
		VALUES ($1, $2, $3, $4) RETURNING ms_id, hall_name, movie_id, start_datetime, type`,
		u.HallName, u.MovieID, u.Date, u.Type)
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

	for i := 0; i < seatsNumber; i++ {
		qres := FR.database.QueryRowx(`
		INSERT INTO seat (movie_session_id, row, seat_number)
		VALUES ($1, $2, $3)`,
			res.MsID, (i+1)/4+1, i+1)
		if err := qres.Err(); err != nil {
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

func (FR FilmRepository) CreateNewFilm(u *models.RegisterProfileFilm) (models.ProfileFilm, error) {
	res := models.ProfileFilm{}
	qres := FR.database.QueryRowx(`
		INSERT INTO film_profile (title, description, director, actors, admin_id, genre, length, production, year)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING film_id, title, director`,
		u.Title, u.Description, u.Director, u.MainActor, u.AdminID, u.Genre, u.Length, u.Production, u.Year)
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

func (FR FilmRepository) UpdateFilmByID(id uint, u *models.ProfileFilm) error {
	if u.Title == "" {
		return nil
	}

	q := strings.Builder{}
	q.WriteString(`
		UPDATE user_profile
		SET `)
	hasBefore := false
	if u.Title != "" {
		q.WriteString("title = :title")
		hasBefore = true
	}
	if u.Description != "" {
		if hasBefore {
			q.WriteString(", ")
		}
		q.WriteString("description = :description")
	}
	if u.Director != "" {
		if hasBefore {
			q.WriteString(", ")
		}
		q.WriteString("director = :director")
	}
	if u.MainActor != "" {
		if hasBefore {
			q.WriteString(", ")
		}
		q.WriteString("mainactor = :mainactor")
	}

	q.WriteString(`
		WHERE film_id = :film_id`)

	_, err := FR.database.NamedExec(q.String(), &models.ProfileFilm{
		Film: models.Film{
			FilmID: id,
		},
		Title:       u.Title,
		Description: u.Description,
		Director:    u.Director,
		MainActor:   u.MainActor,
	})
	if err != nil {
		return err
	}

	return nil
}

func (FR FilmRepository) GetMovieSessionsForToday(movie_id uint) ([]models.RequestFilmTimes, error) {
	res := []models.RequestFilmTimes{}
	resOne := models.RequestFilmTimes{}

	qres, err := FR.database.Queryx(`
		SELECT start_datetime, ms_id, hall_name FROM movie_session
		WHERE movie_id = $1 AND start_datetime > now() at time zone 'msk' AND start_datetime < (now()::date at time zone 'msk' + interval '24h')`,
		movie_id)
	if err != nil {
		return res, err
	}

	for qres.Next() {
		err = qres.StructScan(&resOne)
		res = append(res, resOne)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.FilmNotFoundError{"title"}
		}
		return []models.RequestFilmTimes{}, err
	}

	return res, nil
}

func (FR FilmRepository) GetFilmsForToday() ([]models.ProfileFilm, error) {
	res := []models.MovieSession{}
	resOne := models.MovieSession{}
	film := models.ProfileFilm{}
	films := []models.ProfileFilm{}

	qres, err := FR.database.Queryx(`
		SELECT movie_id FROM movie_session
		 WHERE start_datetime > now() at time zone 'msk' AND start_datetime < (now()::date at time zone 'msk' + interval '24h')`)
	if err != nil {
		return films, err
	}

	for qres.Next() {
		err = qres.StructScan(&resOne)
		res = append(res, resOne)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return films, errors.FilmNotFoundError{"title"}
		}
		return films, err
	}

	var filmsIDs []uint

	for _, value := range res {
		filmsIDs = append(filmsIDs, value.MovieID)
	}

	query, args, err := sqlx.In("SELECT film_id, title, description, director, actors, admin_id, genre, length, production, year, rating, poster_popup, poster, trailer FROM film_profile WHERE film_id IN (?);", filmsIDs)

	// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
	query = FR.database.Rebind(query)
	qresfilms, err := FR.database.Queryx(query, args...)
	if err != nil {
		return films, err
	}

	for qresfilms.Next() {
		err = qresfilms.StructScan(&film)
		films = append(films, film)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return films, errors.FilmNotFoundError{"title"}
		}
		return films, err
	}

	return films, nil
}

func (FR FilmRepository) GetFilmsForSoon() ([]models.ProfileFilm, error) {
	res := []models.MovieSession{}
	resOne := models.MovieSession{}
	film := models.ProfileFilm{}
	films := []models.ProfileFilm{}

	qres, err := FR.database.Queryx(`
		SELECT movie_id FROM movie_session
		 WHERE start_datetime > (now()::date at time zone 'msk' + interval '24h')`)
	if err != nil {
		return films, err
	}

	for qres.Next() {
		err = qres.StructScan(&resOne)
		res = append(res, resOne)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return films, errors.FilmNotFoundError{"title"}
		}
		return films, err
	}

	var filmsIDs []uint

	for _, value := range res {
		filmsIDs = append(filmsIDs, value.MovieID)
	}

	query, args, err := sqlx.In("SELECT film_id, title, description, director, actors, admin_id, genre, length, production, year, rating, poster_popup, poster, trailer FROM film_profile WHERE film_id IN (?);", filmsIDs)

	// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
	query = FR.database.Rebind(query)
	qresfilms, err := FR.database.Queryx(query, args...)
	if err != nil {
		return films, err
	}

	for qresfilms.Next() {
		err = qresfilms.StructScan(&film)
		films = append(films, film)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return films, errors.FilmNotFoundError{"title"}
		}
		return films, err
	}

	return films, nil
}

func (FR FilmRepository) GetSeatsByMSID(movie_session_id uint) ([]models.Seat, error) {
	res := []models.Seat{}
	resOne := models.Seat{}

	qres, err := FR.database.Queryx(`
		SELECT seat_id, movie_session_id, is_taken, row, seat_number FROM seat
		WHERE movie_session_id = $1`,
		movie_session_id)
	if err != nil {
		return res, err
	}

	for qres.Next() {
		err = qres.StructScan(&resOne)
		res = append(res, resOne)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.MSNotFoundError{"movie_session_id"}
		}
		return []models.Seat{}, err
	}

	return res, nil
}

func (FR FilmRepository) GetFilmProfileByID(id uint) (models.ProfileFilm, error) {
	res := models.ProfileFilm{}
	qres := FR.database.QueryRowx(`
		SELECT film_id, title, description, director, actors, admin_id, genre, length, production, year, rating, poster_popup, poster, trailer FROM film_profile
		WHERE film_id = $1`,
		id)
	if err := qres.Err(); err != nil {
		return res, err
	}
	err := qres.StructScan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.FilmNotFoundError{"id"}
		}
		return res, err
	}

	return res, nil
}

func (FR FilmRepository) GetFilmProfileByTitle(title string) (models.ProfileFilm, error) {
	res := models.ProfileFilm{}
	qres := FR.database.QueryRowx(`
		SELECT film_id, title, description, director, actors, admin_id, genre, length, production, year, rating, poster_popup, poster, trailer FROM film_profile
		WHERE title = $1`,
		title)
	if err := qres.Err(); err != nil {
		return res, err
	}
	err := qres.StructScan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.FilmNotFoundError{"title"}
		}
		return res, err
	}

	return res, nil
}

func (FR FilmRepository) CheckExistenceOfTitle(n string) (bool, error) {
	res := models.ProfileFilm{}
	qres := FR.database.QueryRowx(`
		SELECT FROM film_profile
		WHERE title = $1`,
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

func (FR FilmRepository) GetCountOfFilms() (int, error) {
	res := 0
	// TODO: optimize it
	qres := FR.database.QueryRowx(`
		SELECT COUNT(*) FROM film_profile`)
	if err := qres.Err(); err != nil {
		return 0, err
	}
	err := qres.Scan(&res)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (FR FilmRepository) UploadAvatarFilm(uID uint, path string) error {
	qres, err := FR.database.Exec(`
		UPDATE film_profile
		SET avatar = $2
		WHERE film_id = $1`,
		uID, path)
	if err != nil {
		return err
	}
	res, err := qres.RowsAffected()
	if err != nil {
		return err
	}
	if res == 0 {
		return &errors.FilmNotFoundError{"id"}
	}

	return nil
}

func (FR FilmRepository) DeleteAvatarFilm(uID uint) error {
	qres, err := FR.database.Exec(`
		UPDATE film_profile
		SET avatar = NULL
		WHERE film_id = $1`,
		uID)
	if err != nil {
		return err
	}
	res, err := qres.RowsAffected()
	if err != nil {
		return err
	}
	if res == 0 {
		return &errors.FilmNotFoundError{"id"}
	}

	return nil
}

func (FR FilmRepository) GetFilmsForDate(startTime, lastTime time.Time, movie_id uint) (bool, error) {
	res := []models.RequestFilmTimes{}
	resOne := models.RequestFilmTimes{}

	qres, err := FR.database.Queryx(`
		SELECT start_datetime, ms_id, hall_name FROM movie_session
		WHERE movie_id = $1`,
		movie_id)

	if err != nil {
		fmt.Println("erroring")
		return false, err
	}

	for qres.Next() {
		err = qres.StructScan(&resOne)
		res = append(res, resOne)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.FilmNotFoundError{"title"}
		}
		return false, err
	}
	fmt.Println("in res", res)

	for _, value := range res {
		fmt.Println("starttime", startTime, "lasttime", lastTime, "value", value)
		fmt.Println(value.Date.After(startTime))
		if value.Date.After(startTime) && value.Date.Before(lastTime) {
			return true, nil
		}
	}

	return false, nil
}

func (FR FilmRepository) GetFilmsForPrice(minPrice, maxPrice int, filmId uint) (bool, error) {
	res := []models.MovieSessionSeat{}
	resOne := models.MovieSessionSeat{}

	qres, err := FR.database.Queryx(`
		SELECT a.movie_id, b.price FROM movie_session a INNER JOIN seat b ON a.ms_id = b.movie_session_id
		 WHERE b.price > $1 AND b.price < $2 AND a.movie_id = $3`, minPrice, maxPrice, filmId)

	if err != nil {
		fmt.Println("erro1")
		return false, err
	}

	for qres.Next() {
		err = qres.StructScan(&resOne)
		res = append(res, resOne)
	}

	if err != nil {
		fmt.Println("erro2")
		if err == sql.ErrNoRows {
			return false, errors.FilmNotFoundError{"title"}
		}
		return false, err
	}
	fmt.Println("resdate", res)

	if len(res) > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (FR FilmRepository) GetAllFilms() ([]models.ProfileFilm, error) {
	res := []models.ProfileFilm{}
	resOne := models.ProfileFilm{}

	qres, err := FR.database.Queryx(`
		SELECT film_id, title, description, director, actors, admin_id, genre, length, production, year, rating, poster_popup, poster, trailer FROM film_profile
		WHERE is_deleted = $1`,
		false)
	if err != nil {
		return res, err
	}

	for qres.Next() {
		err = qres.StructScan(&resOne)
		res = append(res, resOne)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.FilmNotFoundError{"title"}
		}
		return []models.ProfileFilm{}, err
	}

	return res, nil
}

func (FR FilmRepository) GetRecommendedFilms(wantedGenre string) ([]models.ProfileFilm, error) {
	res := []models.ProfileFilm{}
	result := []models.ProfileFilm{}
	resOne := models.ProfileFilm{}

	qres, err := FR.database.Queryx(`
		SELECT film_id, title, description, director, actors, admin_id, genre, length, production, year, rating, poster_popup, poster, trailer FROM film_profile
		WHERE is_deleted = $1 AND genre = $2`,
		false, wantedGenre)
	if err != nil {
		return res, err
	}

	for qres.Next() {
		err = qres.StructScan(&resOne)
		res = append(res, resOne)
	}

	if len(res) <= 3 {
		result = res
	} else {
		rand1 := rand.Intn(len(res))
		rand2 := rand.Intn(len(res))
		for rand2 == rand1 {
			rand2 = rand.Intn(len(res))
		}
		rand3 := rand.Intn(len(res))
		for rand3 == rand1 || rand3 == rand2 {
			rand3 = rand.Intn(len(res))
		}
		for index, value := range res {
			if index == rand1 || index == rand2 || index == rand3 {
				result = append(result, value)
			}
		}
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.FilmNotFoundError{"title"}
		}
		return []models.ProfileFilm{}, err
	}

	return result, nil
}

func (FR FilmRepository) VoteForFilm(u *models.RegisterVote) (models.Vote, error) {
	res := models.Vote{}
	qres := FR.database.QueryRowx(`
		INSERT INTO rating (user_id, film_id)
		VALUES ($1, $2) RETURNING user_id, film_id, vote_id`,
		u.UserID, u.MovieID)
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

	_, err = FR.database.Queryx(`
		UPDATE film_profile SET rating = rating + 1 WHERE film_id = $1 AND
		is_deleted = $2`,
		u.MovieID, false)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (FR FilmRepository) IsVoted(u *models.RegisterVote) (bool, error) {
	res := []models.Vote{}
	resOne := models.Vote{}

	qres, err := FR.database.Queryx(`
		SELECT vote_id, user_id, film_id FROM rating
		WHERE film_id = $1 AND user_id = $2`,
		u.MovieID, u.UserID)
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
