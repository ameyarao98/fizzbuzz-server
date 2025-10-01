package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ameyarao98/fizzbuzz-server/server/internal"
	"github.com/ameyarao98/fizzbuzz-server/server/internal/handler"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(
		&redis.Options{
			Addr: os.Getenv("REDIS_DSN"),
		})
	defer func(rdb *redis.Client) {
		if err := rdb.Close(); err != nil {
			log.Printf("Warning: %v", err)
		}
	}(rdb)

	handler := handler.NewHandler(rdb)
	router := internal.NewRouter(handler)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))
}
