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

	"github.com/skjune12/birr/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Check the connectivity to birrd.",
	Run: func(cmd *cobra.Command, args []string) {
		address := fmt.Sprintf("%s:%s", cfg.Client.Host, cfg.Client.Port)
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("grpc.Dial: %s\n", err)
		}
		defer conn.Close()

		c := api.NewBirrClient(conn)

		response, err := c.Ping(context.Background(), &api.PingMessage{Ping: "Ping"})
		if err != nil {
			log.Fatalf("Error when calling Ping: %s\n", err)
		}
		log.Printf("Response from server: %s", response.Ping)
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
