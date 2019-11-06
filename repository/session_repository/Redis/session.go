package Redis

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"kino_backend/logger"
	"errors"
)

var Sm *SessionManager
var Rd *redis.Conn

var (
	ErrKeyNotFound = errors.New("key not found")
)


type SessionManager struct {
	redisConn redis.Conn
}

type SessionsRepository struct{
	database *redis.Conn
}

func NewSessionsRepository(db *redis.Conn) *SessionsRepository{
	return &SessionsRepository{
		database:db,
	}
}

func (sm *SessionManager) Close() {
	sm.redisConn.Close()
}


const (
	host     = "localhost"
	port     = 6379
	user     = "redis"
	password = "docker"
	dbname   = "redis"
)

func ConnectSessionDB(address, database string) *SessionManager {
	//Sm = &SessionManager{}
	var err error

	//redisInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s",
	//	host, port, user, password, dbname)

	//Sm.redisConn, err = redis.DialURL("redis://" + address + "/" + database)
	//fmt.Println(redisInfo)
	Sm.redisConn, err = redis.DialURL("redis://redis:docker@localhost:6379/0?")
	if err != nil {
		logger.Panic(err)
	}

	logger.Infof("Successfully connected to %v, database %v", address, database)

	return Sm
}

func (s SessionsRepository) Create(ctx context.Context, sID string, uID uint) (bool, error) {
	res, err := Sm.redisConn.Do("SET", sID, uID, "NX", "EX", 30*24*60*60)
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

func (s SessionsRepository) Get(ctx context.Context, sID string) (uint, error) {
	res, err := redis.Uint64(Sm.redisConn.Do("GET", sID))
	if err != nil {
		if err == redis.ErrNil {
			return 0, ErrKeyNotFound
		}
		return 0, err
	}

	return uint(res), nil
}

func (s SessionsRepository) Delete(ctx context.Context, sID string) error {
	_, err := redis.Int(Sm.redisConn.Do("DEL", sID))
	if err != nil {
		return err
	}

	return nil
}