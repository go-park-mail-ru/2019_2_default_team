package chat_support_delivery

import (
	"github.com/gorilla/websocket"
	"kino_backend/logger"
	"kino_backend/useCase"
	"kino_backend/utilits/middleware"
	"net/http"
)

type Handler struct {
	supuseCase useCase.SupportChatUseCase
}

func NewHandler(useCase useCase.SupportChatUseCase) *Handler {
	return &Handler{
		supuseCase: useCase,
	}
}

func (h *Handler) ConnectChat(w http.ResponseWriter, r *http.Request) {
	u := &useCase.User{}
	if r.Context().Value(middleware.KeyIsAuthenticated).(bool) {
		u.SessionID = r.Context().Value(middleware.KeySessionID).(string)
		u.UID = r.Context().Value(middleware.KeyUserID).(uint)
	} else {
		u.Anon = true
	}
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("Cannot upgrade connection: ", err)
		return
	}
	u.Conn = conn

	h.supuseCase.JoinUser(u)
}
