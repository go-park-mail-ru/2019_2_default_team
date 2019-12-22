package models

type Profile struct {
	User
	Nickname  string  `json:"nickname" example:"Nick"`
	Avatar    *string `json:"avatar,omitempty"`
	FirstName string  `json:"first_name" example:"Nick" db:"first_name"`
	LastName  string  `json:"last_name" example:"Nicker" db:"last_name"`
}

type FullProfile struct {
	User
	Nickname  string          `json:"nickname" example:"Nick"`
	Avatar    *string         `json:"avatar,omitempty"`
	FirstName string          `json:"first_name" example:"Nick" db:"first_name"`
	LastName  string          `json:"last_name" example:"Nicker" db:"last_name"`
	Tickets   []TicketProfile `json:"tickets"`
}

type RegisterProfile struct {
	Nickname string `json:"nickname" example:"Nick"`
	UserPassword
	FirstName string `json:"first_name" example:"Nick"`
	LastName  string `json:"last_name" example:"Nicker"`
}

type User struct {
	UserID uint `json:"id" db:"user_id"`
	UserPassword
}

type UserPassword struct {
	Email    string `json:"email" example:"email@email.com" valid:"required~Почта не может быть пустой,email~invalid_email"`
	Password string `json:"password,omitempty" example:"password" valid:"stringlength(8|32)~error_length"`
}

type ProfileError struct {
	Field string `json:"field" example:"nickname"`
	Text  string `json:"text" example:"Этот никнейм уже занят"`
}

type ProfileErrorList struct {
	Errors []ProfileError `json:"error"`
}

type RequestProfile struct {
	ID       uint   `json:"reqid"`
	Nickname string `json:"reqnick"`
}

type SessionCheck struct {
	Username string `json:"username"`
}
