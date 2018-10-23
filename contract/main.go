package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"os/exec"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func setupContract() error {
	if err := os.RemoveAll("SimpleStorage"); err != nil {
		return err
	}

	// ABIおよびbinの作成
	fmt.Println("solc SimpleStorage.sol --abi --bin -o SimpleStorage")
	out, err := exec.Command("solc", "SimpleStorage.sol", "--abi", "--bin", "-o", "SimpleStorage").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, out)
		return err
	}

	// contract.go の生成
	fmt.Println("abigen --sol=SimpleStorage.sol --pkg=main --out=contract.go")
	out, err = exec.Command("abigen", "--sol=SimpleStorage.sol", "--pkg=main", "--out=contract.go").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, out)
		return err
	}

	return nil
}

func main() {
	err := setupContract()
	if err != nil {
		log.Fatal(err)
	}

	// create a key and make a genesis account with a bunch of ether
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}

	// initiate the simulated background
	// backends.NewSimulatedBackend creates a new binding backend using simulated blockchain for testing purposes.
	sim := backends.NewSimulatedBackend(alloc, uint64(133700000))

	// Deploy Contract
	addr, _, contract, err := DeploySimpleStorage(auth, sim)
	if err != nil {
		log.Fatalf("Could not deploy contract: %v", err)
	}

	fmt.Printf("Contract deployed to %s\n", addr.String())

	// コントラクト内部のデータの読み出し
	storedData, _ := contract.StoredData(nil)
	fmt.Printf("Post-mining Stored Data: %v\n", storedData)

	// コントラクトへのデータの書き込み
	fmt.Println("Adding new object...")
	contract.Set(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: uint64(2381623),
		Value:    big.NewInt(1000),
	}, "Hello")

	// マイニング
	fmt.Println("Mining...")
	sim.Commit()

	// コントラクト内部のデータの読み出し
	storedData, _ = contract.StoredData(nil)
	fmt.Printf("Post-mining Stored Data: %v\n", storedData)

}
