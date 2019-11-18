package repository

import (
	"context"
	"errors"
	"github.com/gomodule/redigo/redis"
	"kino_backend/logger"
)

var Sm *sessionManager
var Rd *redis.Conn

var (
	ErrKeyNotFound = errors.New("key not found")
)

type sessionManager struct {
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

func (sm *sessionManager) Close() {
	sm.RedisConn.Close()
}

const (
	host     = "localhost"
	port     = 6379
	user     = "redis"
	password = "docker"
	dbname   = "redis"
)

func ConnectSessionDB(address, database string) *sessionManager {
	//Sm = &sessionManager{}
	var err error

	//redisInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s",
	//	host, port, user, password, dbname)

	//Sm.RedisConn, err = redis.DialURL("redis://" + address + "/" + database)
	//fmt.Println(redisInfo)
	Sm.RedisConn, err = redis.DialURL("redis://redis:docker@localhost:6379/0?")
	if err != nil {
		logger.Panic(err)
	}

	logger.Infof("Successfully connected to %v, database %v", address, database)

	return Sm
}

func (s SessionRepository) Create(sID string, uID uint) (bool, error) {
	res, err := s.database.Do("SET", sID, uID, "NX", "EX", 30*24*60*60)
	if err != nil {
		return false, err
	}
	if res != "OK" {
		logger.Infow("collision, session not created",
			"sID", sID,
			"uID", uID,
		)
		return false, nil
	}

	logger.Infow("session created",
		"sID", sID,
		"uID", uID,
	)

	return true, nil
}

func (s SessionRepository) Get(ctx context.Context, sID string) (uint, error) {
	res, err := redis.Uint64(s.database.Do("GET", sID))
	if err != nil {
		if err == redis.ErrNil {
			return 0, ErrKeyNotFound
		}
		return 0, err
	}

	return uint(res), nil
}

func (s SessionRepository) Delete(ctx context.Context, sID string) error {
	_, err := redis.Int(s.database.Do("DEL", sID))
	if err != nil {
		return err
	}

	return nil
}
