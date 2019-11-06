package usecase

import (
	"2019_2_default_team/films"
	"2019_2_default_team/models"
	"context"
)

type FilmUseCase struct {
	filmRepo films.Repository
}

func NewFilmUseCase(filmRepo films.Repository) *FilmUseCase {
	return &FilmUseCase{
		filmRepo: filmRepo,
	}
}

func (f FilmUseCase) GetFilm(ctx context.Context, params *models.RequestProfileFilm) (models.ProfileFilm, error) {
	return f.filmRepo.GetFilm(ctx, params)
}

func (f FilmUseCase) PostFilmUse(ctx context.Context, u *models.RegisterProfileFilm) (models.ProfileFilm, error) {
	return f.filmRepo.PostFilm(ctx, u)
}

func (f FilmUseCase) PutFilm(ctx context.Context, filmInfo *models.ProfileFilm) error {
	return f.filmRepo.UpdateFilm(ctx, filmInfo)
}
