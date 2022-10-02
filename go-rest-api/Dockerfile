FROM golang:1.18.4-alpine AS build

COPY go-rest-api /build
WORKDIR /build
RUN go build

FROM alpine AS final

WORKDIR /app
COPY --from=build /bin/go-rest-api /app

CMD ["./go-rest-api"]
