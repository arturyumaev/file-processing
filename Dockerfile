FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go/bin/api ./cmd/api/main.go

FROM scratch

COPY --from=builder /go/bin/api /go/bin/api

EXPOSE ${APPLICATION_PORT}

ENTRYPOINT ["/go/bin/api"]
