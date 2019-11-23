package useCase

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"kino_backend/logger"
	"kino_backend/models"
	"kino_backend/repository"
	"sync"
)

type SupportChatUseCase interface {
	Run()
	acceptJoiningUsers()
	acceptLeavingUsers()
	acceptSendingMessages()
	sendWSMessageToSession(m *ProcessWSMessage)
	getAllMessages(m *ProcessWSMessage)
	sendMessage(m *ProcessWSMessage)
	JoinUser(u *User)
}

type supportChatsUseCase struct {
	chatRepo repository.SupportChatRepository
	c        *Chat
}

func NewSupportChatsUseCase(chatRepo repository.SupportChatRepository, chat *Chat) *supportChatsUseCase {
	return &supportChatsUseCase{
		chatRepo: chatRepo,
		c:        chat,
	}
}

var chat *Chat

type Chat struct {
	Users *sync.Map
	SC    supportChatsUseCase

	Join  chan *User
	Leave chan *User
	Send  chan *ProcessWSMessage
}

//easyjson:json
type WSMessageToSend struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

func (sc supportChatsUseCase) Run() {
	go sc.acceptJoiningUsers()
	go sc.acceptLeavingUsers()
	go sc.acceptSendingMessages()
}

func (sc supportChatsUseCase) acceptJoiningUsers() {
	for {
		u := <-sc.c.Join
		if u.SessionID == "" {
			u.SessionID = uuid.NewV4().String()
		}
		sc.c.Users.Store(u.SessionID, &u.Data)
		go u.Listen()
		if !u.Anon {
			logger.Infof("User with id %v joined chat with session %v", u.UID, u.SessionID)
		} else {
			logger.Infof("anonymous user joined chat with temp session %v", u.SessionID)
		}
	}
}

func (sc supportChatsUseCase) acceptLeavingUsers() {
	for {
		u := <-sc.c.Leave
		sc.c.Users.Delete(u.SessionID)
		if !u.Anon {
			logger.Infof("User with id %v left chat with session %v", u.UID, u.SessionID)
		} else {
			logger.Infof("anonymous user left chat with temp session %v", u.SessionID)
		}
	}
}

func (sc supportChatsUseCase) acceptSendingMessages() {
	for {
		m := <-sc.c.Send
		wsm := m.WSM.(*ReceivedWSMessage)

		if !m.From.Anon {
			logger.Infof("Got message from %v: action = %v, payload = %v", m.From.UID, wsm.Action, string(wsm.Payload))
		} else {
			logger.Infof("Got message from %v: action = %v, payload = %v", m.From.SessionID, wsm.Action, string(wsm.Payload))
		}

		switch wsm.Action {
		case "get":
			sc.getAllMessages(m)
		case "send":
			sc.sendMessage(m)
		default:
			sc.sendWSMessageToSession(&ProcessWSMessage{
				From: m.From,
				WSM: &WSMessageToSend{
					Action:  "error",
					Payload: "unknown action type",
				},
			})
		}
	}
}

func (sc supportChatsUseCase) sendWSMessageToSession(m *ProcessWSMessage) {
	u, ok := sc.c.Users.Load(m.From.SessionID)
	if !ok {
		logger.Info("user cannot be found")
		return
	}
	d := u.(*Data)
	wsm := m.WSM.(*WSMessageToSend)
	j, err := wsm.MarshalJSON()
	if err != nil {
		logger.Error(err)
	}
	logger.Infof("sending the message: %v", string(j))
	err = d.Conn.WriteMessage(websocket.TextMessage, j)
	if err != nil {
		logger.Infof("Error while sending to user %v: %v", *d, err)
	}
}

func (sc supportChatsUseCase) getAllMessages(m *ProcessWSMessage) {
	res, err := sc.chatRepo.GetAllGlobalMessages()
	logger.Infof("gotcha all messages, request from %v", m.From.SessionID)
	if err != nil {
		logger.Error(err)
		return
	}
	sc.sendWSMessageToSession(&ProcessWSMessage{
		From: m.From,
		WSM: &WSMessageToSend{
			Action: "get",
			Payload: models.Messages{
				Msgs: res,
			},
		},
	})
}

