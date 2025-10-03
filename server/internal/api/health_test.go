package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	h := NewHandler(nil) // redis is not used here

	req := httptest.NewRequest(http.MethodGet, HealthEndpoint, nil)
	w := httptest.NewRecorder()

	h.Health(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body := w.Body.String()
	assert.Equal(t, healthResponse, body)
}
