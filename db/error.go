package db

import (
	"errors"
	"fmt"
)

var (
	ErrSessionNotFound = errors.New("no session in database")
	ErrNotNullConstraintViolation = errors.New("not null constraint violation")
	ErrUniqueConstraintViolation  = errors.New("unique constraint violation")
)

type UserNotFoundError struct {
	Field string
}

type UserNotAuthError struct{
	Field string
}

type FilmNotFoundError struct {
	Field string
}

type TicketNotFoundError struct {
	Field string
}

func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("no user with this %v found", e.Field)
}

func (e FilmNotFoundError) Error() string {
	return fmt.Sprintf("no film with this %v found", e.Field)
}

func (e TicketNotFoundError) Error() string {
	return fmt.Sprintf("no ticket with this %v found", e.Field)
}

func (e UserNotAuthError) Error() string {
	return fmt.Sprintf("no auth with this %v found", e.Field)
}

