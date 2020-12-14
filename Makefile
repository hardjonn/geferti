.PHONY: build
%:
    @:

build:
	go build -v ./cmd/rest-server

run:
	./rest-server

start: build run

test:
	go test -v -race ./...

migration.make:
	go run ./cmd/migration make $(name)

migration.up:
	go run ./cmd/migration up --step=0

.DEFAULT_GOAL := build
