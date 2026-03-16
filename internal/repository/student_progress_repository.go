package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type StudentProgressRepository interface {
	StartLevel(ctx context.Context, studentId, levelId int) (int, error)
	FinishLevel(ctx context.Context, studentID, levelID, stars int) (*model.StudentProgress, error)
	GetStudentProgress(ctx context.Context, studentId int) ([]model.StudentProgress, error)
	GetLevelProgress(ctx context.Context, studentId, levelId int) ([]model.StudentProgress, error)

	// student stats
	GetComplitedLevels(ctx context.Context, studentId int) (int, error)
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

func (r *StudentProgressRepositoryStruct) StartLevel(ctx context.Context, studentId, levelId int) (int, error) {
	query := `
		INSERT INTO student_progress(student_id, level_id) 
		VALUES ($1, $2)
		RETURNING id	
	`

	var id int
	err := r.db.QueryRowContext(ctx, query, studentId, levelId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *StudentProgressRepositoryStruct) FinishLevel(ctx context.Context, e model.StudentProgress) (*model.StudentProgress, error) {
	query := `
		UPDATE student_progress SET count_starts = $1, finished_at = $2
		WHERE id = $3
		RETURNING *
	`

	var progress model.StudentProgress
	err := r.db.QueryRowContext(ctx, query, e.CountStarts, e.FinishedAt, e.Id).Scan(&progress.Id, &progress.StudentId, &progress.LevelId, &progress.CountStarts, &progress.FinishedAt)
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
			&progress.LevelId,
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

func (r *StudentProgressRepositoryStruct) GetLevelProgress(ctx context.Context, studentId, levelId int) ([]model.StudentProgress, error) {
	query := `
		SELECT id, student_id, level_id, count_stars, finished_at
		WHERE student_id = $1 AND level_id = $2
	`

	rows, err := r.db.QueryContext(ctx, query, studentId, levelId)
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
			&progress.LevelId,
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

func (r *StudentProgressRepositoryStruct) GetComplitedLevels(ctx context.Context, studentId int) (int, error) {
	query := `
		SELECT COUNT(*) 
		FROM student_progress
		WHERE Student_progress.Student_Id = $1 and finished_at is not null
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
		SELECT SUM(COALESCE(Count_stars, 0))
		FROM student_progress
		WHERE Student_progress.Student_Id = $1
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, studentId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil	
}
