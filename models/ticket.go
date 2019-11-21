package models

type RegisterTicket struct {
	UserID uint `json:"user_id" db:"profile_id"`
	FilmID uint `json:"film_id" db:"movie_session_id"`
	SeatID uint `json:"seat_id" db:"seat_id"`
	Price  uint `json:"price" db:"price"`
}

type Ticket struct {
	RegisterTicket
	TicketID uint `json:"ticket_id" db:"ticket_id"`
}

type RequestTicket struct {
	TicketID uint `json:"ticket_id" db:"ticket_id"`
}

type Seat struct {
	SeatID     uint   `json:"seat_id" db:"seat_id"`
	HallName   string `json:"hall_name" db:"hall_name"`
	Row        int    `json:"row" db:"row"`
	SeatNumber int    `json:"seat_number" db:"seat_number"`
}
