package main

import (
	"log"
	"net"

	"github.com/skjune12/birr/api"
	"google.golang.org/grpc"
)

func main() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", ":7777")
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
