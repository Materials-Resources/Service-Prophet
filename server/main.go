package main

import (
	accountV1 "buf.build/gen/go/materials-resources/prophet/grpc/go/account/v1/accountv1grpc"
	inventoryV1 "buf.build/gen/go/materials-resources/prophet/grpc/go/inventory/v1/inventoryv1grpc"
	orderV1 "buf.build/gen/go/materials-resources/prophet/grpc/go/order/v1/orderv1grpc"
	"fmt"
	"github.com/materials-resources/Service-Prophet/core/services/account"
	"github.com/materials-resources/Service-Prophet/core/services/inventory"
	"github.com/materials-resources/Service-Prophet/core/services/order"
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
	listenOn := "0.0.0.0:50058"
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
