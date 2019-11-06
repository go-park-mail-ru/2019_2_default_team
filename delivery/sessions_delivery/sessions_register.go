package sessions_delivery

import (
	"kino_backend/useCase"
	"net/http"
	"sync"
)

type MyHandlerSessions struct{
	mu      *sync.Mutex
	useCase useCase.SessionsUseCase
}

func NewMyHandlerFilms(uc useCase.SessionsUseCase) *MyHandlerSessions {
	return &MyHandlerSessions{
		mu: &sync.Mutex{},
		useCase: uc,
	}
}

func (apis *MyHandlerSessions) ProfileSessionsHandler(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apis.useCase)

	switch r.Method {
	case http.MethodGet:
		h.getSession(w, r)
	case http.MethodPost:
		h.postLoginHandler(w, r)
	case http.MethodDelete:
		h.deleteSession(w, r)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}


