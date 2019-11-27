package session_microservice

import (
	"context"
	"errors"
	"github.com/gomodule/redigo/redis"
	"kino_backend/logger"
	"kino_backend/session_microservice_client"
)

var Sm *SessionManager
var Rd *redis.Conn

var (
	ErrKeyNotFound = errors.New("key not found")
)

type SessionManager struct {
	RedisConn redis.Conn
}

type SessionRepository struct {
	database redis.Conn
}

func NewSessionsRepository(db redis.Conn) SessionRepository {
	return SessionRepository{
		database: db,
	}
}

func (sm *SessionManager) Close() {
	sm.RedisConn.Close()
}

const (
	host     = "localhost"
	port     = 6379
	user     = "redis"
	password = "docker"
	dbname   = "redis"
)

func ConnectSessionDB(address, database string) *SessionManager {
	Sm = &SessionManager{}
	var err error

	//redisInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s",
	//	host, port, user, password, dbname)

	//Sm.PostgresConn, err = redis.DialURL("redis://" + address + "/" + database)
	//fmt.Println(redisInfo)
	Sm.RedisConn, err = redis.DialURL("redis://redis:docker@localhost:6379/0?")
	if err != nil {
		logger.Panic(err)
	}

	logger.Infof("Successfully connected to %v, database %v", address, database)

	return Sm
}

func (s *SessionManager) Create(ctx context.Context, in *session_microservice_client.Session) (*session_microservice_client.Success, error) {
	res, err := s.RedisConn.Do("SET", in.SID, in.UID, "NX", "EX", 30*24*60*60)
	if err != nil {
		return &session_microservice_client.Success{Success: false}, err
	}
	if res != "OK" {
		logger.Infow("collision, session not created",
			"sID", in.SID,
			"uID", in.UID,
		)
		return &session_microservice_client.Success{Success: false}, nil
	}

	logger.Infow("session created",
		"sID", in.SID,
		"uID", in.UID,
	)

	return &session_microservice_client.Success{Success: true}, nil
}

func (s *SessionManager) Get(ctx context.Context, in *session_microservice_client.SessionID) (*session_microservice_client.Session, error) {
	res, err := redis.Uint64(s.RedisConn.Do("GET", in.SID))
	if err != nil {
		if err == redis.ErrNil {
			return &session_microservice_client.Session{}, ErrKeyNotFound
		}
		return &session_microservice_client.Session{}, err
	}

	return &session_microservice_client.Session{UID: res}, nil
}

func (s *SessionManager) Delete(ctx context.Context, in *session_microservice_client.SessionID) (*session_microservice_client.Nothing, error) {
	_, err := redis.Int(s.RedisConn.Do("DEL", in.SID))
	if err != nil {
		return &session_microservice_client.Nothing{}, err
	}

	return &session_microservice_client.Nothing{}, nil
}
