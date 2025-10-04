package testutils

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func SetupRedis(t *testing.T) (*redis.Client, func()) {
	redisAddr := os.Getenv("REDIS_DSN")
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	require.NoError(t, err, "Invalid REDIS_DB value")
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   redisDB,
	})

	cleanup := func() {
		err := rdb.FlushDB(context.Background()).Err()
		require.NoError(t, err, "Failed to flush Redis DB")
		err = rdb.Close()
		require.NoError(t, err, "Failed to close Redis connection")
	}

	return rdb, cleanup
}
