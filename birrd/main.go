package main

import (
	"fmt"
	"log"
	"net"

	"github.com/skjune12/birr/api"
	"google.golang.org/grpc"
)

func main() {
	config := api.LoadConfiguration()

	// create a listener on TCP port 7777
	address := fmt.Sprintf("%s:%s", config.Daemon.Host, config.Daemon.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server listen on port :7777")

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
