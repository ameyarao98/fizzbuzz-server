package api

import (
	"encoding/json"
	"net/http"

	"github.com/ameyarao98/fizzbuzz-server/server/internal/redis"
)

func (h Handler) Statistics(w http.ResponseWriter, r *http.Request) {
	key, count, err := redis.GetHighestCount(r.Context(), h.rdb)
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
