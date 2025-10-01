FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o fizzbuzz-server ./server
CMD ["./fizzbuzz-server"]
