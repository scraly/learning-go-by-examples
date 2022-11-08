# go-gopher-opentelemetry

[![Go Report Card](https://goreportcard.com/badge/github.com/scraly/learning-go-by-examples)](https://goreportcard.com/report/github.com/scraly/learning-go-by-examples)

This repo contains a simple application with instrumentation, with a basic code organization.
We use:
* OpenTelemetry SDK
* a Jaeger instance

go-gopher-opentelemetry use [Taskfile](https://dev.to/stack-labs/introduction-to-taskfile-a-makefile-alternative-h92) (a Makefile alternative). 

Please read the [Learning Go by examples: part 10 - Instrument your Go app with OpenTelemetry and send traces to Jaeger - Distributed Tracing](xx) article in order to know more about this Git repository.

## Pre-requisites

Install Go in 1.19 version minimum.

or:

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/scraly/learning-go-by-examples/tree/main/go-gopher-opentelemetry)


## Build the app

`$ go build -o bin/go-gopher-opentelemetry main.go`

or

`$ task build`

## Run the app

`$ ./bin/go-gopher-opentelemetry`

or

`$ task run`

## Test the app

Usage:

```
$ export MY_NAME=scraly ; go run main.go
2022/03/01 09:56:21 Listening on localhost:8080
```

```
$ curl localhost:8080
```