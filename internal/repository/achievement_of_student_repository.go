package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type AchievementOfStudentRepository interface {
	CreateAchieveOfStud(ctx context.Context, e model.AchievementOfStudent) error
	GetAllAchievementOfStudents(ctx context.Context) ([]model.AchievementOfStudent, error)
	GetAchievementOfStudentsByStudentId(ctx context.Context, studentId int) ([]model.AchievementOfStudent, error)
	GetAchievementOfStudentsByAchievementId(ctx context.Context, achievementId int) ([]model.AchievementOfStudent, error)
}

type AchievementOfStudentRepositoryStruct struct {
	db *sql.DB
}

func NewAchievementOfStudentRepositoryStruct(db *sql.DB) *AchievementOfStudentRepositoryStruct {
	return &AchievementOfStudentRepositoryStruct{
		db: db,
	}
}

func (r *AchievementOfStudentRepositoryStruct) CreateAchieveOfStud(ctx context.Context, e model.AchievementOfStudent) error {
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

func (r *AchievementOfStudentRepositoryStruct) GetAllAchievementOfStudents(ctx context.Context) ([]model.AchievementOfStudent, error) {
	query := `
		SELECT student_id, achievement_id, got_at
		FROM achievements_of_students 
	`

	rows, err := r.db.QueryContext(ctx, query)
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

func (r *AchievementOfStudentRepositoryStruct) GetAchievementOfStudentsByStudentId(ctx context.Context, studentId int) ([]model.AchievementOfStudent, error) {
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

func (r *AchievementOfStudentRepositoryStruct) GetAchievementOfStudentsByAchievementId(ctx context.Context, achievementId int) ([]model.AchievementOfStudent, error) {
	query := `
		SELECT student_id, achievement_id, got_at
		FROM achievements_of_students 
		WHERE achievement_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, achievementId)
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
