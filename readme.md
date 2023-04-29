# SENAO PRETEST

## Prepare Dev Environment
* install golang 1.20
* git clone
* go mod init senao
* go mod tidy
* go install github.com/swaggo/swag/cmd/swag@v1.8.1
* go build -o ./cmd/senao ./cmd

## Run Reset DB
* ./cmd/senao resetdb

## Run server
* ./cmd/senao server
* http://localhost:8080/v1/docs/index.html


