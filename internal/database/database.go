package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Initialize db and create users table
func Init() {
	var err error
	DB, err = sql.Open("sqlite3", ".users.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXIST users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);
	`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}

}
