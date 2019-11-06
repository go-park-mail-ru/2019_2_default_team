package handlers

import (
	"2019_2_default_team/db"
	"2019_2_default_team/logger"
	"2019_2_default_team/middleware"
	"2019_2_default_team/models"
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/asaskevich/govalidator"
)

type HandlerDB struct {
	DB *sql.DB
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func NewMyHandler() *MyHandler {
	return &MyHandler{
		sessions: make(map[string]uint64, 0),
		usersAuth: map[string]*User{
			"testUser": {1, "testuser", "test"},
		},
		users: make([]User, 0),
		mu:    &sync.Mutex{},
	}
}

//вспомогательные функции

func (api *MyHandler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getProfile(w, r)
	case http.MethodPost:
		postSignupProfile(w, r)
	case http.MethodPut:
		putEditUserProfile(w, r)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
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

//основные методы к профилю юзера

func (api *MyHandler) GetMyProfile(w http.ResponseWriter, r *http.Request) {

	if !r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Println("hello from myprofile")
	session, _ := r.Cookie("session_id")
	encoder := json.NewEncoder(w)
	w.Write([]byte("your profile"))
	api.mu.Lock()
	err := encoder.Encode(api.users[api.sessions[session.Value]])
	api.mu.Unlock()
	if err != nil {
		log.Printf("error marshal json %s", err)
	}
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

func getProfile(w http.ResponseWriter, r *http.Request) {
	//data parse
	params := &models.RequestProfile{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(params)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//end data parse

	if params.ID != 0 {
		profile, err := db.GetUserProfileByID(params.ID)
		if err != nil {
			switch err.(type) {
			case db.UserNotFoundError:
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
	} else if params.Nickname != "" {
		profile, err := db.GetUserProfileByNickname(params.Nickname)
		if err != nil {
			switch err.(type) {
			case db.UserNotFoundError:
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
	} else {

		//get auth
		if !r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		profile, err := db.GetUserProfileByID(r.Context().Value(middleware.KeyUserID).(uint))
		if err != nil {
			switch err.(type) {
			case db.UserNotFoundError:
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

func postSignupProfile(w http.ResponseWriter, r *http.Request) {
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
		newU, err := db.CreateNewUser(u)
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

		err = loginUser(w, newU.UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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

func putEditUserProfile(w http.ResponseWriter, r *http.Request) {

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
		err := db.UpdateUserByID(id, editUser)

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

func getPhoto(id int) (os.File, error) {
	fileName := strconv.Itoa(id)
	file, err := os.Open("./imagesupload/" + fileName + ".jpg")
	if err != nil {
		log.Printf("An error occurred: %v", err)
		return *file, err
	}
	return *file, nil
}

func Download(file multipart.File, id string) (returnErr error) {
	defer func() {
		err := file.Close()

		if err != nil && returnErr == nil {
			log.Printf("error: %v", err)
			returnErr = err
		}
	}()

	tempFile, err := ioutil.TempFile("imagesupload", "upload-*.jpg")
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	defer func() {
		err := tempFile.Close()

		if err != nil && returnErr == nil {
			log.Printf("error: %v", err)
			returnErr = err
		}
	}()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	err = os.Rename(tempFile.Name(), "imagesupload/"+id+".jpg")

	if err != nil {
		log.Printf("An error occurred: %v", err)
		return err
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	return nil
}

func (api *MyHandler) UploadPage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 * 1024 * 1025)
	file, handler, err := r.FormFile("my_file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "handler.Filename %v\n", handler.Filename)
	fmt.Fprintf(w, "handler.Header %#v\n", handler.Header)
	session, err := r.Cookie("session_id")
	id := api.sessions[session.Value]
	strid := strconv.Itoa(int(id))
	error := Download(file, strid)
	if error != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (api *MyHandler) GetPhoto(w http.ResponseWriter, r *http.Request) {

	authorized := false
	session, err := r.Cookie("session_id")
	if err == nil && session != nil {
		_, authorized = api.sessions[session.Value]
	}

	if authorized {
		id := api.sessions[session.Value]
		file, err := getPhoto(int(id))
		if err != nil {
			log.Printf("An error occurred: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		reader := bufio.NewReader(&file)
		bytes := make([]byte, 10<<20)
		_, err = reader.Read(bytes)

		w.Header().Set("content-type", "multipart/form-data;boundary=1")

		_, err = w.Write(bytes)
		if err != nil {
			log.Printf("An error occurred: %v", err)
			w.WriteHeader(500)
			return
		}

		log.Println("Successfully Uploaded File")

	} else {
		w.Write([]byte("not autrorized"))
	}
}
