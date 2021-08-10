/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "github.com/scraly/learning-go-by-examples/go-gopher-grpc/pkg/gopher"
	"github.com/spf13/cobra"
)

const (
	address     = "localhost:9000"
	defaultName = "dr-who"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Query the gRPC server",

	Run: func(cmd *cobra.Command, args []string) {
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		client := pb.NewGopherClient(conn)

		var name string

		// Contact the server and print out its response.
		// name := defaultName
		if len(os.Args) > 2 {
			name = os.Args[2]
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.GetGopher(ctx, &pb.GopherRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("URL: %s", r.GetMessage())
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
