package delivery

import (
	"2019_2_default_team/tickets"
	"net/http"
	"sync"
)

type MyHandlerTicket struct {
	mu      *sync.Mutex
	useCase tickets.UseCase
}

func NewMyHandlerTicket(uc tickets.UseCase) *MyHandlerTicket {
	return &MyHandlerTicket{
		mu:      &sync.Mutex{},
		useCase: uc,
	}
}

func (apit *MyHandlerTicket) ProfileTicketHandler(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apit.useCase)

	switch r.Method {
	case http.MethodGet:
		h.getProfileTicket(w, r)
	case http.MethodPost:
		h.postBuyTicket(w, r)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
