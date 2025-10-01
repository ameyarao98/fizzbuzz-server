package internal

import (
	"net/http"

	"github.com/ameyarao98/fizzbuzz-server/server/internal/handler"
)

func NewRouter(handler handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /fizzbuzz", handler.FizzBuzz)
	mux.HandleFunc("GET /statistics", handler.Statistics)

	return mux
}
