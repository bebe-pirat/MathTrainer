package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type LevelRepository interface {
	GetAllLevels(ctx context.Context) ([]model.Level, error)
	GetById(ctx context.Context, id int) (model.Level, error)
	GetTestLevel(ctx context.Context) (model.Level, error)
}

type LevelRepositoryStruct struct {
	db *sql.DB
}

func NewLevelRepositoryStruct(db *sql.DB) *LevelRepositoryStruct {
	return &LevelRepositoryStruct{
		db: db,
	}
}

func (r *LevelRepositoryStruct) GetAllLevels(ctx context.Context) ([]model.Level, error) {
	query := `
		SELECT id, name, test_level, difficulty
		FROM levels	
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	levels := make([]model.Level, 0)
	for rows.Next() {
		var level model.Level
		err := rows.Scan(
			&level.Id,
			&level.Name,
			&level.TestLevel,
			&level.Difficulty,
		)

		if err != nil {
			return nil, err
		}

		levels = append(levels, level)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return levels, err
}

func (r *LevelRepositoryStruct) GetById(ctx context.Context, id int) (*model.Level, error) {
	query := `
		SELECT id, name, test_level, difficulty
		FROM levels	
		WHERE id = $1
	`

	var level model.Level
	err := r.db.QueryRowContext(ctx, query, id).Scan(&level.Id, &level.Name, &level.TestLevel, &level.Difficulty)
	if err != nil {
		return nil, err
	}

	return &level, nil
}

func (r *LevelRepositoryStruct) GetTestLevel(ctx context.Context) (*model.Level, error) {
	query := `
		SELECT id, name, test_level, difficulty
		FROM levels	
		WHERE test_level = true
	`

	var level model.Level
	err := r.db.QueryRowContext(ctx, query).Scan(&level.Id, &level.Name, &level.TestLevel, &level.Difficulty)
	if err != nil {
		return nil, err
	}

	return &level, nil
}
