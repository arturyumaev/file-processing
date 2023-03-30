#!make
-include .env
export

all: clean rollback migrate build run

build:
	go build -o ./bin/api ./cmd/api/main.go

run:
	./bin/api -config ./config/config.yaml

run_local:
	CompileDaemon -build="go build -o ./bin/api ./cmd/api/main.go" -command="./bin/api -config ./config/config.yaml" -exclude-dir=".git" -color -log-prefix=false

migrate:
	goose -v -dir migrations up

rollback:
	goose -v -dir migrations reset

clean:
	go clean

swagger:
	rm -rf docs
	swag init -g cmd/api/main.go

createdb:
	docker pull postgres
	docker run -p 9999:5432 --name postgres -e POSTGRES_PASSWORD=postgres -d postgres

test:
	go test -v -coverprofile cover.out ./...
