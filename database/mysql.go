package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"singlishwords/config"
	"singlishwords/log"
	"sync"
)

var (
	db        *sqlx.DB
	mutexConn sync.RWMutex
	mutexIn   sync.RWMutex
)

func init() {
	dsn := config.MySqlDSN
	var err error
	db, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		db = nil
		log.Logger.Error("Cannot connect to Mysql database. Check your config files.")
		return
	}

	log.Logger.Info("Connect to mysql successful.")
	//db.SetMaxIdleConns(10)
	//db.SetMaxOpenConns(100)
}

func setMySqlConnection() {
	dsn := config.MySqlDSN
	var err error
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		db = nil
		log.Logger.Error("Cannot connect to Mysql database. Check your config files.")
		return
	}
}

func GetMySqlDB() (*sqlx.DB, error) {
	if db == nil {
		mutexConn.Lock()
		if db == nil {
			setMySqlConnection()
		}
		mutexConn.Unlock()
	}

	if db != nil {
		return db, nil
	}

	return nil, fmt.Errorf("not connected to mysql database")
}
