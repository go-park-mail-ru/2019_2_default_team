package Postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"kino_backend/models"
	"kino_backend/db"
)

type FilmRepository struct{
	database *sqlx.DB
}

func NewFilmRepository(db *sqlx.DB) *FilmRepository{
	return &FilmRepository{
		database: db,
	}
}

func (fm FilmRepository) GetFilm(ctx context.Context, params *models.RequestProfileFilm) (models.ProfileFilm, error) {

	var err error
	var profile models.ProfileFilm

	if params.ID != 0 {
		profile, err = db.GetFilmProfileByID(params.ID)
		if err != nil {
			return models.ProfileFilm{}, err
		}
		return profile, nil
	} else if params.Title != "" {
		profile, err = db.GetFilmProfileByTitle(params.Title)

		if err != nil {
			return models.ProfileFilm{}, err
		}

		return profile, nil
	}
	return profile, nil
}

func (fm FilmRepository) PostFilm(ctx context.Context, u *models.RegisterProfileFilm) (models.ProfileFilm, error){
	newF, err := db.CreateNewFilm(u)

	if err != nil {
		return models.ProfileFilm{}, err
	}

	return newF, err
}

func (fm FilmRepository) UpdateFilm(ctx context.Context, filmInfo *models.ProfileFilm) (error){
	err := db.UpdateFilmByID(filmInfo.FilmID, filmInfo)
	return err
}
