package database

import (
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
}

func GetRedisDB() *redis.Client {
	return rdb
}
