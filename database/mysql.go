package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"singlishwords/config"
)

var db *sqlx.DB

func init() {
	dsn := config.MySqlDSN
	var err error
	db, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		db = nil
		log.Default().Println("Cannot connect to Mysql database. Check your config files.")
		return
	}

	//db.SetMaxIdleConns(10)
	//db.SetMaxOpenConns(100)
}

func GetMySQLDB() (*sqlx.DB, error) {
	if db != nil {
		return db, nil
	}
	return nil, fmt.Errorf("not connected to mysql database")
}
