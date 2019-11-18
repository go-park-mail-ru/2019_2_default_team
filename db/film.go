package db

//
//import (
//	"database/sql"
//	"fmt"
//	"github.com/lib/pq"
//	"strings"
//	"kino_backend/models"
//)
//
//
//
//func CreateNewFilm(u *models.RegisterProfileFilm) (models.ProfileFilm, error) {
//	res := models.ProfileFilm{}
//	fmt.Println(u)
//	qres := Db.QueryRowx(`
//		INSERT INTO film_profile (title, description, director, mainactor, admin_id)
//		VALUES ($1, $2, $3, $4, $5) RETURNING film_id, title, director`,
//		u.Title, u.Description, u.Director, u.MainActor, u.AdminID)
//	if err := qres.Err(); err != nil {
//		pqErr := err.(*pq.Error)
//		switch pqErr.Code {
//		case "23502":
//			return res, ErrNotNullConstraintViolation
//		case "23505":
//			return res, ErrUniqueConstraintViolation
//		}
//	}
//	err := qres.StructScan(&res)
//	if err != nil {
//		return res, err
//	}
//
//	return res, nil
//}
//
//func UpdateFilmByID(id uint, u *models.ProfileFilm) error {
//	if u.Title == "" {
//		return nil
//	}
//
//	q := strings.Builder{}
//	q.WriteString(`
//		UPDATE user_profile
//		SET `)
//	hasBefore := false
//	if u.Title != "" {
//		q.WriteString("title = :title")
//		hasBefore = true
//	}
//	if u.Description != "" {
//		if hasBefore {
//			q.WriteString(", ")
//		}
//		q.WriteString("description = :description")
//	}
//	if u.Director != "" {
//		if hasBefore {
//			q.WriteString(", ")
//		}
//		q.WriteString("director = :director")
//	}
//	if u.MainActor != "" {
//		if hasBefore {
//			q.WriteString(", ")
//		}
//		q.WriteString("mainactor = :mainactor")
//	}
//
//	q.WriteString(`
//		WHERE film_id = :film_id`)
//
//	_, err := Db.NamedExec(q.String(), &models.ProfileFilm{
//		Film: models.Film{
//			FilmID: id,
//		},
//		Title: u.Title,
//		Description: u.Description,
//		Director: u.Director,
//		MainActor: u.MainActor,
//	})
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func GetFilmProfileByID(id uint) (models.ProfileFilm, error) {
//	res := models.ProfileFilm{}
//	qres := Db.QueryRowx(`
//		SELECT film_id, title, description, avatar, director, mainactor, admin_id FROM film_profile
//		WHERE film_id = $1`,
//		id)
//	if err := qres.Err(); err != nil {
//		return res, err
//	}
//	err := qres.StructScan(&res)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return res, FilmNotFoundError{"id"}
//		}
//		return res, err
//	}
//
//	return res, nil
//}
//
//func GetFilmProfileByTitle(title string) (models.ProfileFilm, error) {
//	res := models.ProfileFilm{}
//	qres := Db.QueryRowx(`
//		SELECT film_id, title, description, avatar, director, mainactor, admin_id FROM film_profile
//		WHERE title = $1`,
//		title)
//	if err := qres.Err(); err != nil {
//		return res, err
//	}
//	err := qres.StructScan(&res)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return res, FilmNotFoundError{"title"}
//		}
//		return res, err
//	}
//
//	return res, nil
//}
//
//
//func CheckExistenceOfTitle(n string) (bool, error) {
//	res := models.ProfileFilm{}
//	qres := Db.QueryRowx(`
//		SELECT FROM film_profile
//		WHERE title = $1`,
//		n)
//	if err := qres.Err(); err != nil {
//		return false, err
//	}
//	err := qres.StructScan(&res)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return false, nil
//		}
//		return false, err
//	}
//
//	return true, nil
//}
//
//func GetCountOfFilms() (int, error) {
//	res := 0
//	// TODO: optimize it
//	qres := Db.QueryRowx(`
//		SELECT COUNT(*) FROM film_profile`)
//	if err := qres.Err(); err != nil {
//		return 0, err
//	}
//	err := qres.Scan(&res)
//	if err != nil {
//		return 0, err
//	}
//
//	return res, nil
//}
//
//func UploadAvatarFilm(uID uint, path string) error {
//	qres, err := Db.Exec(`
//		UPDATE film_profile
//		SET avatar = $2
//		WHERE film_id = $1`,
//		uID, path)
//	if err != nil {
//		return err
//	}
//	res, err := qres.RowsAffected()
//	if err != nil {
//		return err
//	}
//	if res == 0 {
//		return &FilmNotFoundError{"id"}
//	}
//
//	return nil
//}
//
//func DeleteAvatarFilm(uID uint) error {
//	qres, err := Db.Exec(`
//		UPDATE film_profile
//		SET avatar = NULL
//		WHERE film_id = $1`,
//		uID)
//	if err != nil {
//		return err
//	}
//	res, err := qres.RowsAffected()
//	if err != nil {
//		return err
//	}
//	if res == 0 {
//		return &FilmNotFoundError{"id"}
//	}
//
//	return nil
//}
