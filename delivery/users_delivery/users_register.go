package users_delivery

import (
	"net/http"
	"sync"
	"kino_backend/useCase"
)

type MyHandlerUser struct{
	mu  *sync.Mutex
	useCase useCase.UsersUseCase
}

func NewMyHandlerUser(uc useCase.UsersUseCase) *MyHandlerUser {
	return &MyHandlerUser{
		mu: &sync.Mutex{},
		useCase: uc,
	}
}

func (apiu *MyHandlerUser) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apiu.useCase)

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

