package internal

import "net/http"

const (
	HealthEndpoint     = "/health"
	FizzBuzzEndpoint   = "/fizzbuzz"
	StatisticsEndpoint = "/statistics"
)

func NewRouter(handler Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(HealthEndpoint, handler.Health)
	mux.HandleFunc(FizzBuzzEndpoint, handler.FizzBuzz)
	mux.HandleFunc(StatisticsEndpoint, handler.Statistics)

	return mux
}
