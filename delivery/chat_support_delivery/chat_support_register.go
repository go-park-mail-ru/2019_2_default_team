package chat_support_delivery

import (
	"kino_backend/useCase"
	"net/http"
	"sync"
)

type MyHandlerCS struct {
	mu         *sync.Mutex
	supuseCase useCase.SupportChatUseCase
}

func NewMyHandlerCS(uc useCase.SupportChatUseCase) *MyHandlerCS {
	return &MyHandlerCS{
		mu:         &sync.Mutex{},
		supuseCase: uc,
	}
}

func (apisc *MyHandlerCS) SupportChat(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apisc.supuseCase)

	h.ConnectChat(w, r)
}
