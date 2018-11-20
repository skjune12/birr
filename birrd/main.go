package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/skjune12/birr/api"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type Config struct {
	Daemon DaemonConfig `mapstructure:"daemon"`
}

type DaemonConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func LoadConfiguration() *Config {
	viper := viper.New()
	viper.SetConfigName("birr")
	viper.AddConfigPath("$GOPATH/src/github.com/skjune12/birr/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Couldn't load config:", err)
		os.Exit(1)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Couldn't load config:", err)
		os.Exit(1)
	}

	return &config
}

func main() {
	config := LoadConfiguration()

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
