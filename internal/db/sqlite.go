package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb() *sql.DB {
	db, err := sql.Open("sqlite3", "app.db")
	if err != nil {
		log.Fatal(err)
	}
	CreateTable := `
        CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL,
		done BOOLEAN NOT NULL CHECK (done IN (0,1))
	);
`
	if _, err := db.Exec(CreateTable); err != nil {
		db.Close()
		log.Fatal(err)
	}

	return db
}
