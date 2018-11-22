package api

import (
	"bytes"
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/spf13/viper"
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

func addToIPFS(data []byte) (string, error) {
	sh := shell.NewShell("localhost:5001")

	cid, err := sh.Add(bytes.NewReader(data))
	if err != nil {
		fmt.Fprintln(os.Stderr, "sh.Add:", err)
		return "", err
	}

	return cid, nil
}
