all: build 

build:
	go build -v ./cmd/apiserver

run:
	go run ./cmd/apiserver
.DEFAULY_GOAL := build