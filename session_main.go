package main

import (
	"google.golang.org/grpc"
	"kino_backend/logger"
	"kino_backend/session_microservice"
	"kino_backend/session_microservice_client"
	"net"
)

func main() {
	l := logger.InitLogger()

	//база для redis
	redis := session_microservice.ConnectSessionDB("user@redis:6379", "0")

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Panicf("cant listen port %v", err)
	}

	server := grpc.NewServer()

	session_microservice_client.RegisterSessionManagerServer(server, redis)

	logger.Info("starting server at: ", 8081)
	logger.Panic(server.Serve(lis))

	defer redis.Close()
	defer l.Sync()
	defer session_microservice.Sm.Close()
}
