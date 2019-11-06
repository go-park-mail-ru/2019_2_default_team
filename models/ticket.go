package models


type RegisterTicket struct{
	UserID uint `json:"user_id" db:"user_id"`
	FilmID uint `json:"film_id" db:"film_id"`
}

type Ticket struct{
	RegisterTicket
	TicketID uint `json:"ticket_id" db:"ticket_id"`
}

type RequestTicket struct{
	TicketID uint `json:"ticket_id" db:"ticket_id"`
}