all: production

swagger:
	rm -rf docs
	swag init -g cmd/api/main.go

development:
	docker compose -f docker-compose.development.yml --env-file ./.env.development up --build api

test:
	docker compose -f docker-compose.test.yml down
	rm -rf .database/data
	mkdir .database/data
	docker compose -f docker-compose.test.yml --env-file ./.env.test up --build api

production:
	docker compose -f docker-compose.production.yml --env-file ./.env.production up --build api

test_cover:
# generate mocks
	go generate ./...
# run tests
	go test -coverprofile cover.out ./...
	go tool cover -html cover.out
