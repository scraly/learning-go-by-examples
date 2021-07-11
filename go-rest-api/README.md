# go-rest-api

This repo contains a simple/basic HTTP server in Go, with a basic code organization.
We use:
* net/http package to start and serve HTTP server
* Gorilla mux to handle routes
* Swagger in order to serve a REST API compliant with OpenAPI specs

## Pre-requisits

Install Go in 1.16 version minimum.

## Build the app

`$ go build -o bin/go-rest-api internal/main.go`

or

`$ make build`

## Run the app

`$ ./bin/go-rest-api`

## Test the app

```
$ curl http://localhost:8080/healthz
OK

$ curl http://localhost:8080/hello/aurelie

```

### Request & Response Examples

Swagger doc: [http-go-server](https://github.com/scraly/learning-by-examples/go-rest-api/doc/index.html)

|                 URL					 | Port | HTTP Method			       | Operation														    |
|:-------------------------:|:--------:|:-----------------------:|------------------------------------------------------------------------|
| /healthz							 | 8080 | GET       |  Test if the app is running							    |
| /hello/{name}							 | 8080 | GET       |  Returns message with {name} provided in the query							    |						    |


`$ curl localhost:8080/hello/aurelie`

## Generate swagger files

After editing `pkg/swagger/swagger.yml` file you need to generate swagger files again:

`$ task swagger.gen`

## Test swagger file validity

`$ task swagger.validate`

## Generate swagger documentation

`$ task swagger.doc`