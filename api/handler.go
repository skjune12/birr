package api

import (
	"context"
	"log"
)

// Server represents the gRPC Server
type Server struct {
}

// SayHello generates response to a Ping Request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s\n", in.Ping)

	return &PingMessage{Ping: "pong"}, nil
}
