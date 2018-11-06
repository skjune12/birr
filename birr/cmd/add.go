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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/skjune12/birr/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add file to birrd",
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) != 4 {
			fmt.Printf("usage: %s %s TYPE `filename`\n\n", os.Args[0], os.Args[1])
			fmt.Printf("where  TYPE := { autnum | route | route6 | asset }\n")
			os.Exit(0)
		}

		// load file
		filename, err := filepath.Abs(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// open file
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ioutil.ReadFile:", err)
			os.Exit(1)
			// return &HashValue{Filename: in.Filename}, err
		}

		address := fmt.Sprintf("%s:%s", cfg.Client.Host, cfg.Client.Port)
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("grpc.Dial: %s\n", err)
		}
		defer conn.Close()

		c := api.NewBirrClient(conn)

		objectType := os.Args[2]

		switch os.Args[2] {
		case "autnum":
			fmt.Println("autnum")
		case "route":
			fmt.Println("route")
		case "route6":
			fmt.Println("route6")
		case "asset":
			fmt.Println("as-set")
		}

		response, err := c.AddFile(context.Background(), &api.AddFileMessage{Type: objectType, Filename: filename, Content: data})
		if err != nil {
			log.Fatalf("Error when calling AddFile: %s\n", err)
		}
		log.Printf("Response from server: %s %s", response.Filename, response.Hash)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
