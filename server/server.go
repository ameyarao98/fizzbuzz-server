package main

import (
	"net/http"
	"os"

	"github.com/ameyarao98/fizzbuzz-server/server/internal"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	handler := internal.NewHandler()
	router := internal.NewRouter(handler)

	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
