ARG GO_VERSION=1.22.5
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o main cmd/api/main.go

FROM debian:bookworm-slim

COPY --from=builder /usr/src/app/main /usr/local/bin/main

EXPOSE 8080

CMD ["/usr/local/bin/main"]
