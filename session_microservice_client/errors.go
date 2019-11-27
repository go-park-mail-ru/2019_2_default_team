package session_microservice_client

import "errors"

var (
	ErrConnRefused = errors.New("no session in database")
	ErrKeyNotFound = errors.New("key not found")
)
