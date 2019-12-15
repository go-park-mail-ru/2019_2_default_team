package models

type RegisterTicket struct {
	UserID uint `json:"user_id" db:"profile_id"`
	MSID   uint `json:"ms_id" db:"movie_session_id"`
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
	SeatID         uint   `json:"seat_id" db:"seat_id"`
	HallName       string `json:"hall_name, omitempty" db:"hall_name"`
	MovieSessionID int    `json:"movie_session_id" db:"movie_session_id"`
	IsTaken        bool   `json:"is_taken" db:"is_taken"`
	Row            int    `json:"row" db:"row"`
	SeatNumber     int    `json:"seat_number" db:"seat_number"`
}
