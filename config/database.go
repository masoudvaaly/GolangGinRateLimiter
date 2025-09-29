package config

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func ConnectDatabase() error {
	//db, err := sql.Open("sqlite3", "D:\\java\\go\\database.db")
	db, err := sql.Open("sqlite", "file:database.db?cache=shared&mode=rwc")
	if err != nil {
		return err
	}

	DB = db
	return nil
}
