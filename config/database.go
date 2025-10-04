package config

import (
	"database/sql"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

var GormDB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "root",
		DBName:   "stripe",
	}
	return &dbConfig
}

func ConnectDatabase() error {
	//db, err := sql.Open("sqlite3", "D:\\java\\go\\database.db")
	db, err := sql.Open("sqlite", "file:database.db?cache=shared&mode=rwc")
	if err != nil {
		return err
	}

	DB = db
	return nil
}
