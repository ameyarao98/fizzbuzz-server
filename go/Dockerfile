FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY server /app/server

CMD ["go", "run", "server/server.go"]
