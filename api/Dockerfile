FROM golang:latest

ENV PROJECT_DIR=/api \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /api

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copying all the files
COPY . .

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build -o main cmd/api/main.go" -command="./main"
