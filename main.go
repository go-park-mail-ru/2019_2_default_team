package main

import (
	"log"
	"kino_backend/server"
	"kino_backend/db"
	"kino_backend/sessions"
	"kino_backend/logger"
)

func main() {

	newServer, err := server.CreateServer()

	if err != nil {
		log.Printf("An error occurred: %v", err)
		return
	}

	l := logger.InitLogger()

	db.Db = db.InitDB("postgres@postgres:5432", "some-postgres")
	sessions.Sm = sessions.ConnectSessionDB("user@redis:6379", "0")
	newServer.RunServer()

	defer sessions.Sm.Close()
	defer db.Db.Close()
	defer l.Sync()

}