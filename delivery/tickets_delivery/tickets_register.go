package tickets_delivery

import (
	"net/http"
	"sync"
	"kino_backend/useCase"
)

type MyHandlerTicket struct{
	mu  *sync.Mutex
	useCase useCase.TicketsUseCase
}

func NewMyHandlerTicket(uc useCase.TicketsUseCase) *MyHandlerTicket {
	return &MyHandlerTicket{
		mu: &sync.Mutex{},
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
