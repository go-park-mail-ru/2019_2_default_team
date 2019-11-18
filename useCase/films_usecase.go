package useCase

import (
	"context"
	"kino_backend/models"
	"kino_backend/repository"
)

type FilmsUseCase interface {
	GetFilm(ctx context.Context, params *models.RequestProfileFilm) (models.ProfileFilm, error)
	PostFilmUse(ctx context.Context, u *models.RegisterProfileFilm) (models.ProfileFilm, error)
	PutFilm(ctx context.Context, filmInfo *models.ProfileFilm) error
}

type filmUseCase struct {
	filmRepo repository.FilmRepository
}

func NewFilmUseCase(filmRepo repository.FilmRepository) *filmUseCase {
	return &filmUseCase{
		filmRepo: filmRepo,
	}
}

func (f filmUseCase) GetFilm(ctx context.Context, params *models.RequestProfileFilm) (models.ProfileFilm, error) {
	var err error
	var profile models.ProfileFilm

	if params.ID != 0 {
		profile, err = f.filmRepo.GetFilmProfileByID(params.ID)
		if err != nil {
			return models.ProfileFilm{}, err
		}
		return profile, nil
	} else if params.Title != "" {
		profile, err = f.filmRepo.GetFilmProfileByTitle(params.Title)

		if err != nil {
			return models.ProfileFilm{}, err
		}

		return profile, nil
	}
	return profile, nil
}

func (f filmUseCase) PostFilmUse(ctx context.Context, u *models.RegisterProfileFilm) (models.ProfileFilm, error) {
	newF, err := f.filmRepo.CreateNewFilm(u)

	if err != nil {
		return models.ProfileFilm{}, err
	}

	return newF, err
}

func (f filmUseCase) PutFilm(ctx context.Context, filmInfo *models.ProfileFilm) error {
	err := f.filmRepo.UpdateFilmByID(filmInfo.FilmID, filmInfo)
	return err
}
