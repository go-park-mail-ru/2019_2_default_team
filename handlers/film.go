package handlers

import (
	"2019_2_default_team/db"
	"2019_2_default_team/logger"
	"2019_2_default_team/middleware"
	"2019_2_default_team/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//http methods/ delivery/handler.go

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

//http methods/ delivery/register.go

func (api *MyHandler) ProfileFilmHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getProfileFilm(w, r)
	case http.MethodPost:
		postSignupProfileFilm(w, r)
	case http.MethodPut:
		putEditFilmProfile(w, r)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
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

//http methods handler.go

func getProfileFilm(w http.ResponseWriter, r *http.Request) {
	//этап парсинга данных
	params := &models.RequestProfileFilm{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(params)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//конец парсинга данных можно передавать ctx,
	if params.ID != 0 {
		profile, err := db.GetFilmProfileByID(params.ID)
		if err != nil {
			switch err.(type) {
			case db.FilmNotFoundError:
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
	} else if params.Title != "" {
		profile, err := db.GetFilmProfileByTitle(params.Title)
		if err != nil {
			switch err.(type) {
			case db.FilmNotFoundError:
				w.WriteHeader(http.StatusNotFound)
				return
			default:
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		//этап отправки данных

		w.Header().Set("Content-Type", "application/json")
		json, err := json.Marshal(profile)
		if err != nil {
			log.Println(err, "in profileMethod")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, string(json))
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

func postSignupProfileFilm(w http.ResponseWriter, r *http.Request) {

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
	newF, err := db.CreateNewFilm(u)
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

	logger.Infof("New film with id %v, title %v created", newF.FilmID, newF.Title)

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

func putEditFilmProfile(w http.ResponseWriter, r *http.Request) {
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
	if uid != filmInfo.AdminID {
		log.Println("no access")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//логика

	err = db.UpdateFilmByID(filmInfo.FilmID, filmInfo)
	//ответ с логики
	if err != nil {
		switch err.(type) {
		case db.FilmNotFoundError:
			w.WriteHeader(http.StatusNotFound)
		default:
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	log.Println("Film with id", filmInfo.FilmID, "changed to", filmInfo.Title, filmInfo.Description)

}
