package models

import "time"

type ProfileFilm struct {
	Film
	Title       string  `json:"title" example:"Joker"`
	Description string  `json:"description" example:"Absolutely madness"`
	Avatar      *string `json:"avatar,omitempty"`
	Director    string  `json:"director"  db:"director" example:"Todd Philips"`
	MainActor   string  `json:"mainactor"  db:"actors" example:"Phoenix"`
	AdminID     uint    `json:"admin_id" db:"admin_id"`
	Genre       string  `json:"genre" db:"genre"`
	Length      int     `json:"length" db:"length"`
	Rating      int     `json:"rating" db:"rating"`
	Production  string  `json:"production" db:"production"`
	Poster      string  `json:"poster" db:"poster"`
	PosterPopup string  `json:"poster_popup" db:"poster_popup"`
	Trailer     string  `json:"trailer" db:"trailer"`
	Year        int     `json:"year" db:"year"`
}

type RegisterProfileFilm struct {
	Title       string  `json:"title" example:"Joker"`
	Description string  `json:"description" example:"Absolutely madness"`
	Avatar      *string `json:"avatar,omitempty"`
	Director    string  `json:"director" example:"Todd Philips"`
	MainActor   string  `json:"mainactor" example:"Phoenix"`
	AdminID     uint    `json:"admin_id" db:"admin_id"`
	Genre       string  `json:"genre" db:"genre"`
	Length      int     `json:"length" db:"length"`
	Production  string  `json:"production" db:"production"`
	Year        int     `json:"year" db:"year"`
	//UserPassword
}

type Film struct {
	FilmID uint `json:"id" db:"film_id"`
	//UserPassword
}

//type UserPassword struct {
//	Email    string `json:"email" example:"email@email.com" valid:"required~Почта не может быть пустой,email~Невалидная почта"`
//	Password string `json:"password,omitempty" example:"password" valid:"stringlength(8|32)~Пароль должен быть не менее 8 символов и не более 32 символов"`
//}

type ProfileFilmError struct {
	Field string `json:"field" example:"title"`
	Text  string `json:"text" example:"Этот фильм уже занят"`
}

type ProfileFilmErrorList struct {
	Errors []ProfileError `json:"error"`
}

type RequestProfileFilm struct {
	ID    uint   `json:"reqidfilm"`
	Title string `json:"reqtitle"`
}

type RegisterMovieSession struct {
	HallName string    `json:"hall_name" db:"hall_name"`
	MovieID  uint      `json:"movie_id" db:"movie_id"`
	Date     time.Time `json:"start_datetime" db:"start_datetime"`
	Type     string    `json:"type" db:"type"`
}

type MovieSession struct {
	MsID     uint      `json:"ms_id" db:"ms_id"`
	HallName string    `json:"hall_name" db:"hall_name"`
	MovieID  uint      `json:"movie_id" db:"movie_id"`
	Date     time.Time `json:"start_datetime" db:"start_datetime"`
	Type     string    `json:"type" db:"type"`
}

type RequestFilmTimes struct {
	MovieSessionID uint      `json:"ms_id" db:"ms_id"`
	Date           time.Time `json:"start_datetime" db:"start_datetime"`
	Hall           string    `json:"hall_name" db:"hall_name"`
}

type Vote struct {
	VoteID  uint `json:"vote_id" db:"vote_id"`
	MovieID uint `json:"film_id" db:"film_id"`
	UserID  uint `json:"user_id" db:"user_id"`
}

type RegisterVote struct {
	MovieID uint `json:"film_id" db:"film_id"`
	UserID  uint `json:"user_id" db:"user_id"`
}
