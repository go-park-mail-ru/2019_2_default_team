package main

import (
	"google.golang.org/grpc"
	"kino_backend/comments_microservice"
	"kino_backend/comments_microservice_client"
	"kino_backend/logger"
	"net"
)

func main() {
	l := logger.InitLogger()

	//база для redis
	database := comments_microservice.InitDB("postgres@postgres:5432", "some-postgres")

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		logger.Panicf("cant listen port %v", err)
	}

	server := grpc.NewServer()

	comments_microservice_client.RegisterCommentsManagerServer(server, database)

	logger.Info("starting server at: ", 8082)
	logger.Panic(server.Serve(lis))

	defer database.Close()
	defer l.Sync()
	defer comments_microservice.Cm.Close()
}
