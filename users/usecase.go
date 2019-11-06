package users

import (
	"context"
	"kino_backend/models"
	)

type UseCase interface {
	GetUser(ctx context.Context, params *models.RequestProfile, auth bool, id uint) (models.Profile, error)
	PostUser(ctx context.Context, u *models.RegisterProfile) (models.Profile, error)
	PutUser(ctx context.Context, id uint, editUser *models.RegisterProfile) (error)
}