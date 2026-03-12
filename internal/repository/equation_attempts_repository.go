package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type EquationAttemptsRepository interface {
	CreateAttempt(ctx context.Context, e model.EquationAttempts) (int, error)
	GetStudentAttempts(ctx context.Context, studentId int) ([]model.EquationAttempts, error)
	GetErrorStats(ctx context.Context, studentId int) (int, error)
	GetAllStats(ctx context.Context, studentId int) (int, error)
	// GetAttemptsByLevel(ctx context.Context, studentId int, levelId int) ([]model.EquationAttempts, error)
}

type EquationAttemptsRepositoryStruct struct {
	db *sql.DB
}

func NewEquationAttemptsRepositoryStruct(db *sql.DB) *EquationAttemptsRepositoryStruct {
	return &EquationAttemptsRepositoryStruct{
		db: db,
	}
}

func (r *EquationAttemptsRepositoryStruct) CreateAttempt(ctx context.Context, e model.EquationAttempts) (int, error) {
	query := `
		INSERT INTO equation_attempts(student_id, equation_id, given_answer, correct, attempted_at) 
		VALUES($1, $2, $3, $4, $5, $6) 
		RETURNING id
	`

	var id int
	err := r.db.QueryRowContext(ctx, query, e.StudentId, e.EquationId, e.GivenAnswer, e.Correct, e.AttemptedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *EquationAttemptsRepositoryStruct) GetStudentAttempts(ctx context.Context, studentId int) ([]model.EquationAttempts, error) {
	query := `
		SELECT id, student_id, equation_id, given_answer, correct, attempted_at
		FROM equation_attempts
		WHERE student_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, studentId)
	if err != nil {
		return nil, err
	}

	attempts := make([]model.EquationAttempts, 0)
	for rows.Next() {
		var att model.EquationAttempts
		err := rows.Scan(
			&att.Id,
			&att.StudentId,
			&att.EquationId,
			&att.GivenAnswer,
			&att.Correct,
			&att.AttemptedAt)
		if err != nil {
			return nil, err
		}

		attempts = append(attempts, att)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attempts, nil
}

func (r *EquationAttemptsRepositoryStruct) GetErrorStats(ctx context.Context, studentId int) (int, error) {
	query := `
		select count(*) from equation_attempts where correct = false and student_id = $1
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, studentId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *EquationAttemptsRepositoryStruct) GetAllStats(ctx context.Context, studentId int) (int, error) {
	query := `
		select count(*) from equation_attempts where student_id = $1
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, studentId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// func (r *EquationAttemptsRepositoryStruct) GetAttemptsByLevel(ctx context.Context, studentId int, levelId int) ([]model.EquationAttempts, error) {
// 	query := `
// 		SELECT id, student_id, equation_id, given_answer, correct, attempted_at
// 		FROM equation_attempts
// 		WHERE student_id = $1 AND level
// 	`

// 	rows, err := r.db.QueryContext(ctx, query, studentId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	attempts := make([]model.EquationAttempts, 0)
// 	for rows.Next() {
// 		var att model.EquationAttempts
// 		err := rows.Scan(
// 			&att.Id,
// 			&att.StudentId,
// 			&att.EquationId,
// 			&att.GivenAnswer,
// 			&att.Correct,
// 			&att.AttemptedAt)
// 		if err != nil {
// 			return nil, err
// 		}

// 		attempts = append(attempts, att)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return attempts, nil
// }
