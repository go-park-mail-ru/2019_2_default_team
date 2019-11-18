package useCase

import (
	"context"
	"kino_backend/models"
	"kino_backend/repository"
	"kino_backend/utilits/errors"
)

type UsersUseCase interface {
	GetUser(params *models.RequestProfile, auth bool, id uint) (models.Profile, error)
	PostUser(ctx context.Context, u *models.RegisterProfile) (models.Profile, error)
	PutUser(ctx context.Context, id uint, editUser *models.RegisterProfile) error
	CheckExistenceOfEmail(e string) (bool, error)
	CheckExistenceOfNickname(n string) (bool, error)
	GetUserPassword(e string) (models.User, error)
}

type usersUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *usersUseCase {
	return &usersUseCase{
		userRepo: userRepo,
	}
}

func (user usersUseCase) GetUser(params *models.RequestProfile, auth bool, id uint) (models.Profile, error) {

	if params.ID != 0 {
		profile, err := user.userRepo.GetUserProfileByID(params.ID)
		if err != nil {

			return models.Profile{}, err
		}
		return profile, nil

	} else if params.Nickname != "" {
		profile, err := user.userRepo.GetUserProfileByNickname(params.Nickname)
		if err != nil {
			return models.Profile{}, err
		}
		return profile, nil

	} else {
		if !auth {
			return models.Profile{}, errors.UserNotAuthError{}
		}

		profile, err := user.userRepo.GetUserProfileByID(id)
		if err != nil {
			return models.Profile{}, err
		}
		return profile, nil
	}
}

func (user usersUseCase) PostUser(ctx context.Context, u *models.RegisterProfile) (models.Profile, error) {
	newU, err := user.userRepo.CreateNewUser(u)

	if err != nil {
		return models.Profile{}, err
	}

	return newU, nil
}

func (user usersUseCase) PutUser(ctx context.Context, id uint, editUser *models.RegisterProfile) error {
	err := user.userRepo.UpdateUserByID(id, editUser)

	if err != nil {
		return err
	}

	return nil
}

func (user usersUseCase) CheckExistenceOfNickname(n string) (bool, error) {
	return user.userRepo.CheckExistenceOfNickname(n)
}

func (user usersUseCase) CheckExistenceOfEmail(e string) (bool, error) {
	return user.userRepo.CheckExistenceOfEmail(e)
}

func (user usersUseCase) GetUserPassword(e string) (models.User, error) {
	return user.userRepo.GetUserPassword(e)
}
