package main

import (
	"fmt"
	accountV1 "github.com/materials-resources/Service-Prophet/gen/proto/go/prophet-API/account/v1"
	"github.com/materials-resources/Service-Prophet/internal/services/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
func run() error {
	listenOn := "0.0.0.0:50057"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to start service on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()

	database := connectBun()

	accountV1.RegisterAccountServiceServer(server, &account.Server{DB: database})

	reflection.Register(server)

	log.Println("Listening on", listenOn)

	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}
	return nil
}
