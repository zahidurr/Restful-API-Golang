check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models

generate_client:
	cd sdk && swagger generate client -f ../swagger.yaml -A restfulapi

## Build all binaries
build:
	go build -o bin/restfulapi main.go

## Run main
run:
	go run main.go

mod init:
	go mod init github.com/zahidurr/Restful-API-Golang