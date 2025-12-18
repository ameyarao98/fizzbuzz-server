package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ameyarao98/fizzbuzz-server/go/internal"
	"github.com/redis/go-redis/v9"
)

func main() {
	redisAddr := os.Getenv("REDIS_DSN")
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("Invalid REDIS_DB value: %v", err)
	}
	rdb := redis.NewClient(
		&redis.Options{
			Addr: redisAddr,
			DB:   redisDB,
		})
	defer func(rdb *redis.Client) {
		if err := rdb.Close(); err != nil {
			log.Print(err)
		}
	}(rdb)

	handler := internal.NewHandler(rdb)
	router := internal.NewRouter(handler)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))
}
