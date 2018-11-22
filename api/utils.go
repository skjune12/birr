package api

import (
	"bytes"
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/k0kubun/pp"
	"github.com/spf13/viper"
)

type BIRRConfig struct {
	Daemon DaemonConfig `mapstructure:"daemon"`
}

type IpfsConfig struct {
	Host string `mapstructure:"host"`
}

type EthereumConfig struct {
	Host string `mapstructure:"host"`
}

type DaemonConfig struct {
	Host     string         `mapstructure:"host"`
	Port     string         `mapstructure:"port"`
	Ethrerum EthereumConfig `mapstructure:"ethereum"`
	Ipfs     IpfsConfig     `mapstructure:"ipfs"`
}

func LoadConfiguration() *BIRRConfig {
	viper := viper.New()
	viper.SetConfigName("birr")
	viper.AddConfigPath("$GOPATH/src/github.com/skjune12/birr/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Couldn't load config:", err)
		os.Exit(1)
	}

	var config BIRRConfig
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Couldn't load config:", err)
		os.Exit(1)
	}

	return &config
}

func addToIPFS(data []byte) (string, error) {
	sh := shell.NewShell(Config.Daemon.Ipfs.Host)

	cid, err := sh.Add(bytes.NewReader(data))
	if err != nil {
		fmt.Fprintln(os.Stderr, "sh.Add:", err)
		return "", err
	}

	return cid, nil
}

var Config *BIRRConfig

func init() {
	fmt.Fprintln(os.Stderr, "[INFO] Loading Configuration")
	Config = LoadConfiguration()
	pp.Println(Config)
	fmt.Fprintln(os.Stderr, "[INFO] Finish loading configuration")
}
