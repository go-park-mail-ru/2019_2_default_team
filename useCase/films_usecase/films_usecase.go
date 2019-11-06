package films_usecase

import (
	"context"
	"kino_backend/repository"
	"kino_backend/models"
)

type FilmUseCase struct{
	filmRepo repository.FilmsRepository
}

func NewFilmUseCase(filmRepo repository.FilmsRepository) *FilmUseCase{
	return &FilmUseCase{
		filmRepo: filmRepo,
	}
}

func (f FilmUseCase) GetFilm(ctx context.Context, params *models.RequestProfileFilm) (models.ProfileFilm, error){
	return f.filmRepo.GetFilm(ctx, params)
}

func (f FilmUseCase) PostFilmUse(ctx context.Context, u *models.RegisterProfileFilm) (models.ProfileFilm, error){
	return f.filmRepo.PostFilm(ctx, u)
}

func (f FilmUseCase) PutFilm(ctx context.Context, filmInfo *models.ProfileFilm) error{
	return f.filmRepo.UpdateFilm(ctx, filmInfo)
}


