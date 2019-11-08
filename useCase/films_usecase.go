package useCase

import (
	"context"
	"kino_backend/db"
	"kino_backend/models"
	"kino_backend/repository"
)

type FilmsUseCase interface{
	GetFilm(ctx context.Context, params *models.RequestProfileFilm) (models.ProfileFilm, error)
	PostFilmUse(ctx context.Context, u *models.RegisterProfileFilm) (models.ProfileFilm, error)
	PutFilm(ctx context.Context, filmInfo *models.ProfileFilm) error
}

type filmUseCase struct{
	filmRepo repository.FilmRepository
}

func NewFilmUseCase(filmRepo repository.FilmRepository) *filmUseCase {
	return &filmUseCase{
		filmRepo: filmRepo,
	}
}

func (f filmUseCase) GetFilm(ctx context.Context, params *models.RequestProfileFilm) (models.ProfileFilm, error){
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

func (f filmUseCase) PostFilmUse(ctx context.Context, u *models.RegisterProfileFilm) (models.ProfileFilm, error){
	newF, err := db.CreateNewFilm(u)

	if err != nil {
		return models.ProfileFilm{}, err
	}

	return newF, err
}

func (f filmUseCase) PutFilm(ctx context.Context, filmInfo *models.ProfileFilm) error{
	err := db.UpdateFilmByID(filmInfo.FilmID, filmInfo)
	return err
}

