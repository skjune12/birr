package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	sh := shell.NewShell("localhost:5001")

	if len(os.Args) < 2 {
		showUsage()
	}

	switch os.Args[1] {
	case "add":
		filename := os.Args[2]
		data, err := readFile(filename)

		cid, err := sh.Add(bytes.NewReader(data))

		if err != nil {
			fmt.Fprintf(os.Stderr, "err")
		}

		log.Println("added:", cid)

	case "get":
		ipfsHash := os.Args[2]

		sh := shell.NewShell("localhost:5001")

		obj, err := sh.ObjectGet(ipfsHash)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(obj.Data)

	default:
		showUsage()
	}
}

func showUsage() {
	fmt.Fprintf(os.Stderr, "usage: %s TYPE\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "where  TYPE := { add | get } COMMAND\n")

	os.Exit(1)
}

func readFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, err
}
