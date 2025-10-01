package handler

import "net/http"

func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("fizz buzz")); err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
