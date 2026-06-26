package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	DB *sql.DB
}

func NewSQLiteStorage() *SQLiteStorage {
	return &SQLiteStorage{
		DB: initDb(),
	}
}

func initDb() *sql.DB {
	db, err := sql.Open("sqlite3", "storage/data.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`PRAGMA foreign_keys = ON`)
	if err != nil {
		log.Fatal(err)
	}

	for _, script := range Tables {
		_, err = db.Exec(script)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, script := range Indices {
		_, err = db.Exec(script)
		if err != nil {
			log.Fatal(err)
		}
	}

	return db
}
