package sessions

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"kino_backend/logger"
)

var Sm *SessionManager

var (
	ErrKeyNotFound = errors.New("key not found")
)

type SessionManager struct {
	redisConn redis.Conn
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
	var err error
	Sm = &SessionManager{}

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

func Create(sID string, uID uint) (bool, error) {
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

func Get(sID string) (uint, error) {
	res, err := redis.Uint64(Sm.redisConn.Do("GET", sID))
	if err != nil {
		if err == redis.ErrNil {
			return 0, ErrKeyNotFound
		}
		return 0, err
	}

	return uint(res), nil
}

func Delete(sID string) error {
	_, err := redis.Int(Sm.redisConn.Do("DEL", sID))
	if err != nil {
		return err
	}

	return nil
}