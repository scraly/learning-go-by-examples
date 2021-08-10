# go-gopher-grpc

[![Go Report Card](https://goreportcard.com/badge/github.com/scraly/learning-go-by-examples)](https://goreportcard.com/report/github.com/scraly/learning-go-by-examples)

This repo contains a simple CLI (Command Line Interface) application with a gRPC server in Go, with a basic code organization.
We use:
* net/http package to retrieve our cute Gophers
* Cobra for creating powerful modern CLI applications
* Viper to ...
* ...

go-gopher-grpc use [Taskfile](https://dev.to/stack-labs/introduction-to-taskfile-a-makefile-alternative-h92) (a Makefile alternative). 

Please read the [Learning Go by examples: part 6 - xx](xx) article in order to know more about this Git repository.

## Pre-requisites

Install Go in 1.16 version minimum.

## Generate protoc files

`$ task generate`

## Build the app

`$ go build -o bin/go-gopher-grpc main.go`

or

`$ task build`

## Run the app

`$ ./bin/go-gopher-grpc server`

or

`$ task run`

## Test the app

Usage:
```
$ ./bin/gopher-grpc
gRPC application written in Go.

Usage:
  go-gopher-grpc [command]

Available Commands:
  client      Query the gRPC server
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  server      Starts the Schema gRPC server

Flags:
      --config string   config file (default is $HOME/.go-gopher-grpc.yaml)
  -h, --help            help for go-gopher-grpc
  -t, --toggle          Help message for toggle

Use "go-gopher-grpc [command] --help" for more information about a command.
```

Run the gRPC server:

```
./bin/gopher-grpc server
task: [build] GOFLAGS=-mod=mod go build -o bin/gopher-grpc main.go
2021/08/09 14:22:39 GRPC server listening on [::]:9000
```

In another tab of your terminal, run the gRPC client:

```
$ ./bin/gopher-grpc client gandalf
2021/08/09 14:22:42 URL: https://raw.githubusercontent.com/scraly/gophers/main/gandalf.png
```

Cool! 
you should have one log file in your server, like this:
```
2021/08/09 14:22:41 Received: gandalf
```