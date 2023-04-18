all: docker_compose

build:
	go build -o ./bin/api ./cmd/api/main.go

run:
	./bin/api

run_local:
	CompileDaemon -build="go build -o ./bin/api ./cmd/api/main.go" -command="./bin/api" -exclude-dir=".git" -color -log-prefix=false

swagger:
	rm -rf docs
	swag init -g cmd/api/main.go

test:
	go test ./...

docker_compose:
	docker compose --env-file ./.env up --build api
