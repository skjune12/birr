package main

import (
	"fmt"
	"log"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/skjune12/birr/contract"
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
	addr, _, birrContract, err := contract.DeployBirrContract(auth, sim)
	if err != nil {
		log.Fatalf("Could not deploy contract: %v", err)
	}

	fmt.Printf("Contract deployed to %s\n", addr.String())

	// verity the transaction
	sim.Commit()

	addrLen, err := birrContract.NumOfAddrs(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addrLen)

	sim.Commit()

	_, err = birrContract.NewObject(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, 2500)
	if err != nil {
		log.Fatal(err)
	}

	sim.Commit()

	addrLen, err = birrContract.NumOfAddrs(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addrLen)

	owner, _ := birrContract.AddressList(nil, big.NewInt(0))
	log.Println(reflect.TypeOf(owner))
	fmt.Println(owner.String())

	asNumber, _, _, _, _, err := birrContract.GetObject(nil, owner)
	fmt.Println(asNumber)
}
