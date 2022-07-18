all: build 

build:
	go build -v ./cmd/apiserver

run:
	go run ./cmd/apiserver
test:
	go test -v -race ./...
.DEFAULY_GOAL := build