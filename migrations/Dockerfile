# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

CMD ["make"]
