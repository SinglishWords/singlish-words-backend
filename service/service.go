package service

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db  *sqlx.DB
	rdb *redis.Client
)

func init() {
	dsn := "csqsiew:123456@tcp(localhost:3306)/singlishwords"
	db = sqlx.MustConnect("mysql", dsn)

	// db.SetMaxIdleConns(10)
	// db.SetMaxOpenConns(100)

	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}

func GetMySQLDB() *sqlx.DB {
	return db
}

func GetRedisDB() *redis.Client {
	return rdb
}
