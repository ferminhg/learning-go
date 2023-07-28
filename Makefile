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

dev:
	go run cmd/api/server.go
.PHONY:devserver

test:
	go test ./...
.PHONY:test

build: test
	go build -o build/server cmd/api/server.go
.PHONE:build
