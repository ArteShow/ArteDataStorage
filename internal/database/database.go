package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() error {
	var err error
	DB, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		return err
	}
	return DB.Ping()
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
