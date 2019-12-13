package errors

import (
	"errors"
	"fmt"
)

var (
	ErrSessionNotFound            = errors.New("no session in database")
	ErrNotNullConstraintViolation = errors.New("not null constraint violation")
	ErrUniqueConstraintViolation  = errors.New("unique constraint violation")
	ErrMovSessionNotFound         = errors.New("no session in database")
)

type TicketNotFoundError struct {
	Field string
}

type MSNotFoundError struct {
	Field string
}

func (e MSNotFoundError) Error() string {
	return fmt.Sprintf("no ms with this %v found", e.Field)
}

func (e TicketNotFoundError) Error() string {
	return fmt.Sprintf("no ticket with this %v found", e.Field)
}
