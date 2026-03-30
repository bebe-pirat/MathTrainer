package database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func OpenDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(10 * time.Second)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)

	return db, nil
}

func CloseDB(db *sql.DB) error {
	if db == nil {
		return nil
	}

	return db.Close()
}
