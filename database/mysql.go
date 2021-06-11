package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"singlishwords/config"
)

var db *sqlx.DB

func init() {
	//dsn := "csqsiew:123456@tcp(localhost:3306)/singlishwords?parseTime=True"
	dsn := config.MySqlDSN
	fmt.Println(dsn)
	db = sqlx.MustConnect("mysql", dsn)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
}

func GetMySQLDB() *sqlx.DB {
	return db
}
