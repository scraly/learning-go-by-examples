# go-rest-api

This repo contains a simple/basic HTTP server in Go, with a basic code organization.
We use:
* net/http package to start and serve HTTP server
* Gorilla mux to handle routes
* Swagger in order to serve a REST API compliant with OpenAPI specs

go-rest-api use [Taskfile](https://dev.to/stack-labs/introduction-to-taskfile-a-makefile-alternative-h92) (a Makefile alternative). 

Please read the [Learning Go by examples: part 3 - Create an HTTP REST API Server in Go](https://dev.to/aurelievache/learning-go-by-examples-part-2-create-an-http-rest-api-server-in-go-1cdm) article in order to know more about this Git repository.

## Pre-requisits

Install Go in 1.16 version minimum.

## Build the app

`$ go build -o bin/go-rest-api internal/main.go`

or

`$ task build`

## Run the app

`$ ./bin/go-rest-api`

or

`$ task run`

## Test the app

```
$ curl http://localhost:8080/healthz
OK

$ curl http://localhost:8080/hello/aurelie
"Hello aurelie!"

$ curl -O localhost:8080/gopher/dr-who
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  992k    0  992k    0     0  9185k      0 --:--:-- --:--:-- --:--:-- 9185k

$ file dr-who
dr-who: PNG image data, 1700 x 1460, 8-bit/color RGBA, non-interlaced
```

### Request & Response Examples

Swagger doc: [go-rest-api](https://github.com/scraly/learning-go-by-examples/blob/main/go-rest-api/doc/index.html)

|                 URL					 | Port | HTTP Method			       | Operation														    |
|:-------------------------:|:--------:|:-----------------------:|------------------------------------------------------------------------|
| /healthz							 | 8080 | GET       |  Test if the app is running							    |
| /hello/{name}							 | 8080 | GET       |  Returns message with {name} provided in the query							    |						    |
| /gopher/{name}							 | 8080 | GET       |  Returns gopher image by {name} provided in the query							    |						    |

`$ curl localhost:8080/hello/aurelie`

## Generate swagger files

After editing `pkg/swagger/swagger.yml` file you need to generate swagger files again:

`$ task swagger.gen`

## Test swagger file validity

`$ task swagger.validate`

## Generate swagger documentation

`$ task swagger.doc`