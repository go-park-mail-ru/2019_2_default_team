package delivery

import (
	"2019_2_default_team/auth"
	"2019_2_default_team/db"
	"2019_2_default_team/logger"
	"2019_2_default_team/middleware"
	"2019_2_default_team/models"
	"2019_2_default_team/users"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
)

type Handler struct {
	useCase users.UseCase
}

func NewHandler(useCase users.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func readProfile(r *http.Request, p *models.RegisterProfile) error {
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

func validateNickname(s string) ([]models.ProfileError, error) {
	var errors []models.ProfileError
	fmt.Println(s)
	isValid := govalidator.StringLength(s, "4", "32")
	if !isValid {
		fmt.Println("inside", "isvalid", isValid)
		errors = append(errors, models.ProfileError{
			Field: "nickname",
			Text:  "Никнейм должен быть не менее 4 символов и не более 32 символов",
		})
		return errors, nil
	}

	fmt.Println("outside", "isvalid", isValid)
	exists, err := db.CheckExistenceOfNickname(s)
	if err != nil {
		fmt.Println("exists", exists)
		logger.Error(err)
		fmt.Println("errors", errors, "err", err)
		return errors, err
	}
	if exists {
		errors = append(errors, models.ProfileError{
			Field: "nickname",
			Text:  "Этот никнейм уже занят",
		})
	}
	fmt.Println("outside2", "is  ", errors)

	return errors, nil
}

func validateEmail(s string) ([]models.ProfileError, error) {
	var errors []models.ProfileError

	isValid := govalidator.IsEmail(s)
	if !isValid {
		errors = append(errors, models.ProfileError{
			Field: "email",
			Text:  "Невалидная почта",
		})
		return errors, nil
	}

	exists, err := db.CheckExistenceOfEmail(s)
	if err != nil {
		logger.Error(err)
		return errors, err
	}
	if exists {
		errors = append(errors, models.ProfileError{
			Field: "email",
			Text:  "Данная почта уже занята",
		})
	}

	return errors, nil
}

func validatePassword(s string) []models.ProfileError {
	var errors []models.ProfileError

	isValid := govalidator.StringLength(s, "8", "32")
	if !isValid {
		errors = append(errors, models.ProfileError{
			Field: "password",
			Text:  "Пароль должен быть не менее 8 символов и не более 32 символов",
		})
	}

	return errors
}

func validateFields(u *models.RegisterProfile) ([]models.ProfileError, error) {
	var errors []models.ProfileError

	valErrors, dbErr := validateNickname(u.Nickname)
	if dbErr != nil {
		fmt.Println("errnick")
		return []models.ProfileError{}, dbErr
	}
	errors = append(errors, valErrors...)

	valErrors, dbErr = validateEmail(u.Email)
	if dbErr != nil {
		fmt.Println("erremail")
		return []models.ProfileError{}, dbErr
	}
	errors = append(errors, valErrors...)
	errors = append(errors, validatePassword(u.Password)...)

	return errors, nil
}

// @Title Получить профиль
// @Summary Получить профиль пользователя по ID, email или из сессии
// @ID get-profile
// @Produce json
// @Param id query int false "ID"
// @Param nickname query string false "Никнейм"
// @Success 200 {object} models.Profile "Пользователь найден, успешно"
// @Failure 400 "Неправильный запрос"
// @Failure 401 "Не залогинен"
// @Failure 404 "Не найдено"
// @Failure 500 "Ошибка в бд"
// @Router /profile [GET]

func (h *Handler) getProfile(w http.ResponseWriter, r *http.Request) {
	//data parse
	params := &models.RequestProfile{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(params)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//end data parse
	var id uint

	if r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		id = r.Context().Value(middleware.KeyUserID).(uint)
	} else {
		id = 0
	}

	profile, err := h.useCase.GetUser(r.Context(), params, r.Context().Value(middleware.KeyIsAuthenticated).(bool), id)

	if err != nil {
		switch err.(type) {
		case db.UserNotFoundError:
			w.WriteHeader(http.StatusNotFound)
			return
		case users.UserNotAuthError:
			w.WriteHeader(http.StatusUnauthorized)
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

	//if params.ID != 0 {
	//	profile, err := db.GetUserProfileByID(params.ID)
	//	if err != nil {
	//		switch err.(type) {
	//		case db.UserNotFoundError:
	//			w.WriteHeader(http.StatusNotFound)
	//			return
	//		default:
	//			log.Println(err)
	//			w.WriteHeader(http.StatusInternalServerError)
	//			return
	//		}
	//	}
	//
	//	w.Header().Set("Content-Type", "application/json")
	//	json, err := json.Marshal(profile)
	//	if err != nil {
	//		log.Println(err, "in profileMethod")
	//		w.WriteHeader(http.StatusInternalServerError)
	//		return
	//	}
	//	fmt.Fprintln(w, string(json))
	//} else if params.Nickname != "" {
	//	profile, err := db.GetUserProfileByNickname(params.Nickname)
	//	if err != nil {
	//		switch err.(type) {
	//		case db.UserNotFoundError:
	//			w.WriteHeader(http.StatusNotFound)
	//			return
	//		default:
	//			log.Println(err)
	//			w.WriteHeader(http.StatusInternalServerError)
	//			return
	//		}
	//	}
	//
	//	w.Header().Set("Content-Type", "application/json")
	//	json, err := json.Marshal(profile)
	//	if err != nil {
	//		log.Println(err, "in profileMethod")
	//		w.WriteHeader(http.StatusInternalServerError)
	//		return
	//	}
	//	fmt.Fprintln(w, string(json))
	//} else {
	//
	//	//get auth
	//	if !r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
	//		w.WriteHeader(http.StatusUnauthorized)
	//		return
	//	}
	//	profile, err := db.GetUserProfileByID(r.Context().Value(middleware.KeyUserID).(uint))
	//	if err != nil {
	//		switch err.(type) {
	//		case db.UserNotFoundError:
	//			w.WriteHeader(http.StatusNotFound)
	//			return
	//		default:
	//			log.Println(err)
	//			w.WriteHeader(http.StatusInternalServerError)
	//			return
	//		}
	//	}
	//
	//	w.Header().Set("Content-Type", "application/json")
	//	json, err := json.Marshal(profile)
	//	if err != nil {
	//		log.Println(err, "in profileMethod")
	//		w.WriteHeader(http.StatusInternalServerError)
	//		return
	//	}
	//	fmt.Fprintln(w, string(json))
	//}
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

func (h *Handler) postSignupProfile(w http.ResponseWriter, r *http.Request) {
	//parsedata
	u := &models.RegisterProfile{}
	err := readProfile(r, u)
	if err != nil {
		switch err.(type) {
		case models.ParseJSONError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	if u.Nickname == "" || u.Email == "" || u.Password == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	fieldErrors, err := validateFields(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(fieldErrors) != 0 {
		json, err := json.Marshal(models.ProfileErrorList{Errors: fieldErrors})
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, string(json))
	} else {

		//logic
		newU, err := h.useCase.PostUser(r.Context(), u)

		//newU, err := db.CreateNewUser(u)
		if err != nil {
			if err == db.ErrUniqueConstraintViolation ||
				err == db.ErrNotNullConstraintViolation {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}
			fmt.Println("4")
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = auth.LoginUser(w, newU.UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//проблемы с тестами инициализация логгера доделать
		//fmt.Print("New film with id , title created", newU.UserID, newU.Email)
		logger.Infof("New user with id %v, email %v and nickname %v logged in", newU.UserID, newU.Email, newU.Nickname)
	}
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

func (h *Handler) putEditUserProfile(w http.ResponseWriter, r *http.Request) {

	if !r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	editUser := &models.RegisterProfile{}
	err := readProfile(r, editUser)

	if err != nil {
		switch err.(type) {
		case models.ParseJSONError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	var fieldErrors []models.ProfileError

	if editUser.Nickname != "" {
		valErrors, dbErr := validateNickname(editUser.Nickname)
		if dbErr != nil {
			log.Println(dbErr)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fieldErrors = append(fieldErrors, valErrors...)
	}
	if editUser.Email != "" {
		valErrors, dbErr := validateEmail(editUser.Email)
		if dbErr != nil {
			log.Println(dbErr)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fieldErrors = append(fieldErrors, valErrors...)
	}
	if editUser.Password != "" {
		fieldErrors = append(fieldErrors, validatePassword(editUser.Password)...)
	}

	if len(fieldErrors) != 0 {
		jsonObject, err := json.Marshal(models.ProfileErrorList{Errors: fieldErrors})
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, string(jsonObject))
	} else {
		id := r.Context().Value(middleware.KeyUserID).(uint)

		//logic

		err := h.useCase.PutUser(r.Context(), id, editUser)
		//err := db.UpdateUserByID(id, editUser)

		if err != nil {
			switch err.(type) {
			case db.UserNotFoundError:
				w.WriteHeader(http.StatusNotFound)
			default:
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		log.Println("User with id", id, "changed to", editUser.Nickname, editUser.Email)
	}
}
