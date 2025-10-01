package internal

import "net/http"

func NewRouter(handler Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /fizzbuzz", handler.FizzBuzz)

	return mux
}
