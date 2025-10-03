package testutils

import (
	"context"
	"os"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func SetupRedis(t *testing.T) (*redis.Client, func()) {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_DSN"),
		DB:   1,
	})

	cleanup := func() {
		err := rdb.FlushDB(context.Background()).Err()
		require.NoError(t, err, "Failed to flush Redis DB")
		err = rdb.Close()
		require.NoError(t, err, "Failed to close Redis connection")
	}

	return rdb, cleanup
}
