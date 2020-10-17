.PHONY: build
build:
	go build -v ./cmd/apiserver

test:
	go test -v -race -timeout 5s ./...

.DEFAULT_GOAL := build