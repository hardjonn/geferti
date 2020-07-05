.PHONY: build
build:
	go build -v ./cmd/apiserver

run:
	./apiserver

start: build run

test:
	go test ./...

migration.new:
	docker-compose run --rm migrate create -ext sql -dir /migrations $(name)

migration.up:
	docker-compose run --rm migrate -path /migrations -database "mysql://gefertiapp:password@tcp(db:3306)/geferti" up $(steps)

migration.down:
	docker-compose run --rm migrate -path /migrations -database "mysql://gefertiapp:password@tcp(db:3306)/geferti" down $(steps)

.DEFAULT_GOAL := build
