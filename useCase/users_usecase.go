package useCase

import (
	"context"
	"kino_backend/db"
	"kino_backend/models"
	"kino_backend/repository"
	"kino_backend/utilits/errors"
)

type UsersUseCase interface {
	GetUser(params *models.RequestProfile, auth bool, id uint) (models.Profile, error)
	PostUser(ctx context.Context, u *models.RegisterProfile) (models.Profile, error)
	PutUser(ctx context.Context, id uint, editUser *models.RegisterProfile) (error)
}


type usersUseCase struct{
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *usersUseCase {
	return &usersUseCase{
		userRepo: userRepo,
	}
}

func (user usersUseCase) GetUser( params *models.RequestProfile, auth bool, id uint) (models.Profile, error){

	if params.ID != 0 {
		profile, err := db.GetUserProfileByID(params.ID)
		if err != nil {

			return models.Profile{}, err
		}
		return profile, nil

	} else if params.Nickname != "" {
		profile, err := db.GetUserProfileByNickname(params.Nickname)
		if err != nil {
			return models.Profile{}, err
		}
		return profile, nil

	} else {
		if !auth {
			return models.Profile{}, errors.UserNotAuthError{}
		}

		profile, err := db.GetUserProfileByID(id)
		if err != nil {
			return models.Profile{}, err
		}
		return profile, nil
	}
}

func (user usersUseCase) PostUser(ctx context.Context, u *models.RegisterProfile) (models.Profile, error){
	newU, err := db.CreateNewUser(u)

	if err != nil {
		return models.Profile{}, err
	}

	return newU, nil
}

func (user usersUseCase) PutUser(ctx context.Context, id uint, editUser *models.RegisterProfile) error{
	err := db.UpdateUserByID(id, editUser)

	if err != nil{
		return err
	}

	return nil
}

