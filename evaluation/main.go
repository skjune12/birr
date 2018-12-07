package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/joho/godotenv"

	"github.com/skjune12/eth-ipfs/contract"
)

// GasLimit specify the GasLimit. in Ropsten, it is defined as 4712388
const GasLimit uint = 4712388

func main() {
	if len(os.Args) < 2 {
		showUsage()
	}

	// load .env
	err := godotenv.Load()

	if err != nil {
		log.Println(".env is not specified.")
		if os.Getenv("ETH_SECRET_KEY") == "" {
			log.Fatal("ETH_SECRET_KEY is not specified.")
		}
		if os.Getenv("CONTRACT_ADDR") == "" {
			log.Fatal("CONTRACT_ADDR is not specified.")
		}
	}

	// setup ipfs shell
	sh := shell.NewShell("localhost:5001")
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	// Deploying Ethereum Client
	case "deploy":
		privateKey, err := crypto.HexToECDSA(os.Getenv("ETH_SECRET_KEY"))
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0) // in wei
		auth.GasLimit = uint64(GasLimit)
		auth.GasPrice = gasPrice

		input := "1.0"
		address, tx, _, err := contract.DeployStore(auth, client, input)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Contract", address.Hex())
		fmt.Println("Transaction", tx.Hash().Hex())

		// Add content to IPFS and write the IPFS hash to Smart Contract
	case "add":
		// Add to IPFS
		if len(os.Args) < 4 {
			fmt.Printf("Usage: %s %s { route | route6 | aut-num | as-set } FILENAME\n", os.Args[0], os.Args[1])
			os.Exit(1)
		}

		objectType := os.Args[2]
		filename := os.Args[3]

		data, err := readFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		cid, err := sh.Add(bytes.NewReader(data))

		if err != nil {
			fmt.Fprintf(os.Stderr, "err")
		}

		log.Println("IPFS Hash", cid)

		// Add to Ethereum
		privateKey, err := crypto.HexToECDSA(os.Getenv("ETH_SECRET_KEY")) // pass the string
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public ket to ECDSA")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		auth := bind.NewKeyedTransactor(privateKey)

		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(GasLimit)
		auth.GasPrice = gasPrice

		// contract address (string)
		address := common.HexToAddress(os.Getenv("CONTRACT_ADDR"))
		instance, err := contract.NewStore(address, client)
		if err != nil {
			log.Fatal(err)
		}

		multihash, err := GetMultiHashFromIPFSHash(cid)
		if err != nil {
			log.Fatal(err)
		}

		var tx *types.Transaction

		switch objectType {
		case "route":
			tx, err = instance.SetRoute(auth, multihash.Digest, multihash.HashFunction, multihash.Size)
		case "route6":
			tx, err = instance.SetRoute6(auth, multihash.Digest, multihash.HashFunction, multihash.Size)
		case "aut-num":
			tx, err = instance.SetAutNum(auth, multihash.Digest, multihash.HashFunction, multihash.Size)
		case "as-set":
			tx, err = instance.SetAsSet(auth, multihash.Digest, multihash.HashFunction, multihash.Size)
		default:
			fmt.Fprintf(os.Stderr, "object type '%s' does not support", objectType)
			os.Exit(1)
		}

		if err != nil {
			log.Fatal("instance.SetObject", err)
		}

		fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

		// get the value from IPFS
	case "get":
		if len(os.Args) < 4 {
			fmt.Printf("Usage: %s %s { route | route6 | aut-num | as-set } ACCOUNT_ADDRESS\n", os.Args[0], os.Args[1])
			os.Exit(1)
		}

		objectType := os.Args[2]
		accountAddr := os.Args[3]

		address := common.HexToAddress(os.Getenv("CONTRACT_ADDR"))
		instance, err := contract.NewStore(address, client)
		if err != nil {
			log.Fatal(err)
		}

		var item struct {
			Digest       [32]byte
			HashFunction uint8
			Size         uint8
		}

		switch objectType {
		case "route":
			item, err = instance.GetRoute(nil, common.HexToAddress(accountAddr))
		case "route6":
			item, err = instance.GetRoute6(nil, common.HexToAddress(accountAddr))
		case "aut-num":
			item, err = instance.GetAutNum(nil, common.HexToAddress(accountAddr))
		case "as-set":
			item, err = instance.GetAsSet(nil, common.HexToAddress(accountAddr))
		default:
			fmt.Fprintf(os.Stderr, "object type '%s' does not support", objectType)
			os.Exit(1)
		}

		if err != nil {
			log.Fatal(err)
		}

		multihash := &MultiHash{}
		multihash.Digest = item.Digest
		multihash.HashFunction = item.HashFunction
		multihash.Size = item.Size

		ipfsHash := GetIPFSHashFromMultiHash(multihash)

		fmt.Println("IPFS Hash:", ipfsHash)

		obj, err := sh.ObjectGet(ipfsHash)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf(obj.Data)

	case "kill":
		privateKey, err := crypto.HexToECDSA(os.Getenv("ETH_SECRET_KEY")) // pass the string
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public ket to ECDSA")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		auth := bind.NewKeyedTransactor(privateKey)

		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(GasLimit)
		auth.GasPrice = gasPrice

		// contract address (string)
		address := common.HexToAddress(os.Getenv("CONTRACT_ADDR"))
		instance, err := contract.NewStore(address, client)
		if err != nil {
			log.Fatal(err)
		}

		tx, err := instance.Kill(auth)

		if err != nil {
			log.Fatal("instance.SetObject", err)
		}

		fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	default:
		showUsage()
	}
}
