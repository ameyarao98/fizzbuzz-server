package api

import (
	"context"
	"encoding/json"
	"maps"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ameyarao98/fizzbuzz-server/server/internal/redis"
	"github.com/ameyarao98/fizzbuzz-server/server/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzHandler(t *testing.T) {
	rdbClient, cleanup := testutils.SetupRedis(t)
	defer cleanup()

	h := Handler{rdb: rdbClient}

	const (
		HealthEndpoint     = "/health"
		FizzBuzzEndpoint   = "/fizzbuzz"
		StatisticsEndpoint = "/statistics"
	)

	// Base valid query as map
	validQuery := map[string]string{
		"int1":  "2",
		"int2":  "3",
		"limit": "6",
		"str1":  "F",
		"str2":  "B",
	}

	t.Run("returns 400 for missing required params", func(t *testing.T) {
		for key := range validQuery {
			query := map[string]string{}
			maps.Copy(query, validQuery)
			delete(query, key)

			q := url.Values{}
			for k, v := range query {
				q.Set(k, v)
			}

			req := httptest.NewRequest(http.MethodGet, FizzBuzzEndpoint+"?"+q.Encode(), nil)
			w := httptest.NewRecorder()

			h.FizzBuzz(w, req)
			resp := w.Result()
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "missing param: %s", key)
		}
	})

	t.Run("returns 400 for invalid integers", func(t *testing.T) {
		intKeys := []string{"int1", "int2", "limit"}

		for _, key := range intKeys {
			// Copy map
			query := map[string]string{}
			maps.Copy(query, validQuery)
			query[key] = "abc" // invalid value

			q := url.Values{}
			for k, v := range query {
				q.Set(k, v)
			}

			req := httptest.NewRequest(http.MethodGet, FizzBuzzEndpoint+"?"+q.Encode(), nil)
			w := httptest.NewRecorder()

			h.FizzBuzz(w, req)
			resp := w.Result()
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "invalid int param: %s", key)
		}
	})

	t.Run("returns 400 for negative integers", func(t *testing.T) {
		intKeys := []string{"int1", "int2", "limit"}

		for _, key := range intKeys {
			query := map[string]string{}
			maps.Copy(query, validQuery)
			query[key] = "-5"

			q := url.Values{}
			for k, v := range query {
				q.Set(k, v)
			}

			req := httptest.NewRequest(http.MethodGet, FizzBuzzEndpoint+"?"+q.Encode(), nil)
			w := httptest.NewRecorder()

			h.FizzBuzz(w, req)
			resp := w.Result()
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "negative int param: %s", key)
		}
	})

	t.Run("returns 200 and correct fizzbuzz result", func(t *testing.T) {
		q := url.Values{}
		for k, v := range validQuery {
			q.Set(k, v)
		}

		req := httptest.NewRequest(http.MethodGet, FizzBuzzEndpoint+"?"+q.Encode(), nil)
		w := httptest.NewRecorder()

		h.FizzBuzz(w, req)
		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		var body []string
		err := json.NewDecoder(resp.Body).Decode(&body)
		assert.NoError(t, err)
		assert.Equal(t, []string{"1", "F", "B", "F", "5", "FB"}, body)

		key := redis.GenerateRedisKey(2, 3, 6, "F", "B")
		count, err := rdbClient.ZScore(context.Background(), redis.RequestsKey, key).Result()
		assert.NoError(t, err)
		assert.Equal(t, 1.0, count)
	})
}
