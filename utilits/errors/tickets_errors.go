package errors

import (
	"errors"
	"fmt"
)

var (
	ErrSessionNotFound = errors.New("no session in database")
	ErrNotNullConstraintViolation = errors.New("not null constraint violation")
	ErrUniqueConstraintViolation  = errors.New("unique constraint violation")
)

type TicketNotFoundError struct {
	Field string
}


func (e TicketNotFoundError) Error() string {
	return fmt.Sprintf("no ticket with this %v found", e.Field)
}


