package models

type ProfileFilm struct {
	Film
	Title string  `json:"title" example:"Joker"`
	Description string  `json:"description" example:"Absolutely madness"`
	Avatar   *string `json:"avatar,omitempty"`
	Director string   `json:"director" example:"Todd Philips"`
	MainActor string   `json:"mainactor" example:"Phoenix"`
	AdminID uint `json:"admin_id" db:"admin_id"`
}

type RegisterProfileFilm struct {
	Title string `json:"title" example:"Joker"`
	Description string  `json:"description" example:"Absolutely madness"`
	Avatar   *string `json:"avatar,omitempty"`
	Director string   `json:"director" example:"Todd Philips"`
	MainActor string   `json:"mainactor" example:"Phoenix"`
	AdminID uint `json:"admin_id" db:"admin_id"`
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
	ID       uint    `json:"reqidfilm"`
	Title string   `json:"reqtitle"`
}