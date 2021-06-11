package cache

import (
	"singlishwords/database"
)

var rdb = database.GetRedisDB()

type cacheMissError struct{}

func (cacheMissError) Error() string {
	return "Cache miss error."
}
