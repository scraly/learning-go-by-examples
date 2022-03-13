# go-gopher-cli

[![Go Report Card](https://goreportcard.com/badge/github.com/scraly/learning-go-by-examples)](https://goreportcard.com/report/github.com/scraly/learning-go-by-examples)

This repo contains a simple CLI (Command Line Interface) application in Go, with a basic code organization.
We use:
* net/http package to retrieve our cute Gophers
* Cobra for creating powerful modern CLI applications
* Viper to ...

go-gopher-cli use [Taskfile](https://dev.to/stack-labs/introduction-to-taskfile-a-makefile-alternative-h92) (a Makefile alternative). 

Please read the [Learning Go by examples: part 3 - Create a CLI app in Go](https://dev.to/aurelievache/learning-go-by-examples-part-3-create-a-cli-app-in-go-1h43) article in order to know more about this Git repository.

## Pre-requisits

Install Go in 1.16 version minimum.

or:

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/scraly/learning-go-by-examples/tree/main/go-gopher-cli)

## Build the app

`$ go build -o bin/go-gopher-cli main.go`

or

`$ task build`

## Run the app

`$ ./bin/go-gopher-cli`

or

`$ task run`

## Test the app

```
$ ./bin/gopher-cli
Gopher CLI application written in Go.

Usage:
  go-gopher-cli [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  get         This command will get the desired Gopher
  help        Help about any command
...


$ ./bin/gopher-cli get friends
Try to get 'friends' Gopher...
Perfect! Just saved in friends.png!

$ file friends.png
friends.png: PNG image data, 1156 x 882, 8-bit/color RGBA, non-interlaced
```

or

```
$ task run -- [command]

$ task run -- get friends
task: [run] GOFLAGS=-mod=mod go run main.go get friends
Try to get 'friends' Gopher...
Perfect! Just saved in friends.png!


$ file friends.png
friends.png: PNG image data, 1838 x 1408, 8-bit/color RGBA, non-interlaced
```