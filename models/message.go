package models

import "time"

type Message struct {
	ID       uint      `json:"id,omitempty" db:"message_id"`
	Author   *uint     `json:"author,omitempty" db:"author_id"`
	To       *uint     `json:"to,omitempty" db:"to_user"`
	Created  time.Time `json:"created"`
	IsEdited bool      `json:"is_edited" db:"is_edited"`
	Message  string    `json:"message" db:"message_text"`
}

type Messages struct {
	Msgs *[]Message `json:"messages"`
}

type USERTOSUP struct {
	UTSID  uint `json:"uts_id" db:"uts_id"`
	USERID uint `json:"user_id" db:"user_id"`
	SUPID  uint `json:"suo_id" db:"suo_id"`
}

type OneMessage struct {
	Message string    `json:"message" db:"message"`
	Time    time.Time `json:"time" db:"time"`
}

type Status struct {
	Status string `json:"status" db:"status"`
}

type SupportDialog struct {
	dialod_id int `json:"uts_id" db:"uts_id"`
}
