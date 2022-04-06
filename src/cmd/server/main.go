package main

import (
	"log"
	"net"

	"github.com/jasonzou/ezgRPC/src/internal/server"
	"google.golang.org/grpc/reflection"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func main() {
	log.Println("Starting listening on port 8080")

	lis, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", CONN_PORT)
	srv := server.NewGRPCServer()
	log.Printf("created a new grpc server")
	// Register reflection service on gRPC server.
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
