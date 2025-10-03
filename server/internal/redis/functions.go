package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func IncreaseCounter(ctx context.Context, rdb *redis.Client, key string) error {
	if err := rdb.ZIncrBy(ctx, RequestsKey, 1, key).Err(); err != nil {
		return fmt.Errorf("failed to increment counter for key %s: %w", key, err)
	}
	return nil
}

func GetHighestCount(ctx context.Context, rdb *redis.Client) (string, float64, error) {
	result, err := rdb.ZRevRangeWithScores(ctx, RequestsKey, 0, 0).Result()
	if err != nil {
		return "", 0, fmt.Errorf("failed to get highest count: %w", err)
	}
	if len(result) == 0 {
		return "", 0, nil
	}

	return result[0].Member.(string), result[0].Score, nil
}
