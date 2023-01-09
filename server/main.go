package main

import (
	"fmt"
	accountV1 "github.com/materials-resources/Service-Prophet/gen/proto/go/prophet-API/account/v1"
	inventoryV1 "github.com/materials-resources/Service-Prophet/gen/proto/go/prophet-API/inventory/v1"
	orderV1 "github.com/materials-resources/Service-Prophet/gen/proto/go/prophet-API/order/v1"
	"github.com/materials-resources/Service-Prophet/internal/services/account"
	"github.com/materials-resources/Service-Prophet/internal/services/inventory"
	"github.com/materials-resources/Service-Prophet/internal/services/order"
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
	inventoryV1.RegisterInventoryServiceServer(server, &inventory.Server{DB: database})
	orderV1.RegisterOrderServiceServer(server, &order.Server{DB: database})

	reflection.Register(server)

	log.Println("Listening on", listenOn)

	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}
	return nil
}
