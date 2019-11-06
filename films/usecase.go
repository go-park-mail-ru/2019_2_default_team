package films

import (
	"2019_2_default_team/models"
	"context"
)

type UseCase interface {
	GetFilm(ctx context.Context, params *models.RequestProfileFilm) (models.ProfileFilm, error)
	PostFilmUse(ctx context.Context, u *models.RegisterProfileFilm) (models.ProfileFilm, error)
	PutFilm(ctx context.Context, filmInfo *models.ProfileFilm) error
}
