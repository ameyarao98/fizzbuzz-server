package internal

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("fizz buzz")); err != nil {
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

	// Parse integers
	int1, err := strconv.Atoi(int1Str)
	if err != nil {
		http.Error(w, "int1 must be a valid integer", http.StatusBadRequest)
		return
	}
	int2, err := strconv.Atoi(int2Str)
	if err != nil {
		http.Error(w, "int2 must be a valid integer", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, "limit must be a valid integer", http.StatusBadRequest)
		return
	}

	result, err := GenerateFizzBuzz(int1, int2, limit, str1, str2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}
