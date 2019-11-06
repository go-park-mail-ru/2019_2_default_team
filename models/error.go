package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"kino_backend/logger"
)

type Error struct {
	Message string `json:"error" example:"Error: oops"`
}

type ParseJSONError struct {
	SomeError error
}

func (e ParseJSONError) Error() string {
	return fmt.Sprintf("error while parsing JSON: %v", e.SomeError)
}

func SendError(w http.ResponseWriter, r *http.Request, e error, status int) {
	m, err := json.Marshal(Error{Message: e.Error()})
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintln(w, string(m))
}