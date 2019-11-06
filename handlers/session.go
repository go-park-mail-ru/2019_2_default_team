package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"time"
	"kino_backend/models"
	"kino_backend/logger"
	"kino_backend/sessions"
	"kino_backend/middleware"
	"kino_backend/db"
)


//вспомогательные методы

func (api *MyHandler)  SessionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getSession(w, r)
	case http.MethodPost:
		postLoginHandler(w, r)
	case http.MethodDelete:
		deleteSession(w, r)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
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

func loginUser(w http.ResponseWriter, userID uint) error {
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


// @Summary Получить сессию
// @Description Получить сессию пользователя, если есть сессия, то она в куке session_id
// @ID get-session
// @Produce json
// @Success 200 {object} models.Session "Пользователь залогинен, успешно"
// @Failure 401 "Не залогинен"
// @Failure 500 "Ошибка в бд"
// @Router /session [GET]
func getSession(w http.ResponseWriter, r *http.Request) {
	if r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		sID, err := json.Marshal(models.Session{SessionID: r.Context().Value(middleware.KeySessionID).(string)})
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		csrfToken := generateCSRFToken()
		sessions.AddCSRFToken(sID, csrfToken)
		w.Header().Set("X-csrf-token", csrfToken)

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
func postLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		// user has already logged in
		return
	}

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

	dbResponse, err := db.GetUserPassword(u.Email)

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
		err := loginUser(w, dbResponse.UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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
func deleteSession(w http.ResponseWriter, r *http.Request) {
	if !r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		// user has already logged out
		return
	}
	err := sessions.Delete(r.Context().Value(middleware.KeySessionID).(string))
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

// @Summary Сгенерировать CSRF-токен
// @Descriptiond Функция генерации CSRF-токена (возвращает строку)
func generateCSRFToken() string {
	timeNanoString := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(timeNanoString)
	hashInBytes := md5.Sum([]byte(timeNanoString))
	return hex.EncodeToString(hashInBytes[:])
}
/*

func loginUser(w http.ResponseWriter, userID uint) error {
	sessionID := ""
	for {
		var err error
		u, _ := uuid.NewV4()
		sessionID = u.String()
		fmt.Println("ses id ", sessionID)

		//sessionID = uuid.NewV4().String()
		exists, err := db.CheckExistenceOfSession(sessionID)
		if err != nil {
			log.Println(err)
			return err
		}
		if !exists {
			break
		}
	}

	err := db.CreateNewSession(sessionID, userID)
	if err != nil {
		log.Println(err)
		return err
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	fmt.Println(cookie.Value)


	return nil
}

// @Title Получить сессию
// @Summary Получить сессию пользователя, если есть сессия, то она в куке session_id
// @ID get-session
// @Produce json
// @Success 200 {object} models.Session "Пользователь залогинен, успешно"
// @Failure 401 "Не залогинен"
// @Failure 500 "Ошибка в бд"
// @Router /session [GET]
func getSession(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if err == nil && session != nil {
			fmt.Println("see  ", session.Value)
		} else{
			fmt.Println("no sess")
		}

	if r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		sID, err := json.Marshal(models.Session{SessionID: r.Context().Value(middleware.KeySessionID).(string)})
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(sID))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// @Title Залогинить
// @Summary Залогинить пользователя (создать сессию)
// @ID post-session
// @Accept json
// @Produce json
// @Param UserPassword body models.UserPassword true "Почта и пароль"
// @Success 200 {object} models.Session "Успешный вход / пользователь уже залогинен"
// @Failure 400 "Неверный формат JSON"
// @Failure 422 "Неверная пара пользователь/пароль"
// @Failure 500 "Внутренняя ошибка"
// @Router /session [POST]

func postLoginHandler(w http.ResponseWriter, r http.Request){
	if r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		// logged in
		return
	}

	u := &models.UserPassword{}
	err := readLoginInfo(&r, u)
	if err != nil {
		switch err.(type) {
		case models.ParseJSONError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			log.Println(err, "in sessionHandler in getUserFromRequestBody")
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	if u.Email == "" || u.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbResponse, err := db.GetUserPassword(u.Email)

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
		err := loginUser(w, dbResponse.UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println("User logged in:", dbResponse.UserID, dbResponse.Email)
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
}

// @Title Разлогинить
// @Summary Разлогинить пользователя (удалить сессию)
// @ID delete-session
// @Success 200 "Успешный выход / пользователь уже разлогинен"
// @Router /session [DELETE]
func deleteSession(w http.ResponseWriter, r *http.Request) {
	if !r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		// logged out
		fmt.Println("hekllodelete1")
		return
	}
	err := db.DeleteSession(r.Context().Value(middleware.KeySessionID).(string))
	if err != nil { // but we continue
		log.Println(err)
	}
	fmt.Println("hekllodelete2")
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Expires:  time.Now().AddDate(0, 0, -1),
		Secure:   true,
		HttpOnly: true,
	})
}






func (api *MyHandler) Login(w http.ResponseWriter, r *http.Request) {

	ok := false
	var user models.Profile
	for _, value := range db.Users{
		if value.Nickname == r.FormValue("login"){
			if value.Password == r.FormValue("password"){
				ok = true
				user = value
			}
		}
	}
	if !ok {
		http.Error(w, `invalid login or password`, 404)
		return
	}


	SID := RandStringRunes(32)

	_, err := db.CheckExistenceOfSession(SID)
	if err != nil {
		log.Println(err)
		return
	}


	err = db.CreateNewSession(SID, user.UserID)
	if err != nil {
	log.Println(err)
	return
	}

	//db.Sessions[SID] = user.ID
	fmt.Println(user)

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   SID,
		Expires: time.Now().Add(10 * time.Hour),
		HttpOnly:true,
	}
	http.SetCookie(w, cookie)
	w.Write([]byte(SID))

}

func (api *MyHandler) Logout(w http.ResponseWriter, r *http.Request) {

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Error(w, `no sess`, 401)
		return
	}

	if _, ok := api.sessions[session.Value]; !ok {
		http.Error(w, `no sess`, 401)
		return
	}

	delete(api.sessions, session.Value)

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}


*/