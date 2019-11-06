package usecase

import (
	"2019_2_default_team/models"
	"2019_2_default_team/users"
	"context"
)

type UserUseCase struct {
	userRepo users.Repository
}

func NewUserUseCase(userRepo users.Repository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (user UserUseCase) GetUser(ctx context.Context, params *models.RequestProfile, auth bool, id uint) (models.Profile, error) {
	return user.userRepo.GetUser(ctx, params, auth, id)
}

func (user UserUseCase) PostUser(ctx context.Context, u *models.RegisterProfile) (models.Profile, error) {
	return user.userRepo.PostUser(ctx, u)
}

func (user UserUseCase) PutUser(ctx context.Context, id uint, userInfo *models.RegisterProfile) error {
	return user.userRepo.PutUser(ctx, id, userInfo)
}
