package main

import (
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"kino_backend/logger"
	"kino_backend/session_microservice"
	"kino_backend/session_microservice_client"
	"net"
	"time"
)

func main() {

	l := logger.InitLogger()
	dbConnStr := flag.String("db_connstr", "user@localhost:6379", "redis connection string")
	dbName := flag.String("db_name", "0", "redis database name")
	flag.Parse()

	//база для redis
	//redis := session_microservice.ConnectSessionDB("user@redis:6379", "0")

	redis := session_microservice.NewSessionManager(*dbConnStr, *dbName)
	defer redis.Close()

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Panicf("cant listen port %v", err)
	}
	//CP := keepalive.ClientParameters{
	//	Time:                2 * time.Minute,
	//	Timeout:             350,
	//	PermitWithoutStream: true,
	//}
	//
	EP := keepalive.EnforcementPolicy{
		MinTime:             1 * time.Minute,
		PermitWithoutStream: true,
	}

	opts := []grpc.ServerOption{grpc.KeepaliveEnforcementPolicy(EP)}

	server := grpc.NewServer(opts...)

	session_microservice_client.RegisterSessionManagerServer(server, redis)

	logger.Info("starting server at: ", 8081)
	logger.Panic(server.Serve(lis))

	//defer redis.Close()
	defer l.Sync()
	//defer session_microservice.Sm.Close()
}
