package main

import (
	"kino_backend/db"
	"kino_backend/logger"
	"kino_backend/repository"
	"kino_backend/server"
	"kino_backend/sessions"
	"log"
)

func main() {

	l := logger.InitLogger()

	//db.Db = db.InitDB("postgres@postgres:5432", "some-postgres")

	//база для postgres
	database := db.InitDB("postgres@postgres:5432", "some-postgres")

	//sessions.Sm = sessions.ConnectSessionDB("user@redis:6379", "0")

	//база для redis
	redis := sessions.ConnectSessionDB("user@redis:6379", "0")

	newServer, err := server.CreateServer(database, redis)

	if err != nil {
		log.Printf("An error occurred: %v", err)
		return
	}

	newServer.RunServer()

	defer redis.Close()
	defer repository.Sm.Close()
	//defer db.Db.Close()
	defer database.Close()
	defer l.Sync()

}
