package database

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"singlishwords/config"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Password, // no password set
		DB:       config.Redis.DB,       // use default DB
	})

	rdb.Set("questionIndex", 0, 0)
	rdb.Del("questionList")
}

func GetRedisDB() (*redis.Client, error) {
	if rdb != nil {
		return rdb, nil
	}
	return nil, fmt.Errorf("not connected to redis database")
}
