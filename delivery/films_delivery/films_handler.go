package films_delivery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kino_backend/logger"
	"kino_backend/models"
	"kino_backend/useCase"
	"kino_backend/utilits/errors"
	"kino_backend/utilits/middleware"
	"log"
	"net/http"
)

type Handler struct{
	useCase useCase.FilmsUseCase
}

func NewHandler(useCase useCase.FilmsUseCase) *Handler{
	return &Handler{
		useCase: useCase,
	}
}

func readProfileFilm(r *http.Request, p *models.ProfileFilm) error {
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

func readRegisterProfileFilm(r *http.Request, p *models.RegisterProfileFilm) error {
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
// @Summary Получить профиль фильма по ID или названию Title
// @ID get-profilefilm
// @Produce json
// @Param id query int false "ID"
// @Success 200 {object} models.ProfileFilm "Film найден, успешно"
// @Failure 400 "Неправильный запрос"
// @Failure 404 "Не найдено"
// @Failure 500 "Ошибка в бд"
// @Router /profilefilm [GET]

//http methods tickets_handler.go

func (h *Handler) getProfileFilm(w http.ResponseWriter, r *http.Request){
	//этап парсинга данных
	params := &models.RequestProfileFilm{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(params)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//конец парсинга данных можно передавать ctx,

	profile, err := h.useCase.GetFilm(r.Context(), params)
	if err != nil {
		switch err.(type) {
		case errors.FilmNotFoundError:
			w.WriteHeader(http.StatusNotFound)
			return
		default:
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(profile)
	if err != nil {
		log.Println(err, "in profileMethod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))

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

func (h *Handler) postSignupProfileFilm(w http.ResponseWriter, r *http.Request){

	//начала парсинга данных
	u := &models.RegisterProfileFilm{}
	err := readRegisterProfileFilm(r, u)
	if err != nil {
		switch err.(type) {
		case models.ParseJSONError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	if u.Title == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//логика

	newF, err := h.useCase.PostFilmUse(r.Context(), u)

	if err != nil {
		if err == errors.ErrUniqueConstraintViolation ||
			err == errors.ErrNotNullConstraintViolation {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//fmt.Fprintln(w, http.StatusOK)
	w.WriteHeader(http.StatusOK)

	fmt.Print("New film with id , title created", newF.FilmID, newF.Title)
	//logger.Infof("New film with id %v, title %v created", newF.FilmID, newF.Title)
	//очень плохо в тесте реагирует мока
}


// @Title Изменить профиль
// @Summary Изменить профиль, должен быть залогинен
// @ID put-profile
// @Accept json
// @Produce json
// @Param Profile body models.RegisterProfile true "Новые никнейм, и/или почта, и/или пароль"
// @Success 200 "Пользователь найден, успешно изменены данные"
// @Failure 400 "Неверный формат JSON"
// @Failure 401 "Не залогинен"
// @Failure 403 {object} models.ProfileErrorList "Занята почта или ник, пароль не удовлетворяет правилам безопасности, другие ошибки"
// @Failure 500 "Ошибка в бд"
// @Router /profile [PUT]

func (h *Handler) putEditFilmProfile(w http.ResponseWriter, r *http.Request){
	//начало парсинга данных
	if !r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	filmInfo := &models.ProfileFilm{}
	err := readProfileFilm(r, filmInfo)

	if err != nil {
		switch err.(type) {
		case models.ParseJSONError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	uid := r.Context().Value(middleware.KeyUserID).(uint)
	if uid != filmInfo.AdminID{
		log.Println("no access")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//логика
	err = h.useCase.PutFilm(r.Context(), filmInfo)

	//err = db.UpdateFilmByID(filmInfo.FilmID, filmInfo)
	//ответ с логики
	if err != nil{
		switch err.(type) {
		case errors.FilmNotFoundError:
			w.WriteHeader(http.StatusNotFound)
		default:
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	log.Println("Film with id", filmInfo.FilmID, "changed to", filmInfo.Title, filmInfo.Description)
}
