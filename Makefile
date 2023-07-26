.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet


server:
	go run cmd/api/server.go
.PHONY:server

test:
	go test ./...
.PHONY:test
