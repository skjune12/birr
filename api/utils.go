package api

import (
	"bytes"
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

func addToIPFS(data []byte) (string, error) {
	sh := shell.NewShell("localhost:5001")

	cid, err := sh.Add(bytes.NewReader(data))
	if err != nil {
		fmt.Fprintln(os.Stderr, "sh.Add:", err)
		return "", err
	}

	return cid, nil
}
