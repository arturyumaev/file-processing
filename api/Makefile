build:
	go build -o ./bin/api ./cmd/api/main.go

run:
	./bin/api -config ./config/config.yaml

run_local:
	CompileDaemon -build="go build -o ./bin/api ./cmd/api/main.go" -command="./bin/api -config ./config/config.yaml" -color
