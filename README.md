# file-processing
REST written in Go for file processing

# postgres

Create postgres container and forward port to :9999
```
$ docker pull postgres
$ docker run -p 9999:5432 --name postgres -e POSTGRES_PASSWORD=postgres -d postgres
```

# migrations

Download migrations utility
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Apply migrations from `postgres` directory
```
$ ./migrate up
```

Rollback migrations
```
$ ./migrate down
```
