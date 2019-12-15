package tickets_delivery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kino_backend/db"
	"kino_backend/logger"
	"kino_backend/models"
	"kino_backend/useCase"
	"kino_backend/utilits/middleware"
	"log"
	"net/http"
)

type Handler struct {
	useCase useCase.TicketsUseCase
}

func NewHandler(useCase useCase.TicketsUseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func readRegisterProfileTicket(r *http.Request, p *models.RegisterTicket) error {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, p)
	if err != nil {
		return models.ParseJSONError{err}
	}

	return nil
}

func readProfileTicket(r *http.Request, p *models.Ticket) error {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, p)
	if err != nil {
		return models.ParseJSONError{err}
	}

	return nil
}

// @Title Получить профиль
// @Summary Получить профиль билета по ID
// @Produce json
// @Param id query int false "ID"
// @Success 200 {object} models.Ticket "Ticket найден, успешно"
// @Failure 400 "Неправильный запрос"
// @Failure 404 "Не найдено"
// @Failure 500 "Ошибка в бд"
// @Router /profileticket [GET]

func (h *Handler) getProfileTicket(w http.ResponseWriter, r *http.Request) {
	params := &models.RequestTicket{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(params)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//logic

	if params.TicketID != 0 {

		//logic
		profile, err := h.useCase.GetTicket(r.Context(), params)
		//profile, err := db.GetTicketProfileByID(params.TicketID)
		if err != nil {
			switch err.(type) {
			case db.TicketNotFoundError:
				w.WriteHeader(http.StatusNotFound)
				return
			default:
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json, err := profile.MarshalJSON()
		if err != nil {
			log.Println(err, "in profileMethod")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, string(json))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// @Title Зарегистрироваться и залогиниться по новому профилю
// @Summary Зарегистрировать по никнейму, почте и паролю и автоматически залогинить
// @ID post-profile
// @Accept json
// @Produce json
// @Param Profile body models.RegisterProfile true "Никнейм, почта и пароль"
// @Success 200 "Пользователь зарегистрирован и залогинен успешно"
// @Failure 400 "Неверный формат JSON"
// @Failure 403 {object} models.ProfileErrorList "Занята почта или ник, пароль не удовлетворяет правилам безопасности, другие ошибки"
// @Failure 422 "При регистрации не все параметры"
// @Failure 500 "Ошибка в бд"
// @Router /profile [POST]

func (h *Handler) postBuyTicket(w http.ResponseWriter, r *http.Request) {
	if !r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	id_user := r.Context().Value(middleware.KeyUserID).(uint)
	u := &models.RegisterTicket{}
	err := readRegisterProfileTicket(r, u)
	if err != nil {
		switch err.(type) {
		case models.ParseJSONError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	//билетик в подарок?
	//uid := r.Context().Value(middleware.KeyUserID).(uint)
	//if uid != u.UserID{
	//	log.Println("no access")
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}

	//logic
	u.UserID = id_user

	result, err := h.useCase.CheckTicket(r.Context(), u)

	if err != nil {
		if err == db.ErrUniqueConstraintViolation ||
			err == db.ErrNotNullConstraintViolation {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if result {
		w.WriteHeader(http.StatusAlreadyReported)
		fmt.Print("Ticket was bought")
		return
	}

	newT, err := h.useCase.PostTicket(r.Context(), u)
	//newT, err := db.CreateNewTicket(u)
	if err != nil {
		if err == db.ErrUniqueConstraintViolation ||
			err == db.ErrNotNullConstraintViolation {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Print("New ticket", newT.TicketID)
	//logger.Infof("New ticket with id %v ", newT.TicketID)

}
