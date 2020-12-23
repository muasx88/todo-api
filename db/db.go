package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {

	dsn := os.Getenv("MYSQL_DSN")
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		panic("Connection Error")
	}

	if err = db.Ping(); err != nil {
		panic("Dsn invalid")
	}

}

func CreateCon() *sql.DB {
	return db
}
