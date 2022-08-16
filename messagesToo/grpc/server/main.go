package main

import (
	"log"
	"net"
	"os"

	messagesToocore "github.com/EvanFr/messagesToo/messagesToo/core"
	messagesToogrpc "github.com/EvanFr/messagesToo/messagesToo/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// configure our core service
	clientService := messagesToocore.NewService()

	// configure our gRPC service controller
	clientServiceController := NewClientServiceController(clientService)

	// start a gRPC server
	server := grpc.NewServer()
	messagesToogrpc.RegisterClientServiceServer(server, clientServiceController)
	reflection.Register(server)

	con, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
	if err != nil {
		panic(err)
	}

	log.Printf("Starting gRPC user service on %s...\n", con.Addr().String())
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
