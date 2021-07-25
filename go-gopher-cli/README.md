# go-gopher-cli

[![Go Report Card](https://goreportcard.com/badge/github.com/scraly/learning-go-by-examples)](https://goreportcard.com/report/github.com/scraly/learning-go-by-examples)

This repo contains a simple CLI (Command Line Interface) application in Go, with a basic code organization.
We use:
* net/http package to retrieve our cute Gophers
* Cobra for creating powerful modern CLI applications
* Viper to ...

go-gopher-cli use [Taskfile](https://dev.to/stack-labs/introduction-to-taskfile-a-makefile-alternative-h92) (a Makefile alternative). 

## Pre-requisits

Install Go in 1.16 version minimum.

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
