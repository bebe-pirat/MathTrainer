package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type AchievementRepository interface {
	CreateAchievement(ctx context.Context, e model.Achievement) (int, error)
	UpdateAchievement(ctx context.Context, e model.Achievement) (*model.Achievement, error)
	DeleteAchievement(ctx context.Context, id int) error
	GetAllAchievement(ctx context.Context, e model.Achievement) ([]model.Achievement, error)
	GetAchievementById(ctx context.Context, id int) (*model.Achievement, error)
}

type AchievementRepositoryStruct struct {
	db *sql.DB
}

func NewAchievementRepositoryStruct(db *sql.DB) *AchievementRepositoryStruct {
	return &AchievementRepositoryStruct{
		db: db,
	}
}

func (r *AchievementRepositoryStruct) CreateAchievement(ctx context.Context, e model.Achievement) (int, error) {
	query := `
		INSERT INTO achievements(name, description)
		VALUES ($1, $2)
		RETURNING id
	`

	var id int
	err := r.db.QueryRowContext(ctx, query, e.Name, e.Description).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AchievementRepositoryStruct) UpdateAchievement(ctx context.Context, e model.Achievement) (*model.Achievement, error) {
	query := `
		UPDATE achievements 
		SET name = $1, description = $2
		WHERE id = $3
	`

	var ach model.Achievement
	err := r.db.QueryRowContext(ctx, query, e.Name, e.Description, e.Id).Scan(&ach.Id, &ach.Name, &ach.Description)
	if err != nil {
		return nil, err
	}

	return &ach, nil
}

func (r *AchievementRepositoryStruct) DeleteAchievement(ctx context.Context, id int) error {
	query := `
		DELETE FROM achievements
		WHERE id = $1
	`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *AchievementRepositoryStruct) GetAllAchievement(ctx context.Context, e model.Achievement) ([]model.Achievement, error) {
	query := `
		SELECT name, description
		FROM achievements
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	achs := make([]model.Achievement, 0)
	for rows.Next() {
		var ach model.Achievement
		err := rows.Scan(
			&ach.Id,
			&ach.Name,
			&ach.Description,
		)

		if err != nil {
			return nil, err
		}

		achs = append(achs, ach)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return achs, nil
}

func (r *AchievementRepositoryStruct) GetAchievementById(ctx context.Context, id int) (*model.Achievement, error) {
	query := `
		SELECT name, description
		FROM achievements
		WHERE id = $1
	`

	var ach model.Achievement
	err := r.db.QueryRowContext(ctx, query, id).Scan(&ach.Id, &ach.Name, &ach.Description)
	if err != nil {
		return nil, err
	}

	return &ach, nil
}
