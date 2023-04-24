all: docker_compose

run_local:
	CompileDaemon -build="go build -o ./bin/api ./cmd/api/main.go" -command="./bin/api" -exclude-dir=".git" -color -log-prefix=false

swagger:
	rm -rf docs
	swag init -g cmd/api/main.go

docker_compose:
	docker compose --env-file ./.env up --build api

test:
# generate mocks
	go generate ./...
# run tests
	go test -coverprofile cover.out ./...
	go tool cover -html cover.out
