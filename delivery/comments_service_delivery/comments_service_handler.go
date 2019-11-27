package comments_service_delivery

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"io/ioutil"
	"kino_backend/comments_microservice_client"
	"kino_backend/logger"
	"kino_backend/models"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	manager *comments_microservice_client.CommentsManager
}

func NewHandler(man *comments_microservice_client.CommentsManager) *Handler {
	return &Handler{
		manager: man,
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

func readComment(r *http.Request, p *comments_microservice_client.Comment) error {
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

func SanitizeMe(comment comments_microservice_client.Comment) comments_microservice_client.Comment {
	sanitizer := bluemonday.UGCPolicy()
	comment.Username = sanitizer.Sanitize(comment.Username)
	comment.FilmTitle = sanitizer.Sanitize(comment.FilmTitle)
	comment.Text = sanitizer.Sanitize(comment.Text)

	return comment
}

func SanitizeMeComments(comment comments_microservice_client.CommentsResponse) comments_microservice_client.CommentsResponse {
	sanitizer := bluemonday.UGCPolicy()
	for _, value := range comment.Comments {
		value.Username = sanitizer.Sanitize(value.Username)
		value.FilmTitle = sanitizer.Sanitize(value.FilmTitle)
		value.Text = sanitizer.Sanitize(value.Text)
	}
	return comment
}

// @Title Добавить коммент
// @Summary
// @ID post-profile
// @Accept json
// @Produce json
// @Param Comment
// @Success 200 "Added"

// @Failure 500 "Ошибка"
// @Router /comment [POST]

func (h *Handler) postComment(w http.ResponseWriter, r *http.Request) {

	//начала парсинга данных
	u := &comments_microservice_client.Comment{}
	err := readComment(r, u)
	if err != nil {
		switch err.(type) {
		case models.ParseJSONError:
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	fmt.Println(u)
	if u.FilmTitle == "" || u.Username == "" || u.Text == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//логика
	_, err = h.manager.CreateComment(*u)

	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//fmt.Fprintln(w, http.StatusOK)
	w.WriteHeader(http.StatusOK)

	logger.Info("New comment was added ", u.FilmTitle, u.Username)
}

// @Title Get comment wit comment id
// @Summary
// @ID get-comment
// @Accept json
// @Produce json
// @Param
// @Success 200 "added"
// @Failure 500 "Ошибка в бд"
// @Router /comment [GET]

func (h *Handler) getComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err, "error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	params := &comments_microservice_client.CommentID{}
	params.CID = uint64(ID)

	//конец парсинга данных можно передавать ctx,

	comment, err := h.manager.GetComment(*params)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	comment = SanitizeMe(comment)
	json, err := json.Marshal(comment)
	if err != nil {
		log.Println(err, "in profileMethod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))

}

// @Title Get comment by film
// @Summary
// @ID get-comment
// @Accept json
// @Produce json
// @Param
// @Success 200 "added"
// @Failure 500 "Ошибка в бд"
// @Router /comment [GET]

func (h *Handler) getCommentByFilmTitle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	filmTitle := vars["film"]

	params := &comments_microservice_client.CommentID{}
	params.Film = filmTitle

	//конец парсинга данных можно передавать ctx,

	comments, err := h.manager.GetCommentsByFilmID(*params)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	comments = SanitizeMeComments(comments)
	json, err := json.Marshal(comments)
	if err != nil {
		log.Println(err, "in profileMethod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))

}

// @Title Get comment by film
// @Summary
// @ID get-comment
// @Accept json
// @Produce json
// @Param
// @Success 200 "added"
// @Failure 500 "Ошибка в бд"
// @Router /comment [GET]

func (h *Handler) getCommentByUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]

	params := &comments_microservice_client.CommentID{}
	params.User = username

	//конец парсинга данных можно передавать ctx,

	comments, err := h.manager.GetCommentsByUserID(*params)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	comments = SanitizeMeComments(comments)
	json, err := json.Marshal(comments)
	if err != nil {
		log.Println(err, "in profileMethod")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(json))

}
