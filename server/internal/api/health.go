package api

import "net/http"

const healthResponse = "fizz buzz"

func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(healthResponse)); err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
