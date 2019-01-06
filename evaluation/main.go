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

	"github.com/skjune12/birr/evaluation/contract"
)

// GasLimit specify the GasLimit. in Ropsten, it is defined as 4712388
const GasLimit uint = 4712388

func main() {
	if len(os.Args) < 2 {
		ShowUsage()
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
		log.Fatal("ethclient.Dial:", err)
	}

	switch os.Args[1] {
	// Deploying Ethereum Client
	case "deploy":
		privateKey, err := crypto.HexToECDSA(os.Getenv("ETH_SECRET_KEY"))
		if err != nil {
			log.Fatal("crypto.HexToECDSA:", err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal("client.PendingNonceAt:", err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal("client.SuggestGasPrice:", err)
		}

		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0) // in wei
		auth.GasLimit = uint64(GasLimit)
		auth.GasPrice = gasPrice

		input := "1.0"
		address, tx, _, err := contract.DeployBirrContract(auth, client, input)
		if err != nil {
			log.Fatal("contract.DeployStore:", err)
		}

		fmt.Println("Contract", address.Hex())
		fmt.Println("Transaction", tx.Hash().Hex())

	// Add content to IPFS and write the IPFS hash to Smart Contract
	case "add":
		// Add to IPFS
		if len(os.Args) < 5 {
			fmt.Printf("Usage: %s %s { route | route6 | aut-num | as-set } FILENAME KEY\n", os.Args[0], os.Args[1])
			os.Exit(1)
		}

		objectType := os.Args[2]
		filename := os.Args[3]
		key := StringToBytes32(os.Args[4])

		data, err := ReadFile(filename)
		if err != nil {
			log.Fatal("ReadFile:", err)
		}

		cid, err := sh.Add(bytes.NewReader(data))

		if err != nil {
			log.Fatal("sh.Add:", err)
		}

		log.Println("IPFS Hash", cid)

		// Add to Ethereum
		privateKey, err := crypto.HexToECDSA(os.Getenv("ETH_SECRET_KEY")) // pass the string
		if err != nil {
			log.Fatal("crypto.HexToECDSA:", err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal("PendingNonceAt:", err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal("client.SuggestGasPrice:", err)
		}

		auth := bind.NewKeyedTransactor(privateKey)

		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(GasLimit)
		auth.GasPrice = gasPrice

		// contract address (string)
		address := common.HexToAddress(os.Getenv("CONTRACT_ADDR"))
		instance, err := contract.NewBirrContract(address, client)
		if err != nil {
			log.Fatal("contract.NewBirrContract:", err)
		}

		multihash, err := GetMultiHashFromIPFSHash(cid)
		if err != nil {
			log.Fatal("GetMultiHashFromIPFSHash:", err)
		}

		var tx *types.Transaction

		switch objectType {
		case "route":
			tx, err = instance.SetRoute(auth, key, multihash.Digest, multihash.HashFunction, multihash.Size)

		case "route6":
			tx, err = instance.SetRoute6(auth, key, multihash.Digest, multihash.HashFunction, multihash.Size)

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
		if len(os.Args) < 5 {
			if os.Args[2] != "routelist" {
				fmt.Printf("Usage: %s %s { route | routelist | route6 | aut-num | as-set } ACCOUNT_ADDRESS KEY\n", os.Args[0], os.Args[1])
				os.Exit(1)
			}
		}

		objectType := os.Args[2]

		var accountAddr string
		var key [32]byte

		if os.Args[2] != "routelist" {
			accountAddr = os.Args[3]
			key = StringToBytes32(os.Args[4])
		}

		address := common.HexToAddress(os.Getenv("CONTRACT_ADDR"))
		instance, err := contract.NewBirrContract(address, client)
		if err != nil {
			log.Fatal("contract.NewBirrContract:", err)
		}

		var item struct {
			Digest       [32]byte
			HashFunction uint8
			Size         uint8
		}

		switch objectType {
		case "route":
			item, err = instance.GetRoute(nil, common.HexToAddress(accountAddr), key)

		case "routelist":
			// NOTE: does not work but works fine in remix
			keys, err := instance.GetRouteKeys(nil)
			if err != nil {
				log.Fatal("contract.GetRouteKeys", err)
			}

			fmt.Printf("%#v", keys)
			return

		case "route6":
			item, err = instance.GetRoute6(nil, common.HexToAddress(accountAddr), key)

		case "route6list":
			keys, err := instance.GetRoute6Keys(nil)
			if err != nil {
				log.Fatal("contract.GetRoute6Keys", err)
			}

			fmt.Printf("%#v", keys)
			return

		case "aut-num":
			item, err = instance.GetAutNum(nil, common.HexToAddress(accountAddr))

		case "as-set":
			item, err = instance.GetAsSet(nil, common.HexToAddress(accountAddr))

		default:
			fmt.Fprintf(os.Stderr, "object type '%s' does not support", objectType)
			os.Exit(1)
		}

		if err != nil {
			log.Fatal("instance.Get", err)
		}

		multihash := &MultiHash{}
		multihash.Digest = item.Digest
		multihash.HashFunction = item.HashFunction
		multihash.Size = item.Size

		ipfsHash := GetIPFSHashFromMultiHash(multihash)

		fmt.Println("IPFS Hash:", ipfsHash)

		obj, err := sh.ObjectGet(ipfsHash)
		if err != nil {
			log.Println("sh.ObjectGet:", err)
		}

		fmt.Printf(obj.Data)

	// remove
	case "remove":
		if len(os.Args) < 4 {
			fmt.Printf("Usage: %s %s { route | route6 | aut-num | as-set } KEY\n", os.Args[0], os.Args[1])
			os.Exit(1)
		}

		objectType := os.Args[2]
		key := StringToBytes32(os.Args[3])

		// Exec contract
		privateKey, err := crypto.HexToECDSA(os.Getenv("ETH_SECRET_KEY")) // pass the string
		if err != nil {
			log.Fatal("crypto.HexToECDSA:", err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public ket to ECDSA")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal("client.PendingNonceAt", err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal("client.SuggestGasPrice:", err)
		}

		auth := bind.NewKeyedTransactor(privateKey)

		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(GasLimit)
		auth.GasPrice = gasPrice

		// contract address (string)
		address := common.HexToAddress(os.Getenv("CONTRACT_ADDR"))
		instance, err := contract.NewBirrContract(address, client)
		if err != nil {
			log.Fatal("contract.NewBirrContract", err)
		}

		var tx *types.Transaction

		switch objectType {
		case "route":
			tx, err = instance.RemoveRoute(auth, key)

		case "route6":
			tx, err = instance.RemoveRoute6(auth, key)

		default:
			fmt.Fprintf(os.Stderr, "object type '%s' does not support", objectType)
			os.Exit(1)
		}

		if err != nil {
			log.Fatal("instance.RemoveObject", err)
		}

		fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	// destruct the contract
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
			log.Fatal("client.PendingNonceAt:", err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal("client.SuggestGasPrice:", err)
		}

		auth := bind.NewKeyedTransactor(privateKey)

		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(GasLimit)
		auth.GasPrice = gasPrice

		// contract address (string)
		address := common.HexToAddress(os.Getenv("CONTRACT_ADDR"))
		instance, err := contract.NewBirrContract(address, client)
		if err != nil {
			log.Fatal("contract.NewBirrContract:", err)
		}

		tx, err := instance.Kill(auth)

		if err != nil {
			log.Fatal("instance.SetObject:", err)
		}

		fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	default:
		ShowUsage()
	}
}
