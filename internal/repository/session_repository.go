package repository

import (
	"context"
	"database/sql"
	"time"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, userId int, expiresAt time.Time) (int, error)
	DeleteSession(ctx context.Context, id int) error
}

type SessionRepositoryStruct struct {
	db *sql.DB
}

func NewSessionRepositoryStruct(db *sql.DB) *SessionRepositoryStruct {
	return &SessionRepositoryStruct{
		db: db,
	}
}

func (r *SessionRepositoryStruct) CreateSession(ctx context.Context, userId int, expiresAt time.Time) (int, error) {
	query := `
		INSERT INTO sessions(user_id, expires_at)
		VALUES ($1, $2)
		RETURNING id;
	`

	var id int
	err := r.db.QueryRowContext(ctx, query, userId, expiresAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SessionRepositoryStruct) DeleteSession(ctx context.Context, id int) error {
	query := `
		DELETE FROM sessions 
		WHERE id = $1	
	`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
