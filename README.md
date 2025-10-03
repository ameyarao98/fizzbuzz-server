# fizzbuzz-server

**Go web server providing customizable FizzBuzz sequences with request statistics.**
The server uses Redis for storing request counts and supports Docker-based setup for easy development and testing.

---

## Features

- Generate **customizable FizzBuzz sequences** by specifying integer multiples and replacement strings.
- **Track request statistics** in Redis (e.g., most frequent request).

---

## Requirements

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Yaak](https://github.com/yaak-org/yaak) (optional, to test and explore the API)

---

## Setup
```bash
docker-compose up
```

This will start both the Go server and a Redis instance. The server runs at: http://localhost:8000

### Running Tests

To run all tests inside the Docker container:
```bash
docker-compose exec server go test -p 1 ./...
```
-p 1 ensures tests run sequentially, avoiding race conditions with Redis.

### Using Yaak

Open the Yaak GUI, load this project’s configuration, and you can interactively test the API endpoints.
