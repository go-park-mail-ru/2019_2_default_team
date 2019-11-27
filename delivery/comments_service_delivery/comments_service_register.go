package comments_service_delivery

import (
	"kino_backend/comments_microservice_client"
	"net/http"
	"sync"
)

type MyHandlerSessionsMicroservice struct {
	mu      *sync.Mutex
	manager *comments_microservice_client.CommentsManager
}

func NewMyHandlerFilms(m *comments_microservice_client.CommentsManager) *MyHandlerSessionsMicroservice {
	return &MyHandlerSessionsMicroservice{
		mu:      &sync.Mutex{},
		manager: m,
	}
}

func (apic *MyHandlerSessionsMicroservice) CommentsHandler(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apic.manager)

	switch r.Method {
	case http.MethodPost:
		h.postComment(w, r)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (apic *MyHandlerSessionsMicroservice) CommentsByFilmHandler(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apic.manager)

	h.getCommentByFilmTitle(w, r)
}

func (apic *MyHandlerSessionsMicroservice) CommentsByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apic.manager)

	h.getCommentByUser(w, r)
}

func (apic *MyHandlerSessionsMicroservice) CommentsByIDHandler(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apic.manager)

	h.getComment(w, r)
}
