package repository

import (
	"kino_backend/models"
)

type FilmsRepository interface {
	CreateNewFilm(u *models.RegisterProfileFilm) (models.ProfileFilm, error)
	UpdateFilmByID(id uint, u *models.ProfileFilm) error
	GetFilmProfileByID(id uint) (models.ProfileFilm, error)
	GetFilmProfileByIDSQL(id uint) (models.ProfileFilm, error)
	GetFilmProfileByTitle(title string) (models.ProfileFilm, error)
	CheckExistenceOfTitle(n string) (bool, error)
	GetCountOfFilms() (int, error)
	UploadAvatarFilm(uID uint, path string) error
	DeleteAvatarFilm(uID uint) error
	GetAllFilms() ([]models.ProfileFilm, error)
}