func (sc supportChatsUseCase) sendMessage(m *ProcessWSMessage) {
	wsm := m.WSM.(*ReceivedWSMessage)
	mess := &models.Message{}
	err := mess.UnmarshalJSON(wsm.Payload)
	if err != nil {
		logger.Infof("Message cannot be parsed: %v, message: %v", err, wsm.Payload)
		sc.sendWSMessageToSession(&ProcessWSMessage{
			From: m.From,
			WSM: &WSMessageToSend{
				Action:  "error",
				Payload: "bad payload",
			},
		})
		return
	}

	if !m.From.Anon {
		mess.Author = new(uint)
		*mess.Author = m.From.UID
	}
	mess, err = sc.chatRepo.CreateMessage(mess)
	if err != nil {
		logger.Infof("Message cannot be saved: %v", err)
		return
	}
	logger.Infof("Message saved to database: %v", *mess)

	send := &WSMessageToSend{
		Action:  "send",
		Payload: mess,
	}
	j, err := send.MarshalJSON()
	if err != nil {
		logger.Error("Marshalling ended with error: %v", err)
		return
	}
	logger.Debugf("sending the message: %v", string(j))
	if mess.To == nil {
		if mess.Author != nil {
			logger.Infof("Got global message from %v: %v", *mess.Author, mess.Message)
		} else {
			logger.Info("Got global message from anonym: ", mess.Message)
		}
		sc.c.Users.Range(func(k, v interface{}) bool {
			d := v.(*Data)
			err = d.Conn.WriteMessage(websocket.TextMessage, j)
			if err != nil {
				logger.Info(err)
			}
			if mess.Author != nil {
				logger.Infof("Message sent from %v: %v", *mess.Author, mess.Message)
			} else {
				logger.Infof("Message sent from anonym: %v", mess.Message)
			}
			return true
		})
	} else {
		if mess.Author != nil {
			sent := false
			logger.Info("Got private message from %v to %v: %v", *mess.Author, *mess.To, mess.Message)
			sc.c.Users.Range(func(k, v interface{}) bool {
				d := v.(*Data)
				if d.UID == *mess.To {
					err = d.Conn.WriteMessage(websocket.TextMessage, j)
					if err != nil {
						logger.Info(err)
						return false
					}
					sent = true
					logger.Info("Private message from %v to %v: %v", *mess.Author, *mess.To, mess.Message)
					return false
				}
				return true
			})
			if !sent {
				logger.Info("user %v is offline", *mess.To)
				sc.sendWSMessageToSession(&ProcessWSMessage{
					From: m.From,
					WSM: &WSMessageToSend{
						Action:  "send",
						Payload: "user if offline",
					},
				})
			}
		} else {
			logger.Info("anonymous users can't send private messages")
		}
	}
}

func InitChat() *Chat {
	chat = &Chat{
		Users: &sync.Map{},
		Join:  make(chan *User),
		Leave: make(chan *User),
		Send:  make(chan *ProcessWSMessage),
	}
	return chat
}

func (sc supportChatsUseCase) JoinUser(u *User) {
	sc.c.Join <- u
}

type User struct {
	SessionID string
	Data
}

type Data struct {
	UID  uint
	Conn *websocket.Conn
	Anon bool
}

type ProcessWSMessage struct {
	From *User
	WSM  interface{}
}

//easyjson:json
type ReceivedWSMessage struct {
	Action  string          `json:"action"`
	Payload json.RawMessage `json:"payload"`
}

func (u *User) Listen() {
	for {
		m := &ReceivedWSMessage{}
		_, raw, err := u.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err) {
				if !u.Anon {
					logger.Infof("User %v with session %v was disconnected", u.UID, u.SessionID)
				} else {
					logger.Infof("anonymous user with temp session %v was disconnected", u.SessionID)
				}
			} else {
				logger.Error(err)
			}
			chat.Leave <- u
			return
		}
		err = m.UnmarshalJSON(raw)
		if err != nil {
			logger.Error(err)
			continue
		}

		logger.Infof("Read WSMessage: %v", *m)

		chat.Send <- &ProcessWSMessage{
			From: u,
			WSM:  m,
		}
	}
}
