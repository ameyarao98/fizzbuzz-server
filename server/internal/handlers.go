package internal

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct{}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("healthy")); err != nil {
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
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Missing required parameter: int1"}); err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		}
		return
	}
	if int2Str == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Missing required parameter: int2"}); err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		}
		return
	}
	if limitStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Missing required parameter: limit"}); err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		}
		return
	}
	if str1 == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Missing required parameter: str1"}); err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		}
		return
	}
	if str2 == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Missing required parameter: str2"}); err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		}
		return
	}

	// Parse integers
	int1, err := strconv.Atoi(int1Str)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(ErrorResponse{Error: "int1 must be a valid integer"}); err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		}
		return
	}
	int2, err := strconv.Atoi(int2Str)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(ErrorResponse{Error: "int2 must be a valid integer"}); err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		}
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(ErrorResponse{Error: "limit must be a valid integer"}); err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		}
		return
	}

	result, err := GenerateFizzBuzz(int1, int2, limit, str1, str2)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()}); err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}
