.PHONY: fetch
SHELL := /bin/bash

run:
	PORT=8888 go run main.go

test:
	PORT=8888 go test ./...

coverage:
	go test -v -coverprofile tests/coverage.out ./... && go tool cover -html=tests/coverage.out

lint:
	golangci-lint run && echo "✅"
