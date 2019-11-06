package users

import "fmt"

type UserNotFoundError struct {
	Field string
}


func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("no user with this %v found", e.Field)
}

type UserNotAuthError struct {
	Field string
}

func (e UserNotAuthError) Error() string {
	return fmt.Sprintf("no user auth %v found", e.Field)
}
