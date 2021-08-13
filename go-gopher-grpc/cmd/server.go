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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"

	pb "github.com/scraly/learning-go-by-examples/go-gopher-grpc/pkg/gopher"
	"google.golang.org/grpc"
)

const (
	port         = ":9000"
	KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"
)

// server is used to implement gopher.GopherServer.
type Server struct {
	pb.UnimplementedGopherServer
}

type Gopher struct {
	URL string `json: "url"`
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the Schema gRPC server",

	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()

		// Register services
		pb.RegisterGopherServer(grpcServer, &Server{})

		log.Printf("GRPC server listening on %v", lis.Addr())

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

// GetGopher implements gopher.GopherServer
func (s *Server) GetGopher(ctx context.Context, req *pb.GopherRequest) (*pb.GopherReply, error) {
	res := &pb.GopherReply{}

	// Check request
	if req == nil {
		fmt.Println("request must not be nil")
		return res, xerrors.Errorf("request must not be nil")
	}

	if req.Name == "" {
		fmt.Println("name must not be empty in the request")
		return res, xerrors.Errorf("name must not be empty in the request")
	}

	log.Printf("Received: %v", req.GetName())

	//Call KuteGo API in order to get Gopher's URL
	// https://kutego-api-xxxxx-ew.a.run.app/gophers?name=back-to-the-future
	// [{"name":"back-to-the-future","path":"back-to-the-future.png","url":"https://raw.githubusercontent.com/scraly/gophers/main/back-to-the-future.png"}]
	response, err := http.Get(KuteGoAPIURL + "/gophers?name=" + req.GetName())
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		// Transform our response to a []byte
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		// Put only needed informations of the JSON document in our array of Gopher
		var data []Gopher
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println(err)
		}

		// Create a string with all of the Gopher's name and a blank line as separator
		var gophers strings.Builder
		for _, gopher := range data {
			gophers.WriteString(gopher.URL + "\n")
		}

		res.Message = gophers.String()
	} else {
		fmt.Println("Error: Can't get the Gopher! :-(")
	}

	// res.Message = "Hello " + req.Name

	return res, nil
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
