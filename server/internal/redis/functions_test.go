package redis

import (
	"context"
	"testing"

	"github.com/ameyarao98/fizzbuzz-server/server/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestIncreaseCounter(t *testing.T) {
	rdb, cleanup := testutils.SetupRedis(t)
	defer cleanup()

	t.Run("should increment counter for new key", func(t *testing.T) {
		key1 := "test:key:1"

		err := IncreaseCounter(context.Background(), rdb, key1)
		assert.NoError(t, err)

		count, err := rdb.ZScore(context.Background(), RequestsKey, key1).Result()
		assert.NoError(t, err)
		assert.Equal(t, 1.0, count)
	})

	t.Run("should increment counter for existing key", func(t *testing.T) {
		key1 := "test:key:1"

		err := IncreaseCounter(context.Background(), rdb, key1)
		assert.NoError(t, err)

		count, err := rdb.ZScore(context.Background(), RequestsKey, key1).Result()
		assert.NoError(t, err)
		assert.Equal(t, 2.0, count)
	})

	t.Run("should handle multiple keys independently", func(t *testing.T) {
		key2 := "test:key:2"
		key3 := "test:key:3"

		err := IncreaseCounter(context.Background(), rdb, key2)
		assert.NoError(t, err)

		err = IncreaseCounter(context.Background(), rdb, key3)
		assert.NoError(t, err)

		err = IncreaseCounter(context.Background(), rdb, key2)
		assert.NoError(t, err)

		count, err := rdb.ZScore(context.Background(), RequestsKey, key2).Result()
		assert.NoError(t, err)
		assert.Equal(t, 2.0, count)

		count, err = rdb.ZScore(context.Background(), RequestsKey, key3).Result()
		assert.NoError(t, err)
		assert.Equal(t, 1.0, count)
	})
}
func TestGetHighestCount(t *testing.T) {
	rdb, cleanup := testutils.SetupRedis(t)
	defer cleanup()

	t.Run("should return empty result when no keys exist", func(t *testing.T) {
		key, count, err := GetHighestCount(context.Background(), rdb)
		assert.NoError(t, err)
		assert.Equal(t, "", key)
		assert.Equal(t, 0.0, count)
	})

	t.Run("should return highest count when one key exists", func(t *testing.T) {
		key1 := "test:highest:1"
		err := rdb.ZIncrBy(context.Background(), RequestsKey, 1, key1).Err()
		assert.NoError(t, err)

		key, count, err := GetHighestCount(context.Background(), rdb)
		assert.NoError(t, err)
		assert.Equal(t, key1, key)
		assert.Equal(t, 1.0, count)
	})

	t.Run("should return key with highest count among multiple", func(t *testing.T) {
		key2 := "test:highest:2"
		key3 := "test:highest:3"

		err := rdb.ZIncrBy(context.Background(), RequestsKey, 2, key2).Err()
		assert.NoError(t, err)
		err = rdb.ZIncrBy(context.Background(), RequestsKey, 1, key3).Err()
		assert.NoError(t, err)

		key, count, err := GetHighestCount(context.Background(), rdb)
		assert.NoError(t, err)
		assert.Equal(t, key2, key)
		assert.Equal(t, 2.0, count)
	})

}
