package users_usecase

import (
	"context"
	"kino_backend/models"
	"kino_backend/repository"
)

type UserUseCase struct{
	userRepo repository.UsersRepository
}

func NewUserUseCase(userRepo repository.UsersRepository) *UserUseCase{
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (user UserUseCase) GetUser(ctx context.Context, params *models.RequestProfile, auth bool, id uint) (models.Profile, error){
	return user.userRepo.GetUser(ctx, params, auth, id)
}

func (user UserUseCase) PostUser(ctx context.Context, u *models.RegisterProfile) (models.Profile, error){
	return user.userRepo.PostUser(ctx, u)
}

func (user UserUseCase) PutUser(ctx context.Context, id uint, userInfo *models.RegisterProfile) error{
	return user.userRepo.PutUser(ctx, id, userInfo)
}

