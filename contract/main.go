package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// create a key and make a genesis account with a bunch of ether
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}

	// initiate the simulated background
	// backends.NewSimulatedBackend creates a new binding backend using simulated blockchain for testing purposes.
	sim := backends.NewSimulatedBackend(alloc, uint64(133700000))

	// Deploy Contract
	addr, _, contract, err := DeployBirrContract(auth, sim)
	if err != nil {
		log.Fatalf("Could not deploy contract: %v", err)
	}

	fmt.Printf("Contract deployed to %s\n", addr.String())

	// verity the transaction
	sim.Commit()

	_ = contract

	// NOTE: the example of reading from / writing to contract is below.

	// read the data from contract
	// data, _ := contract.GetHowManyOwners(nil)
	// fmt.Printf("Post-mining data: %v\n", data)
	//
	// asn, hash, ownerAddr, _ := contract.GetObjects(nil)
	// fmt.Printf("Post-mining data: asn=%v hash=%v ownerAddr=%v\n", asn, hash, ownerAddr)
	//
	// owner, _ := contract.GetOwner(nil, 1)
	// fmt.Printf("owner: %v\n", owner)

	// write a data to contract
	// fmt.Println("Adding new object...")
	// contract.Set(&bind.TransactOpts{
	// 	From:     auth.From,
	// 	Signer:   auth.Signer,
	// 	GasLimit: uint64(2381623),
	// 	Value:    big.NewInt(1000),
	// }, "Hello")
	//
	// mining
	// fmt.Println("Mining...")
	// sim.Commit()
	//
	// read the data from contract
	// storedData, _ = contract.StoredData(nil)
	// fmt.Printf("Post-mining Stored Data: %v\n", storedData)

}
