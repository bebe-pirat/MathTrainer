package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type StudentProgressRepository interface {
	CreateStudentProgressLevel(ctx context.Context, progress model.StudentProgress) error

	// старое
	StartLevel(ctx context.Context, studentId, LevelOrder int) error
	FinishLevel(ctx context.Context, e model.StudentProgress) (*model.StudentProgress, error)

	GetStudentProgress(ctx context.Context, studentId int) ([]model.StudentProgress, error)
	GetLevelProgress(ctx context.Context, studentId, LevelOrder int) ([]model.StudentProgress, error)

	// student stats
	GetCountComplitedLevels(ctx context.Context, studentId int) (int, error)
	GetTotalStars(ctx context.Context, studentId int) (int, error)
}

type StudentProgressRepositoryStruct struct {
	db *sql.DB
}

func NewStudentProgressRepositoryStruct(db *sql.DB) *StudentProgressRepositoryStruct {
	return &StudentProgressRepositoryStruct{
		db: db,
	}
}

func (r *StudentProgressRepositoryStruct) CreateStudentProgressLevel(ctx context.Context, progress model.StudentProgress) error {
	query := `
		INSERT INTO student_progress_level(user_id, section_id, level_order, count_stars, finished_at) 
		VALUES ($1, $2, $3, $4, $5);
	`

	res, err := r.db.ExecContext(ctx, query, progress.StudentId, progress.SectionId, progress.LevelOrder, progress.CountStarts, progress.FinishedAt)
	if err != nil {
		return err
	}

	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *StudentProgressRepositoryStruct) StartLevel(ctx context.Context, studentId, LevelOrder int) error {
	query := `
		INSERT INTO student_progress(student_id, level_id) 
		VALUES ($1, $2)
	`

	res, err := r.db.ExecContext(ctx, query, studentId, LevelOrder)
	if err != nil {
		return err
	}
	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *StudentProgressRepositoryStruct) FinishLevel(ctx context.Context, e model.StudentProgress) (*model.StudentProgress, error) {
	query := `
		UPDATE student_progress SET count_starts = $1, finished_at = $2
		WHERE id = $3
		RETURNING *
	`

	var progress model.StudentProgress
	err := r.db.QueryRowContext(ctx, query, e.CountStarts, e.FinishedAt, e.Id).Scan(&progress.Id, &progress.StudentId, &progress.LevelOrder, &progress.CountStarts, &progress.FinishedAt)
	if err != nil {
		return nil, err
	}

	return &progress, nil
}

func (r *StudentProgressRepositoryStruct) GetStudentProgress(ctx context.Context, studentId int) ([]model.StudentProgress, error) {
	query := `
		SELECT id, student_id, level_id, count_stars, finished_at
		WHERE student_id = $1	
	`

	rows, err := r.db.QueryContext(ctx, query, studentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	progresses := make([]model.StudentProgress, 0)
	for rows.Next() {
		var progress model.StudentProgress
		err := rows.Scan(
			&progress.Id,
			&progress.StudentId,
			&progress.LevelOrder,
			&progress.CountStarts,
			&progress.FinishedAt,
		)

		if err != nil {
			return nil, err
		}
		progresses = append(progresses, progress)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return progresses, nil
}

func (r *StudentProgressRepositoryStruct) GetLevelProgress(ctx context.Context, studentId, LevelOrder int) ([]model.StudentProgress, error) {
	query := `
		SELECT id, student_id, level_id, count_stars, finished_at
		WHERE student_id = $1 AND level_id = $2
	`

	rows, err := r.db.QueryContext(ctx, query, studentId, LevelOrder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	progresses := make([]model.StudentProgress, 0)
	for rows.Next() {
		var progress model.StudentProgress
		err := rows.Scan(
			&progress.Id,
			&progress.StudentId,
			&progress.LevelOrder,
			&progress.CountStarts,
			&progress.FinishedAt,
		)

		if err != nil {
			return nil, err
		}
		progresses = append(progresses, progress)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return progresses, nil
}

func (r *StudentProgressRepositoryStruct) GetCountComplitedLevels(ctx context.Context, studentId int) (int, error) {
	query := `
		SELECT COUNT(id) 
		FROM student_progress_level
		WHERE student_progress_level.student_Id = $1 and finished_at is not null
		GROUP BY student_id;
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, studentId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *StudentProgressRepositoryStruct) GetTotalStars(ctx context.Context, studentId int) (int, error) {
	query := `
		SELECT SUM(COALESCE(count_stars, 0))
		FROM student_progress_level
		WHERE student_progress_level.student_id = $1
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, studentId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
