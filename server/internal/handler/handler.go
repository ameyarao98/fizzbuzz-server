package handler

import (
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	rdb *redis.Client
}

func NewHandler(rdb *redis.Client) Handler {
	return Handler{
		rdb: rdb,
	}
}
