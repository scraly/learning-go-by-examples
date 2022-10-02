FROM golang:1.18.4-alpine AS build

COPY . /build
WORKDIR /build
RUN GOFLAGS=-mod=mod go build -o bin/go-rest-api internal/main.go

FROM alpine AS final

WORKDIR /app
COPY --from=build /build/bin/go-rest-api /app

CMD ["./go-rest-api"]
