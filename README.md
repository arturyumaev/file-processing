# postgres

Create postgres container and forward port to :9999
```
make create_db
```

# dependencies

Download deps
```
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/golang/mock/mockgen@v1.6.0
```
