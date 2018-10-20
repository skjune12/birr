// Copyright © 2018 Kohei Suzuki <jingle@sfc.wide.ad.jp>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/skjune12/birr/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the content from IPFS Hash value.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) != 3 {
			fmt.Printf("usage: %s %s `IPFS_HASH`\n", os.Args[0], os.Args[1])
			os.Exit(0)
		}

		ipfsHash := os.Args[2]

		conn, err := grpc.Dial(":7777", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("grpc.Dial: %s\n", err)
		}
		defer conn.Close()

		c := api.NewBirrClient(conn)

		response, err := c.GetFile(context.Background(), &api.GetFileMessage{Hash: ipfsHash})
		if err != nil {
			log.Fatalf("Error when calling GetFile: %s\n", err)
		}

		fmt.Printf("Response from server\nHash: %s\nContent: %s", response.Hash, response.Content)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
