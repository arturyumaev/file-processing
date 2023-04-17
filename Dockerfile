# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest
COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE ${APPLICATION_PORT}

# Run
CMD ["make", "start"]
