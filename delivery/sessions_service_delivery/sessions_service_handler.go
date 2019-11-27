package sessions_service_delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"kino_backend/CSRF"
	"kino_backend/db"
	"kino_backend/logger"
	"kino_backend/models"
	"kino_backend/session_microservice_client"
	"kino_backend/useCase"
	"kino_backend/utilits/middleware"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	manager     *session_microservice_client.SessionManager
	useCaseUser useCase.UsersUseCase
}

func NewHandler(info *session_microservice_client.SessionManager, useCaseUser useCase.UsersUseCase) *Handler {
	return &Handler{
		manager:     info,
		useCaseUser: useCaseUser,
	}
}

func readLoginInfo(r *http.Request, u *models.UserPassword) error {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, u)
	if err != nil {
		return models.ParseJSONError{err}
	}

	return nil
}

func (h *Handler) LoginUser(ctx context.Context, w http.ResponseWriter, userID uint) (bool, error) {
	sessionID := ""
	var successGlobal bool = false
	for {
		// create session, if collision ocquires, generate new sessionID
		var err error
		u := uuid.Must(uuid.NewV4(), err)
		sessionID = u.String()
		fmt.Println(sessionID)

		//sessionID = uuid.NewV4().String()
		success, err := h.manager.Create(sessionID, userID)
		successGlobal = success
		if err != nil {
			logger.Error(err)
			return success, err
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
	tokenExpiration := time.Now().Add(24 * time.Hour)
	csrfToken, _ := CSRF.Tokens.Create(string(userID), cookie.Value, tokenExpiration.Unix())
	w.Header().Set("X-CSRF-Token", csrfToken)

	http.SetCookie(w, &cookie)

	return successGlobal, nil
}

// @Summary Получить сессию
// @Description Получить сессию пользователя, если есть сессия, то она в куке session_id
// @ID get-session
// @Produce json
// @Success 200 {object} models.Session "Пользователь залогинен, успешно"
// @Failure 401 "Не залогинен"
// @Failure 500 "Ошибка в бд"
// @Router /session [GET]
func (h *Handler) getSession(w http.ResponseWriter, r *http.Request) {
	if r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		sID, err := json.Marshal(models.Session{SessionID: r.Context().Value(middleware.KeySessionID).(string)})
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(sID))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// @Summary Залогинить
// @Description Залогинить пользователя (создать сессию)
// @ID post-session
// @Accept json
// @Produce json
// @Param UserPassword body models.UserPassword true "Почта и пароль"
// @Success 200 {object} models.Session "Успешный вход / пользователь уже залогинен"
// @Failure 400 "Неверный формат JSON, невалидные данные"
// @Failure 422 "Неверная пара пользователь/пароль"
// @Failure 500 "Внутренняя ошибка"
// @Router /session [POST]
func (h *Handler) postLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		// user has already logged in
		return
	}
	Suc := models.Success{false}

	u := &models.UserPassword{}
	err := readLoginInfo(r, u)
	if err != nil {
		switch err.(type) {
		case models.ParseJSONError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	isValid := govalidator.IsEmail(u.Email)
	if !isValid {
		models.SendError(w, r, fmt.Errorf("Невалидная почта"), http.StatusBadRequest)
		return
	}
	//исправить !!!
	dbResponse, err := h.useCaseUser.GetUserPassword(u.Email)

	if err != nil {
		switch err.(type) {
		case db.UserNotFoundError:
			w.WriteHeader(http.StatusUnprocessableEntity)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	if u.Email == dbResponse.Email && u.Password == dbResponse.Password {
		suc, err := h.LoginUser(r.Context(), w, dbResponse.UserID)
		Suc = models.Success{suc}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json, err := json.Marshal(Suc)
		if err != nil {
			log.Println(err, "error msrashal")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, string(json))
		logger.Info("user with id %v and email %v logged in", dbResponse.UserID, dbResponse.Email)
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
}

// @Summary Разлогинить
// @Descriptiond Разлогинить пользователя (удалить сессию)
// @ID delete-session
// @Success 200 "Успешный выход / пользователь уже разлогинен"
// @Router /session [DELETE]
func (h *Handler) deleteSession(w http.ResponseWriter, r *http.Request) {
	if !r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		// user has already logged out
		return
	}
	err := h.manager.Delete(r.Context().Value(middleware.KeySessionID).(string))
	if err != nil { // but we continue
		logger.Error(err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Expires:  time.Now().AddDate(0, 0, -1),
		Secure:   true,
		HttpOnly: true,
	})
}

// @Summary Проверка авторизации
// @Description получить булевое значение статуса авторизации
// @ID authorized
// @Produce json
// @Success 200 true/false "Пользователь залогинен, успешно"
// @Failure 500 "Ошибка в бд"
// @Router /session [GET]
func (h *Handler) getOfAuthorized(w http.ResponseWriter, r *http.Request) {
	if r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		params := r.Context().Value(middleware.KeyIsAuthenticated).(bool)
		paramsJSON, err := json.Marshal(models.Authorization{Authorized: params})
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(paramsJSON))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
