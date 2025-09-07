package mysql

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	connection *sql.DB
}

func NewDB() *DB {
	conn, err := sql.Open("mysql", "root:root@(localhost:3306)/game_db?parseTime=true")
	if err != nil {
		log.Panicf("cannot connect to database: %v", err)
	}

	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	return &DB{connection: conn}
}
