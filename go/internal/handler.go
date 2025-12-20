package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Handler struct {
	rdb *redis.Client
}

func NewHandler(rdb *redis.Client) Handler {
	return Handler{
		rdb: rdb,
	}
}

const healthResponse = "fizz buzz: go"

func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(healthResponse)); err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

func (h Handler) FizzBuzz(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	int1Str := query.Get("int1")
	int2Str := query.Get("int2")
	limitStr := query.Get("limit")
	str1 := query.Get("str1")
	str2 := query.Get("str2")

	// Validate required parameters
	if int1Str == "" {
		http.Error(w, "Missing required parameter: int1", http.StatusBadRequest)
		return
	}
	if int2Str == "" {
		http.Error(w, "Missing required parameter: int2", http.StatusBadRequest)
		return
	}
	if limitStr == "" {
		http.Error(w, "Missing required parameter: limit", http.StatusBadRequest)
		return
	}
	if str1 == "" {
		http.Error(w, "Missing required parameter: str1", http.StatusBadRequest)
		return
	}
	if str2 == "" {
		http.Error(w, "Missing required parameter: str2", http.StatusBadRequest)
		return
	}

	int1Val, err := strconv.Atoi(int1Str)
	if err != nil {
		http.Error(w, "int1 must be a valid integer", http.StatusBadRequest)
		return
	}
	if int1Val <= 0 {
		http.Error(w, "int1 must be a positive integer", http.StatusBadRequest)
		return
	}
	int1 := uint(int1Val)

	int2Val, err := strconv.Atoi(int2Str)
	if err != nil {
		http.Error(w, "int2 must be a valid integer", http.StatusBadRequest)
		return
	}
	if int2Val <= 0 {
		http.Error(w, "int2 must be a positive integer", http.StatusBadRequest)
		return
	}
	int2 := uint(int2Val)

	limitVal, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, "limit must be a valid integer", http.StatusBadRequest)
		return
	}
	if limitVal < 0 {
		http.Error(w, "limit must be a positive integer", http.StatusBadRequest)
		return
	}
	limit := uint(limitVal)

	result := GenerateFizzBuzz(int1, int2, limit, str1, str2)

	key := GenerateRedisKey(int1, int2, limit, str1, str2)
	err = IncreaseCounter(r.Context(), h.rdb, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

}
func (h Handler) Statistics(w http.ResponseWriter, r *http.Request) {
	key, count, err := GetHighestCount(r.Context(), h.rdb)
	if err != nil {
		http.Error(w, "Error getting statistics", http.StatusInternalServerError)
		return
	}
	data := map[string]float64{key: count}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonData); err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
