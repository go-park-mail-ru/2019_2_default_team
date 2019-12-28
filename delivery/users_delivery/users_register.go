package users_delivery

import (
	"kino_backend/session_microservice_client"
	"kino_backend/useCase"
	"net/http"
	"sync"
)

type MyHandlerUser struct {
	mu      *sync.Mutex
	useCase useCase.UsersUseCase
	uS      *session_microservice_client.SessionManager
}

func NewMyHandlerUser(uc useCase.UsersUseCase, hS *session_microservice_client.SessionManager) *MyHandlerUser {
	return &MyHandlerUser{
		mu:      &sync.Mutex{},
		useCase: uc,
		uS:      hS,
	}
}

func (apiu *MyHandlerUser) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apiu.useCase, apiu.uS)

	switch r.Method {
	case http.MethodGet:
		h.getProfile(w, r)
	case http.MethodPost:
		h.postSignupProfile(w, r)
	case http.MethodPut:
		h.putEditUserProfile(w, r)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (apiu *MyHandlerUser) ProfileRegHandler(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apiu.useCase, apiu.uS)

	h.postSignupProfile(w, r)
}
