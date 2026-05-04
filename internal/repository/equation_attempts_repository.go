package repository

import (
	"MathTrainer/internal"
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type EquationAttemptsRepository interface {
	GetStudentSectionStats(ctx context.Context, studentId int, sectionId int) (map[int]float32, error)

	CreateAttempt(ctx context.Context, e model.Attempt) error
	GetStudentAttempts(ctx context.Context, studentId int, equation_type_id int) ([]model.AttemptForTeacher, error)

    // school stats
	GetTotalAttemptsBySchoolId(ctx context.Context, schoolId int) (int, error)
	GetWrongAnswersBySchoolId(ctx context.Context, schoolId int) (int, error)
	GetClassesAccuracyBySchoolId(ctx context.Context, schoolId int) ([]model.ClassShortStats, error)
	GetEquationTypeAccuracyBySchoolId(ctx context.Context, schoolId int) ([]model.EquationTypeStats, error)

	// class stats
	GetStudentsShortStatsByClassId(ctx context.Context, classId int) ([]model.StudentShortStats, error)
	GetWrongAnswersByClassId(ctx context.Context, classId int) (int, error)
	GetTotalAttemptsByClassId(ctx context.Context, classId int) (int, error)
	GetEquationTypeAccuracyByClassId(ctx context.Context, classId int) ([]model.EquationTypeStats, error)

	// student stats
	GetCountErrorAttempts(ctx context.Context, studentId int) (int, error)
	GetTotalCountAttempts(ctx context.Context, studentId int) (int, error)
	GetExtendedEquationTypeStats(ctx context.Context, studentId int) ([]model.ExtendedEquationTypeStats, error)
}

type EquationAttemptsRepositoryStruct struct {
	db *sql.DB
}

func NewEquationAttemptsRepositoryStruct(db *sql.DB) *EquationAttemptsRepositoryStruct {
	return &EquationAttemptsRepositoryStruct{
		db: db,
	}
}

func (r *EquationAttemptsRepositoryStruct) GetStudentSectionStats(ctx context.Context, studentId int, sectionId int) (map[int]float32, error) {
	query := `
		SELECT equation_type_id, success_rate
		FROM student_section_stats
		WHERE student_id = $1 AND section_id = $2;
	`

	rows, err := r.db.QueryContext(ctx, query, studentId, sectionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	shortStats := make(map[int]float32)
	for rows.Next() {
		var stat model.ShortEquationTypeStats

		err := rows.Scan(
			&stat.TypeId,
			&stat.Accuracy,
		)
		if err != nil {
			return nil, err
		}

		shortStats[stat.TypeId] = stat.Accuracy
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return shortStats, nil
}

func (r *EquationAttemptsRepositoryStruct) CreateAttempt(ctx context.Context, e model.Attempt) error {
	query := `
		INSERT INTO attempts(student_id, equation_type_id, equation_text, correct_answer, student_answer, answered_at)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var id int
	err := r.db.QueryRowContext(ctx, query, e.StudentId, e.EquationTypeId, e.EquationText, e.CorrectAnswer, e.GivenAnswer, e.AnsweredAt).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (r *EquationAttemptsRepositoryStruct) GetStudentAttempts(ctx context.Context, studentId int, equation_type_id int) ([]model.AttemptForTeacher, error) {
	query := `
		SELECT attempts.equation_text, attempts.equation_type_id, equation_types.name, student_answer, correct_answer, answered_at
		FROM attempts 
		JOIN equation_types ON attempts.equation_type_id = equation_types.id
		WHERE student_id = $1
	`

	parameters := make([]any, 0)
	parameters = append(parameters, studentId)

	if equation_type_id > 0 {
		parameters = append(parameters, equation_type_id)
		query += " AND equation_type_id = $2;"
	}

	rows, err := r.db.QueryContext(ctx, query, parameters...)
	if err != nil {
		return nil, err
	}

	attempts := make([]model.AttemptForTeacher, 0)
	for rows.Next() {
		var att model.AttemptForTeacher
		err := rows.Scan(
			&att.EquationText,
			&att.EquationTypeId,
			&att.EquationTypeName,
			&att.GivenAnswer,
			&att.CorrectAnswer,
			&att.AnsweredAt)
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

func (r *EquationAttemptsRepositoryStruct) GetCountErrorAttempts(ctx context.Context, studentId int) (int, error) {
	query := `
		SELECT count(id) 
		FROM attempts 
		WHERE correct_answer <> student_answer and student_id = $1
		GROUP BY student_id;
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, studentId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *EquationAttemptsRepositoryStruct) GetTotalCountAttempts(ctx context.Context, studentId int) (int, error) {
	query := `
		SELECT count(id) 
		FROM attempts 
		WHERE student_id = $1
		GROUP BY student_id;
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, studentId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *EquationAttemptsRepositoryStruct) GetTotalAttemptsBySchoolId(ctx context.Context, schoolId int) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM attempts
		JOIN users on users.id = attempts.student_id
		JOIN classes on classes.id = users.class_id
		WHERE classes.school_id = $1;
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, schoolId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *EquationAttemptsRepositoryStruct) GetWrongAnswersBySchoolId(ctx context.Context, schoolId int) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM attempts
		JOIN users on users.id = attempts.student_id
		JOIN classes on classes.id = users.class_id
		WHERE classes.school_id = $1 AND attempts.correct_answer = attempts.student_answer;
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, schoolId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *EquationAttemptsRepositoryStruct) GetClassesAccuracyBySchoolId(ctx context.Context, schoolId int) ([]model.ClassShortStats, error) {
	query := `
		SELECT classes.name, COALESCE(SUM(CASE WHEN attempts.correct_answer = attempts.student_answer THEN 1 ELSE 0 END) * 100.0 / NULLIF(COUNT(attempts.id), 0), 0)
		FROM attempts
		JOIN users on users.id = attempts.student_id
		JOIN classes on classes.id = users.class_id
		WHERE classes.school_id = $1
		GROUP BY classes.name;
	`

	rows, err := r.db.QueryContext(ctx, query, schoolId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := make([]model.ClassShortStats, 0)
	for rows.Next() {
		var stat model.ClassShortStats

		err := rows.Scan(
			&stat.Name,
			&stat.Accuracy,
		)
		if err != nil {
			return nil, err
		}

		stats = append(stats, stat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *EquationAttemptsRepositoryStruct) GetEquationTypeAccuracyBySchoolId(ctx context.Context, schoolId int) ([]model.EquationTypeStats, error) {
	query := `
		SELECT equation_types.name, COALESCE(SUM(CASE WHEN attempts.correct_answer = attempts.student_answer THEN 1 ELSE 0 END) * 100.0 / NULLIF(COUNT(attempts.id), 0), 0)
		FROM attempts
		JOIN equation_types ON attempts.equation_type_id = equation_types.id
		JOIN users ON attempts.student_id = users.id
		JOIN classes ON users.class_id = classes.id
		WHERE classes.school_id = 1
		GROUP BY equation_types.name;
	`

	rows, err := r.db.QueryContext(ctx, query, schoolId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := make([]model.EquationTypeStats, 0)
	for rows.Next() {
		var stat model.EquationTypeStats

		err := rows.Scan(
			&stat.Type,
			&stat.Accuracy,
		)
		if err != nil {
			return nil, err
		}

		stats = append(stats, stat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *EquationAttemptsRepositoryStruct) GetStudentsShortStatsByClassId(ctx context.Context, classId int) ([]model.StudentShortStats, error) {
	query := `
		WITH levels_complited AS (
			SELECT user_id, COUNT(student_progress_level.id) as count_levels
			FROM student_progress_level 
			JOIN users ON users.id = student_progress_level.user_id
			WHERE users.class_id = $1 AND finished_at is not null
			GROUP BY user_id
		)
		SELECT
			users.id,
			users.fullname,
			COALESCE(ROUND(SUM(CASE WHEN attempts.correct_answer = attempts.student_answer THEN 1 ELSE 0 END) * 100.0 / NULLIF(COUNT(attempts.id), 0), 2), 0) as accuracy,
			COALESCE(count_levels, 0)
		FROM users	
		LEFT JOIN attempts ON attempts.student_id = users.id
		LEFT JOIN levels_complited ON levels_complited.user_id = users.id
		WHERE users.class_id = $1 AND users.role_id = $2
		GROUP BY users.id, users.fullname, count_levels;
	`

	rows, err := r.db.QueryContext(ctx, query, classId, internal.StudentRoleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := make([]model.StudentShortStats, 0)
	for rows.Next() {
		var stat model.StudentShortStats

		err := rows.Scan(
			&stat.StudentId,
			&stat.Name,
			&stat.Accuracy,
			&stat.LevelsComplited,
		)
		if err != nil {
			return nil, err
		}

		stats = append(stats, stat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *EquationAttemptsRepositoryStruct) GetEquationTypeAccuracyByClassId(ctx context.Context, classId int) ([]model.EquationTypeStats, error) {
	query := `
		SELECT equation_types.name, COALESCE(SUM(CASE WHEN attempts.correct_answer = attempts.student_answer THEN 1 ELSE 0 END) * 100.0 / NULLIF(COUNT(attempts.id), 0), 0)
		FROM equation_types
		LEFT JOIN attempts ON attempts.equation_type_id = equation_types.id
		LEFT JOIN users ON users.id = attempts.student_id
		WHERE users.class_id = $1
		GROUP BY equation_types.name;
	`

	rows, err := r.db.QueryContext(ctx, query, classId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := make([]model.EquationTypeStats, 0)
	for rows.Next() {
		var stat model.EquationTypeStats

		err := rows.Scan(
			&stat.Type,
			&stat.Accuracy,
		)
		if err != nil {
			return nil, err
		}

		stats = append(stats, stat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *EquationAttemptsRepositoryStruct) GetTotalAttemptsByClassId(ctx context.Context, classId int) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM attempts
		JOIN users on users.id = attempts.student_id
		WHERE users.class_id = $1
		GROUP BY users.class_id;
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, classId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *EquationAttemptsRepositoryStruct) GetWrongAnswersByClassId(ctx context.Context, classId int) (int, error) {
	query := `
		SELECT count(*)
		FROM attempts
		JOIN users ON users.id = attempts.student_id
		WHERE users.class_id = $1 AND attempts.correct_answer <> attempts.student_answer
		GROUP BY users.class_id;
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, classId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *EquationAttemptsRepositoryStruct) GetExtendedEquationTypeStats(ctx context.Context, studentId int) ([]model.ExtendedEquationTypeStats, error) {
	query := `
		WITH stats AS (
			SELECT 
				equation_types.name,
				COUNT(*) AS total,
				SUM(CASE WHEN attempts.correct_answer = attempts.student_answer THEN 1 ELSE 0 END) AS right_ans,
				SUM(CASE WHEN attempts.correct_answer = attempts.student_answer THEN 0 ELSE 1 END) AS wrong
			FROM attempts
			JOIN equation_types ON attempts.equation_type_id = equation_types.id
			WHERE student_id = $1
			GROUP BY equation_types.name
		)
		SELECT 
			name,
			total,
			right_ans,
			wrong,
			(right_ans * 100.0 / NULLIF(total, 0)) AS accuracy
		FROM stats;
	`

	rows, err := r.db.QueryContext(ctx, query, studentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := make([]model.ExtendedEquationTypeStats, 0)
	for rows.Next() {
		var stat model.ExtendedEquationTypeStats

		err := rows.Scan(
			&stat.Type,
			&stat.Attempts,
			&stat.Correct,
			&stat.Wrong,
			&stat.Accuracy,
		)
		if err != nil {
			return nil, err
		}

		stats = append(stats, stat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}
