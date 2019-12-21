package repository

import (
	"kino_backend/models"
)

type UsersRepository interface {
	GetUserPassword(e string) (models.User, error)
	CreateNewUser(u *models.RegisterProfile) (models.Profile, error)
	UpdateUserByID(id uint, u *models.RegisterProfile) error
	GetUserProfileByID(id uint) (models.FullProfile, error)
	GetUserProfileByNickname(nickname string) (models.Profile, error)
	CheckExistenceOfEmail(e string) (bool, error)
	CheckExistenceOfNickname(n string) (bool, error)
	GetCountOfUsers() (int, error)
	UploadAvatar(uID uint, path string) error
	DeleteAvatar(uID uint) error
}
