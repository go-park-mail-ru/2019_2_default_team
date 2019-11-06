package films

import (
	"context"
	"kino_backend/models"
)

type Repository interface{
	GetFilm(ctx context.Context, params *models.RequestProfileFilm) (models.ProfileFilm, error)
	PostFilm(ctx context.Context, u *models.RegisterProfileFilm) (models.ProfileFilm, error)
	UpdateFilm(ctx context.Context, filmInfo *models.ProfileFilm) (error)
}