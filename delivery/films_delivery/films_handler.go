package films_delivery

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"io/ioutil"
	"kino_backend/logger"
	"kino_backend/models"
	"kino_backend/useCase"
	"kino_backend/utilits/errors"
	"kino_backend/utilits/middleware"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	useCase useCase.FilmsUseCase
}

func NewHandler(useCase useCase.FilmsUseCase) *Handler {
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

func readRegisterMovieSession(r *http.Request, p *models.RegisterMovieSession) error {
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

func readMovieSession(r *http.Request, p *models.MovieSession) error {
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

func readVote(r *http.Request, p *models.RegisterVote) error {
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

func SanitizeMe(film models.ProfileFilm) models.ProfileFilm {
	sanitizer := bluemonday.UGCPolicy()
	film.Description = sanitizer.Sanitize(film.Description)
	film.MainActor = sanitizer.Sanitize(film.Description)
	film.Director = sanitizer.Sanitize(film.Director)

	return film
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

func (h *Handler) getProfileFilm(w http.ResponseWriter, r *http.Request) {
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
	profile = SanitizeMe(profile)
	profileJSON, err := profile.MarshalJSON()
	//json, err := json.Marshal(profile)
	//if err != nil {
	//	log.Println(err, "in profileMethod")
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	fmt.Fprintln(w, string(profileJSON))

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

func (h *Handler) postSignupProfileFilm(w http.ResponseWriter, r *http.Request) {

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

// @Title Create new movie session
// @Summary  Создать новый сеанс фильма
// @ID post-profile
// @Accept json
// @Produce json
// @Param model.RegisterMovieSession
// @Success 200 "Создан сеанс
// @Failure 400 "Неверный формат JSON"
// @Failure 422 "При регистрации не все параметры"
// @Failure 500 "Ошибка в бд"
// @Router /createmoviesession [POST]

func (h *Handler) postCreateMovieSession(w http.ResponseWriter, r *http.Request) {

	//начала парсинга данных
	u := &models.RegisterMovieSession{}
	err := readRegisterMovieSession(r, u)
	if err != nil {
		switch err.(type) {
		case models.ParseJSONError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	//логика

	newMS, err := h.useCase.CreateNewMovieSession(r.Context(), u, 20)

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

	fmt.Print("New moviesession with id , time created", newMS.MsID, newMS.Date)
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

func (h *Handler) putEditFilmProfile(w http.ResponseWriter, r *http.Request) {
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
	err = h.useCase.PutFilm(r.Context(), filmInfo)

	//err = db.UpdateFilmByID(filmInfo.FilmID, filmInfo)
	//ответ с логики
	if err != nil {
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

func (h *Handler) getOneFilm(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err, "error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	params := &models.RequestProfileFilm{}
	params.ID = uint(ID)

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
	profile = SanitizeMe(profile)

	json, err := profile.MarshalJSON()
	if err != nil {
		log.Println(err, "in profileMethod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))

}

func (h *Handler) getAllFilms(w http.ResponseWriter, r *http.Request) {

	//конец парсинга данных можно передавать ctx,
	vals := r.URL.Query() // Returns a url.Values, which is a map[string][]string
	productTypes, ok := vals["title"]
	var pt string
	if ok {
		if len(productTypes) >= 1 {
			pt = productTypes[0]
		}
	}

	profile, err := h.useCase.GetAllFilms(r.Context())
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

	var finalFilms []models.ProfileFilm

	if pt != "" {
		for _, value := range profile {
			if strings.Contains(strings.ToLower(value.Title), strings.ToLower(pt)) {
				finalFilms = append(finalFilms, value)
			}
		}
	}

	if len(finalFilms) != 0 {
		json, err := json.Marshal(finalFilms)
		if err != nil {
			log.Println(err, "in profileMethod")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, string(json))
	} else {
		json, err := json.Marshal(profile)
		if err != nil {
			log.Println(err, "in profileMethod")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, string(json))
	}
}

func (h *Handler) getTimesForToday(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	movie_id := vars["movie_id"]

	u := &models.MovieSession{}
	var err error
	var intValue int
	intValue, err = strconv.Atoi(movie_id)

	if err != nil {
		switch err.(type) {
		case models.ParseJSONError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	u.MovieID = uint(intValue)

	//u := &models.MovieSession{}
	//err = readMovieSession(r, u)
	//if err != nil {
	//	switch err.(type) {
	//	case models.ParseJSONError:
	//		w.WriteHeader(http.StatusBadRequest)
	//	default:
	//		w.WriteHeader(http.StatusInternalServerError)
	//	}
	//	return
	//}

	times, err := h.useCase.GetMovieSessionsForToday(r.Context(), u.MovieID)
	if err != nil {
		fmt.Println("error")
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

	json, err := json.Marshal(times)
	if err != nil {
		log.Println(err, "in profileMethod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))

}

func (h *Handler) getSeatsByMSID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ms_id := vars["ms_id"]

	u := &models.MovieSession{}
	var err error
	var intValue int
	intValue, err = strconv.Atoi(ms_id)

	if err != nil {
		switch err.(type) {
		case models.ParseJSONError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	u.MsID = uint(intValue)

	//u := &models.MovieSession{}
	//err := readMovieSession(r, u)
	//if err != nil {
	//	switch err.(type) {
	//	case models.ParseJSONError:
	//		w.WriteHeader(http.StatusBadRequest)
	//	default:
	//		w.WriteHeader(http.StatusInternalServerError)
	//	}
	//	return
	//}

	seats, err := h.useCase.GetSeatsByMSID(r.Context(), u.MsID)
	if err != nil {
		fmt.Println("error")
		switch err.(type) {
		case errors.MSNotFoundError:
			w.WriteHeader(http.StatusNotFound)
			return
		default:
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	json, err := json.Marshal(seats)
	if err != nil {
		log.Println(err, "in profileMethod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))

}

func (h *Handler) getIsVoted(w http.ResponseWriter, r *http.Request) {

	u := &models.RegisterVote{}
	err := readVote(r, u)
	if err != nil {
		switch err.(type) {
		case models.ParseJSONError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	result, err := h.useCase.CheckIsVoted(r.Context(), u)
	if err != nil {
		fmt.Println("error")
		switch err.(type) {
		case errors.MSNotFoundError:
			w.WriteHeader(http.StatusNotFound)
			return
		default:
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	json, err := json.Marshal(result)
	if err != nil {
		log.Println(err, "in profileMethod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))

}

func (h *Handler) postVote(w http.ResponseWriter, r *http.Request) {

	if !r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	u := &models.RegisterVote{}
	err := readVote(r, u)
	if err != nil {
		switch err.(type) {
		case models.ParseJSONError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	id_user := r.Context().Value(middleware.KeyUserID).(uint)

	u.UserID = id_user

	result, err := h.useCase.CheckIsVoted(r.Context(), u)
	if err != nil {
		switch err.(type) {
		case errors.MSNotFoundError:
			w.WriteHeader(http.StatusNotFound)
			return
		default:
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if !result {
		newVote, err := h.useCase.Vote(r.Context(), u)

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

		fmt.Print("User id , was registered for film ", newVote.UserID, newVote.MovieID)

		w.WriteHeader(http.StatusCreated)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getFilmsForToday(w http.ResponseWriter, r *http.Request) {

	films, err := h.useCase.GetFilmsForToday(r.Context())
	if err != nil {
		fmt.Println("error")
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

	json, err := json.Marshal(films)
	if err != nil {
		log.Println(err, "in profileMethod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))

}

func (h *Handler) getFilmsForSoon(w http.ResponseWriter, r *http.Request) {

	films, err := h.useCase.GetFilmsForSoon(r.Context())
	if err != nil {
		fmt.Println("error")
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

	json, err := json.Marshal(films)
	if err != nil {
		log.Println(err, "in profileMethod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))

}

func (h *Handler) getRecommendedFilms(w http.ResponseWriter, r *http.Request) {

	//конец парсинга данных можно передавать ctx,
	vals := r.URL.Query() // Returns a url.Values, which is a map[string][]string
	productTypes, ok := vals["genre"]
	var pt string
	if ok {
		if len(productTypes) >= 1 {
			pt = productTypes[0]
		}
	}

	films, err := h.useCase.GetRecommendedFilms(pt, r.Context())
	if err != nil {
		fmt.Println("error")
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

	json, err := json.Marshal(films)
	if err != nil {
		log.Println(err, "in profileMethod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))
}
