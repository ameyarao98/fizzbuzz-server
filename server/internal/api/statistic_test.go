package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ameyarao98/fizzbuzz-server/server/internal/redis"
	"github.com/ameyarao98/fizzbuzz-server/server/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestStatisticsHandler(t *testing.T) {
	rdb, cleanup := testutils.SetupRedis(t)
	defer cleanup()

	h := NewHandler(rdb)

	t.Run("empty Redis returns 0", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, StatisticsEndpoint, nil)
		w := httptest.NewRecorder()

		h.Statistics(w, req)
		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		var data map[string]float64
		err := json.NewDecoder(resp.Body).Decode(&data)
		assert.NoError(t, err)
		assert.Equal(t, map[string]float64{"": 0}, data)
	})

	t.Run("single key in Redis", func(t *testing.T) {
		key := "key:1"
		err := redis.IncreaseCounter(context.Background(), rdb, key)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, StatisticsEndpoint, nil)
		w := httptest.NewRecorder()

		h.Statistics(w, req)
		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var data map[string]float64
		err = json.NewDecoder(resp.Body).Decode(&data)
		assert.NoError(t, err)
		assert.Equal(t, map[string]float64{key: 1}, data)
	})

	t.Run("multiple keys, returns highest count", func(t *testing.T) {
		key1 := "key:1"
		key2 := "key:2"

		// Increase counters
		err := redis.IncreaseCounter(context.Background(), rdb, key1) // 2 total
		assert.NoError(t, err)
		err = redis.IncreaseCounter(context.Background(), rdb, key2) // 1 total
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, StatisticsEndpoint, nil)
		w := httptest.NewRecorder()

		h.Statistics(w, req)
		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var data map[string]float64
		err = json.NewDecoder(resp.Body).Decode(&data)
		assert.NoError(t, err)
		assert.Equal(t, map[string]float64{key1: 2}, data)
	})
}
