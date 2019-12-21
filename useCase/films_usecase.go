package useCase

import (
	"context"
	"kino_backend/models"
	"kino_backend/repository"
	"kino_backend/utilits/middleware"
)

type FilmsUseCase interface {
	GetFilm(ctx context.Context, params *models.RequestProfileFilm) (models.ProfileFilmWithVote, error)
	PostFilmUse(ctx context.Context, u *models.RegisterProfileFilm) (models.ProfileFilm, error)
	PutFilm(ctx context.Context, filmInfo *models.ProfileFilm) error
	GetAllFilms(ctx context.Context) ([]models.ProfileFilm, error)
	CreateNewMovieSession(ctx context.Context, u *models.RegisterMovieSession, seatsNumber int) (models.MovieSession, error)
	GetMovieSessionsForToday(ctx context.Context, movie_id uint) ([]models.RequestFilmTimes, error)
	GetSeatsByMSID(ctx context.Context, movie_session_id uint) ([]models.Seat, error)
	CheckIsVoted(ctx context.Context, u *models.RegisterVote) (bool, error)
	Vote(ctx context.Context, u *models.RegisterVote) (models.Vote, error)
	GetFilmsForToday(ctx context.Context) ([]models.ProfileFilm, error)
	GetFilmsForSoon(ctx context.Context) ([]models.ProfileFilm, error)
	GetRecommendedFilms(wantedGenre string, ctx context.Context) ([]models.ProfileFilm, error)
}

type filmUseCase struct {
	filmRepo repository.FilmRepository
}

func NewFilmUseCase(filmRepo repository.FilmRepository) *filmUseCase {
	return &filmUseCase{
		filmRepo: filmRepo,
	}
}

func (f filmUseCase) GetMovieSessionsForToday(ctx context.Context, movie_id uint) ([]models.RequestFilmTimes, error) {
	newTimes, err := f.filmRepo.GetMovieSessionsForToday(movie_id)

	if err != nil {
		return []models.RequestFilmTimes{}, err
	}

	return newTimes, err
}

func (f filmUseCase) GetFilmsForToday(ctx context.Context) ([]models.ProfileFilm, error) {
	newFilms, err := f.filmRepo.GetFilmsForToday()

	if err != nil {
		return []models.ProfileFilm{}, err
	}

	return newFilms, err
}

func (f filmUseCase) GetRecommendedFilms(wantedGenre string, ctx context.Context) ([]models.ProfileFilm, error) {
	var err error
	var profile []models.ProfileFilm

	profile, err = f.filmRepo.GetRecommendedFilms(wantedGenre)

	if err != nil {
		return []models.ProfileFilm{}, err
	}

	return profile, nil
}

func (f filmUseCase) GetFilmsForSoon(ctx context.Context) ([]models.ProfileFilm, error) {
	newFilms, err := f.filmRepo.GetFilmsForSoon()

	if err != nil {
		return []models.ProfileFilm{}, err
	}

	return newFilms, err
}

func (f filmUseCase) GetSeatsByMSID(ctx context.Context, movie_session_id uint) ([]models.Seat, error) {
	newSeats, err := f.filmRepo.GetSeatsByMSID(movie_session_id)

	if err != nil {
		return []models.Seat{}, err
	}

	return newSeats, err
}

func (f filmUseCase) Vote(ctx context.Context, u *models.RegisterVote) (models.Vote, error) {
	result, err := f.filmRepo.VoteForFilm(u)

	if err != nil {
		return models.Vote{}, err
	}

	return result, err
}

func (f filmUseCase) CheckIsVoted(ctx context.Context, u *models.RegisterVote) (bool, error) {
	result, err := f.filmRepo.IsVoted(u)

	if err != nil {
		return false, err
	}

	return result, err
}

func (f filmUseCase) CreateNewMovieSession(ctx context.Context, u *models.RegisterMovieSession, seatsNumber int) (models.MovieSession, error) {
	newMS, err := f.filmRepo.CreateNewMovieSession(u, seatsNumber)

	if err != nil {
		return models.MovieSession{}, err
	}

	return newMS, err
}

func (f filmUseCase) GetFilm(ctx context.Context, params *models.RequestProfileFilm) (models.ProfileFilmWithVote, error) {
	var err error
	var profile models.ProfileFilm
	var profileVote models.ProfileFilmWithVote
	var u models.RegisterVote
	var voted bool

	if params.ID != 0 {
		profile, err = f.filmRepo.GetFilmProfileByID(params.ID)
		if !ctx.Value(middleware.KeyIsAuthenticated).(bool) {
			voted = false
		} else {
			u.UserID = ctx.Value(middleware.KeyUserID).(uint)
			u.MovieID = profile.FilmID
			res, err := f.CheckIsVoted(ctx, &u)
			if err != nil {
				return models.ProfileFilmWithVote{}, err
			}
			voted = res
		}

		profileVote.ProfileFilm = profile
		profileVote.IsVoted = voted

		if err != nil {
			return models.ProfileFilmWithVote{}, err
		}
		return profileVote, nil
	} else if params.Title != "" {
		profile, err = f.filmRepo.GetFilmProfileByTitle(params.Title)

		if !ctx.Value(middleware.KeyIsAuthenticated).(bool) {
			voted = false
		} else {
			u.UserID = ctx.Value(middleware.KeyUserID).(uint)
			u.MovieID = profile.FilmID
			res, err := f.CheckIsVoted(ctx, &u)
			if err != nil {
				return models.ProfileFilmWithVote{}, err
			}
			voted = res
		}

		profileVote.ProfileFilm = profile
		profileVote.IsVoted = voted
		if err != nil {
			return models.ProfileFilmWithVote{}, err
		}

		return profileVote, nil
	}
	return profileVote, nil
}

func (f filmUseCase) PostFilmUse(ctx context.Context, u *models.RegisterProfileFilm) (models.ProfileFilm, error) {
	newF, err := f.filmRepo.CreateNewFilm(u)

	if err != nil {
		return models.ProfileFilm{}, err
	}

	return newF, err
}

func (f filmUseCase) PutFilm(ctx context.Context, filmInfo *models.ProfileFilm) error {
	err := f.filmRepo.UpdateFilmByID(filmInfo.FilmID, filmInfo)
	return err
}

func (f filmUseCase) GetAllFilms(ctx context.Context) ([]models.ProfileFilm, error) {
	var err error
	var profile []models.ProfileFilm

	profile, err = f.filmRepo.GetAllFilms()

	if err != nil {
		return []models.ProfileFilm{}, err
	}

	return profile, nil
}
