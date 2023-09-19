package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() error {
	db, error := sql.Open("sqlite", "database.db")
	if error != nil {
		return error
	}
	DB = db
	return nil
}
