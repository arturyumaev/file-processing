include .env

all: clean migrate build run

build:
	go build -o ./bin/api ./cmd/api/main.go

run:
	./bin/api -config ./config/config.yaml

run_local:
	CompileDaemon -build="go build -o ./bin/api ./cmd/api/main.go" -command="./bin/api -config ./config/config.yaml" -exclude-dir=".git" -color -log-prefix=false

migrate:
	goose -v -dir ./migrations up

rollback:
	goose -v -dir ./migrations reset

clean:
	go clean

test:
	go test ./...
