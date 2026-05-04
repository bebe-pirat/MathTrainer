package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type AchievementOfStudentRepository interface {
	GiveAchievementToStudent(ctx context.Context, e model.AchievementOfStudent) error
	GetAchievemntsByStudentId(ctx context.Context, studentId int) ([]model.AchievementOfStudent, error)
	GetAchievemntById(ctx context.Context, id int) (*model.Achievement, error)
}

type AchievementOfStudentRepositoryStruct struct {
	db *sql.DB
}

func NewAchievementOfStudentRepositoryStruct(db *sql.DB) *AchievementOfStudentRepositoryStruct {
	return &AchievementOfStudentRepositoryStruct{
		db: db,
	}
}

func (r *AchievementOfStudentRepositoryStruct) GiveAchievementToStudent(ctx context.Context, e model.AchievementOfStudent) error {
	query := `
		INSERT INTO achievements_of_students(student_id, achievement_id, got_at)
		VALUES ($1, $2, $3)
	`

	res, err := r.db.ExecContext(ctx, query, e.StudentId, e.AchievementId, e.GotAt)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *AchievementOfStudentRepositoryStruct) GetAchievemntsByStudentId(ctx context.Context, studentId int) ([]model.AchievementOfStudent, error) {
	query := `
		SELECT student_id, achievement_id, got_at
		FROM achievements_of_students 
		WHERE student_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, studentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	achs := make([]model.AchievementOfStudent, 0)
	for rows.Next() {
		var ach model.AchievementOfStudent
		err := rows.Scan(
			&ach.StudentId,
			&ach.AchievementId,
			&ach.GotAt,
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

func (r *AchievementOfStudentRepositoryStruct) GetAchievemntById(ctx context.Context, id int) (*model.Achievement, error) {
	query := `
		SELECT id, name, description
		FROM achievements
		WHERE id = $1;
	`

	var ach model.Achievement
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&ach.Id,
		&ach.Name,
		&ach.Description,
	)
	if err != nil {
		return nil, err
	}

	return &ach, nil
}
