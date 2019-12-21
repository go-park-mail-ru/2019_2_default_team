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
	CreateNewMovieSession(u *models.RegisterMovieSession, seatsNumber int) (models.MovieSession, error)
	GetMovieSessionsForToday(movie_id int) ([]models.RequestFilmTimes, error)
	GetSeatsByMSID(movie_session_id uint) ([]models.Seat, error)
	IsVoted(u *models.RegisterVote) (bool, error)
	GetFilmsForToday() ([]models.ProfileFilm, error)
	GetFilmsForSoon() ([]models.ProfileFilm, error)
	GetRecommendedFilms(wantedGenre string) ([]models.ProfileFilm, error)
}
