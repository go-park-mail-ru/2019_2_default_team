package sessions_service_delivery

import (
	"fmt"
	"kino_backend/session_microservice_client"
	"kino_backend/useCase"
	"net/http"
	"sync"
)

type MyHandlerSessionsMicroservice struct {
	mu          *sync.Mutex
	manager     *session_microservice_client.SessionManager
	useCaseUser useCase.UsersUseCase
}

func NewMyHandlerSessions(m *session_microservice_client.SessionManager, user useCase.UsersUseCase) *MyHandlerSessionsMicroservice {
	return &MyHandlerSessionsMicroservice{
		mu:          &sync.Mutex{},
		manager:     m,
		useCaseUser: user,
	}
}

func (apis *MyHandlerSessionsMicroservice) ProfileSessionsMicroserviceHandler(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apis.manager, apis.useCaseUser)

	switch r.Method {
	case http.MethodGet:
		h.getSession(w, r)
	case http.MethodPost:
		h.postLoginHandler(w, r)
	case http.MethodDelete:
		fmt.Println("DELETE")
		h.deleteSession(w, r)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
