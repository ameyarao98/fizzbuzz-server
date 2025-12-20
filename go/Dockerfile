FROM golang:1.25-alpine

COPY . /app

WORKDIR /app

RUN go mod download

CMD ["go", "run", "server.go"]
