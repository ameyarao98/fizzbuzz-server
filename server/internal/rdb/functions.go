package rdb

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func IncreaseCounter(rdb *redis.Client, key string) error {
	if err := rdb.ZIncrBy(context.Background(), requestsKey, 1, key).Err(); err != nil {
		return fmt.Errorf("failed to increment counter for key %s: %w", key, err)
	}
	return nil
}

func GetHighestCount(rdb *redis.Client) (string, float64, error) {
	result, err := rdb.ZRevRangeWithScores(context.Background(), requestsKey, 0, 0).Result()
	if err == redis.Nil {
		return "", 0, nil
	} else if err != nil {
		return "", 0, fmt.Errorf("failed to get highest count: %w", err)
	}

	return result[0].Member.(string), result[0].Score, nil
}
