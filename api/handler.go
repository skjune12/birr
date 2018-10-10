package api

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

// Server represents the gRPC Server
type Server struct {
}

// SayHello generates response to a Ping Request
func (s *Server) Ping(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s\n", in.Ping)

	return &PingMessage{Ping: "pong"}, nil
}

func (s *Server) AddFile(ctx context.Context, in *AddFileMessage) (*HashValue, error) {
	log.Printf("Receive message %s\n", in.Filename)

	data, err := ioutil.ReadFile(in.Filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ioutil.ReadFile:", err)
		return &HashValue{Filename: in.Filename}, err
	}

	sh := shell.NewShell("localhost:5001")

	cid, err := sh.Add(bytes.NewReader(data))
	if err != nil {
		fmt.Fprintln(os.Stderr, "sh.Add:", err)
		return &HashValue{Filename: in.Filename}, err
	}

	return &HashValue{Filename: in.Filename, Hash: cid}, nil
}

func (s *Server) GetFile(ctx context.Context, in *GetFileMessage) (*ContentMessage, error) {
	log.Printf("Receive GetFile %s\n", in.Hash)

	sh := shell.NewShell("localhost:5001")

	obj, err := sh.ObjectGet(in.Hash)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(obj)

	return &ContentMessage{Hash: in.Hash, Content: obj.Data}, nil
}
