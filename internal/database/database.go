package database

import (
	"database/sql"
	"fmt"
	"time"
)

func OpenDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgresql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database")
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
