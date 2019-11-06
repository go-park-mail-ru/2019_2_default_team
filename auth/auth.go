package auth

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
	"kino_backend/sessions"
	"kino_backend/logger"
)

func LoginUser(w http.ResponseWriter, userID uint) error {
	sessionID := ""
	for {
		// create session, if collision ocquires, generate new sessionID
		var err error
		u, _ := uuid.NewV4()
		sessionID = u.String()

		//sessionID = uuid.NewV4().String()
		success, err := sessions.Create(sessionID, userID)
		if err != nil {
			logger.Error(err)
			return err
		}
		if success {
			break
		}
	}

	cookie := http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	return nil
}
