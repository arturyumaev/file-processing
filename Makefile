#!make

all: build run

build:
	go build -o ./bin/api ./cmd/api/main.go

run:
	./bin/api

run_local:
	CompileDaemon -build="go build -o ./bin/api ./cmd/api/main.go" -command="./bin/api" -exclude-dir=".git" -color -log-prefix=false

migrate:
	goose -v -dir migrations up

rollback:
	goose -v -dir migrations reset

swagger:
	rm -rf docs
	swag init -g cmd/api/main.go

test:
	go test ./...
