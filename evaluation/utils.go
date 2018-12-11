package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mr-tron/base58"
)

// MultiHash represents base58 decoded IPFS Hash (34bytes length in total)
type MultiHash struct {
	Digest       [32]byte
	HashFunction byte
	Size         byte
}

// GetMultiHashFromIPFSHash generates MultiHash from IPFSHash
func GetMultiHashFromIPFSHash(ipfsHash string) (*MultiHash, error) {
	decoded, err := base58.Decode(ipfsHash)
	if err != nil {
		return nil, err
	}

	digest := [32]byte{}
	copy(digest[:], decoded[2:])

	hashFunction := decoded[0]
	size := decoded[1]

	multihash := MultiHash{}
	multihash.Digest = digest
	multihash.HashFunction = hashFunction
	multihash.Size = size

	return &multihash, nil
}

// GetIPFSHashFromMultiHash generates IPFSHash from MultiHash
func GetIPFSHashFromMultiHash(multihash *MultiHash) string {
	// TODO: concat MultiHash values
	bytes := make([]byte, 34)

	bytes[0] = multihash.HashFunction
	bytes[1] = multihash.Size
	for i, v := range multihash.Digest {
		bytes[i+2] = v
	}

	encoded := base58.Encode(bytes)
	return encoded
}

// StringToBytes32 converts string to [32]bytes
func StringToBytes32(str string) [32]byte {
	bytes := [32]byte{}
	copy(bytes[:], []byte(str))
	return bytes
}

// Bytes32ToString converts [32]bytes to string
func Bytes32ToString(bytes [32]byte) string {
	return fmt.Sprintf("%s", bytes[:])
}

// ShowUsage shows how to use this script
func ShowUsage() {
	fmt.Fprintf(os.Stderr, "usage: %s TYPE\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "where  TYPE := { deploy | add | get | remove | kill } COMMAND\n")

	os.Exit(1)
}

// ReadFile reads file from argument and returns its content
func ReadFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, err
}
