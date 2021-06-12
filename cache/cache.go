package cache

import (
	"singlishwords/database"
)

var rdb, _ = database.GetRedisDB()

type cacheMissError struct{}

func (cacheMissError) Error() string {
	return "Cache miss error."
}

type notConnectedError struct{}

func (notConnectedError) Error() string {
	return "Not connected to redis server."
}
