package main

import (
	"fmt"
	"log"
	"net"

	"github.com/skjune12/birr/api"
	"google.golang.org/grpc"
)

func main() {
	address := fmt.Sprintf("%s:%s", api.Config.Daemon.Host, api.Config.Daemon.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server listen on port :%s\n", api.Config.Daemon.Port)

	// create a server instance
	s := api.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Birr service to the server
	api.RegisterBirrServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
