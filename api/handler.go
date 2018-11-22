package api

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/k0kubun/pp"
)

// Owner represents owner object
type Owner struct {
	Address  string
	AsNumber uint32
	Objects  Object
}

// Object represents each routing object
type Object struct {
	AutNum string
	AsSet  string
	Route  string
	Route6 string
}

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

	filename := in.Filename
	data := in.Content
	objectType := in.Type

	pp.Println(filename, data)

	switch objectType {
	case "autnum":
		fmt.Println("autnum")
		// HandleAutNum
	case "route":
		fmt.Println("route")
		// HandleRoute
	case "route6":
		fmt.Println("route6")
		// HandleRoute6
	case "asset":
		fmt.Println("as-set")
		// HandleAsSet
	}

	sh := shell.NewShell(Config.Daemon.Ipfs.Host)

	cid, err := sh.Add(bytes.NewReader(data))
	if err != nil {
		fmt.Fprintln(os.Stderr, "sh.Add:", err)
		return &HashValue{Filename: in.Filename}, err
	}

	return &HashValue{Filename: in.Filename, Hash: cid}, nil
}

func (s *Server) GetFile(ctx context.Context, in *GetFileMessage) (*ContentMessage, error) {
	log.Printf("Receive GetFile %s\n", in.Hash)

	// TODO: Modity function to call from Ethereum Address, not IPFS HashValue.
	sh := shell.NewShell(Config.Daemon.Ipfs.Host)

	obj, err := sh.ObjectGet(in.Hash)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(obj)

	return &ContentMessage{Hash: in.Hash, Content: obj.Data}, nil
}
