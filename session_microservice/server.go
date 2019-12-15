package session_microservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"kino_backend/logger"
	"kino_backend/session_microservice_client"
	"time"
)

var Sm *SessionManager
var Rd *redis.Conn

var (
	ErrKeyNotFound = errors.New("key not found")
)

type SessionManager struct {
	//	RedisConn redis.Conn
	redisConnPool *redis.Pool
}

type SessionRepository struct {
	database redis.Conn
}

func NewSessionManager(address, database string) *SessionManager {
	sm := &SessionManager{}
	err := sm.Open(address, database)
	if err != nil {
		logger.Panic(err)
	}

	logger.Infof("Successfully connected to %v, database %v", address, database)

	return sm
}

func (sm *SessionManager) Open(address, database string) error {
	sm.redisConnPool = &redis.Pool{
		MaxIdle:     500,
		IdleTimeout: 240 * time.Second,
		MaxActive:   1000,
		Wait:        true,
		Dial:        func() (redis.Conn, error) { return redis.DialURL("redis://redis:docker@redis_db:6379/0?") },
		//redis.DialURL("redis://redis:docker@localhost:6379/0?")
		//Dial:        func() (redis.Conn, error) { return redis.DialURL("redis://" + address + "/" + database) },
	}
	conn := sm.redisConnPool.Get()
	defer conn.Close()
	_, err := conn.Do("PING")
	return err
}

//func NewSessionManager(db redis.Conn) SessionManager {
//	return SessionManager{
//		RedisConn: db,
//	}
//}

func NewSessionsRepository(db redis.Conn) SessionRepository {
	return SessionRepository{
		database: db,
	}
}

func (sm *SessionManager) Close() {
	sm.redisConnPool.Close()
}

const (
	host     = "localhost"
	port     = 6379
	user     = "redis"
	password = "docker"
	dbname   = "redis"
)

//func ConnectSessionDB(address, database string) *SessionManager {
//	Sm = &SessionManager{}
//	var err error
//
//	//redisInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//	//	"password=%s dbname=%s",
//	//	host, port, user, password, dbname)
//
//	//Sm.PostgresConn, err = redis.DialURL("redis://" + address + "/" + database)
//	//fmt.Println(redisInfo)
//	Sm.RedisConn, err = redis.DialURL("redis://redis:docker@localhost:6379/0?")
//	if err != nil {
//		logger.Panic(err)
//	}
//
//	logger.Infof("Successfully connected to %v, database %v", address, database)
//
//	return Sm
//}

func (s *SessionManager) Create(ctx context.Context, in *session_microservice_client.Session) (*session_microservice_client.Success, error) {
	sID := ""
	conn := s.redisConnPool.Get()
	defer conn.Close()
	for {
		sID = in.SID
		res, err := conn.Do("SET", sID, in.UID, "NX", "EX", 30*24*60*60)
		if err != nil {
			return &session_microservice_client.Success{Success: false}, status.Error(codes.Internal, err.Error())
		}
		if res != "OK" {
			logger.Infow("collision, session not created",
				"sID", sID,
				"uID", in.UID,
			)
			continue
		}
		break
	}

	logger.Infow("session created",
		"sID", sID,
		"uID", in.UID,
	)

	return &session_microservice_client.Success{Success: true}, nil

	//res, err := s.RedisConn.Do("SET", in.SID, in.UID, "NX", "EX", 30*24*60*60)
	//if err != nil {
	//	return &session_microservice_client.Success{Success: false}, err
	//}
	//if res != "OK" {
	//	logger.Infow("collision, session not created",
	//		"sID", in.SID,
	//		"uID", in.UID,
	//	)
	//	return &session_microservice_client.Success{Success: false}, nil
	//}
	//
	//logger.Infow("session created",
	//	"sID", in.SID,
	//	"uID", in.UID,
	//)
	//
	//return &session_microservice_client.Success{Success: true}, nil
}

func (s *SessionManager) Get(ctx context.Context, in *session_microservice_client.SessionID) (*session_microservice_client.Session, error) {

	conn := s.redisConnPool.Get()
	defer conn.Close()
	res, err := redis.Uint64(conn.Do("GET", in.SID))
	if err != nil {
		if err == redis.ErrNil {
			return &session_microservice_client.Session{}, status.Error(codes.NotFound, ErrKeyNotFound.Error())
		}
		return &session_microservice_client.Session{}, status.Error(codes.Internal, err.Error())
	}

	return &session_microservice_client.Session{UID: res}, nil

	//res, err := redis.Uint64(s.RedisConn.Do("GET", in.SID))
	//if err != nil {
	//	if err == redis.ErrNil {
	//		return &session_microservice_client.Session{}, ErrKeyNotFound
	//	}
	//	return &session_microservice_client.Session{}, err
	//}
	//
	//return &session_microservice_client.Session{UID: res}, nil
}

func (s *SessionManager) Delete(ctx context.Context, in *session_microservice_client.SessionID) (*session_microservice_client.Nothing, error) {
	conn := s.redisConnPool.Get()
	defer conn.Close()
	fmt.Println(in.SID)
	_, err := redis.Int(conn.Do("DEL", in.SID))
	if err != nil {
		return &session_microservice_client.Nothing{}, status.Error(codes.Internal, err.Error())
	}

	logger.Infow("session deleted",
		"sID", in.SID,
	)

	return &session_microservice_client.Nothing{}, nil

	//_, err := redis.Int(s.RedisConn.Do("DEL", in.SID))
	//if err != nil {
	//	return &session_microservice_client.Nothing{}, err
	//}
	//
	//return &session_microservice_client.Nothing{}, nil
}
