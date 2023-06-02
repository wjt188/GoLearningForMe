package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Init() *sql.DB {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}
