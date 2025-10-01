package rdb

import "fmt"

const requestsKey = "requests"

func GenerateRedisKey(int1, int2, limit uint, str1, str2 string) string {
	return fmt.Sprintf("%d:%d:%d:%s:%s", int1, int2, limit, str1, str2)
}
