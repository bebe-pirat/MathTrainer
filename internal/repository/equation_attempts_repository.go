package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type EquationAttemptsRepository interface {
	GetStudentSectionStats(ctx context.Context, studentId int, sectionId int) (map[int]float32, error)

	// // старое))))))))
	// CreateAttempt(ctx context.Context, e model.EquationAttempts) (int, error)
	// GetStudentAttempts(ctx context.Context, studentId int) ([]model.EquationAttempts, error)

	// // school stats
	// GetTotalAttemptsBySchoolId(ctx context.Context, schoolId int) (int, error)
	// GetWrongAnswersBySchoolId(ctx context.Context, schoolId int) (int, error)
	// GetClassesAccuracyBySchoolId(ctx context.Context, schoolId int) ([]model.ClassShortStats, error)
	// GetEquationTypeAccuracyBySchoolId(ctx context.Context, schoolId int) ([]model.EquationTypeStats, error)

	// // class stats
	// GetStudentsShortStatsByClassId(ctx context.Context, classId int) ([]model.StudentShortStats, error)
	// GetWrongAnswersByClassId(ctx context.Context, classId int) (int, error)
	// GetTotalAttemptsByClassId(ctx context.Context, classId int) (int, error)
	// GetEquationTypeAccuracyByClassId(ctx context.Context, classId int) ([]model.EquationTypeStats, error)

	// // student stats
	// GetErrorStats(ctx context.Context, studentId int) (int, error)
	// GetAllStats(ctx context.Context, studentId int) (int, error)
	// GetExtendedEquationTypeStats(ctx context.Context, studentId int) ([]model.ExtendedEquationTypeStats, error)
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
		WHERE student_id = $1, section_id = $2;
	`

	rows, err := r.db.QueryContext(ctx, query, studentId, sectionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shortStats map[int]float32
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

// func (r *EquationAttemptsRepositoryStruct) CreateAttempt(ctx context.Context, e model.EquationAttempts) (int, error) {
// 	query := `
// 		INSERT INTO equation_attempts(student_id, equation_id, given_answer, correct, attempted_at)
// 		VALUES($1, $2, $3, $4, $5, $6)
// 		RETURNING id
// 	`

// 	var id int
// 	err := r.db.QueryRowContext(ctx, query, e.StudentId, e.EquationId, e.GivenAnswer, e.Correct, e.AttemptedAt).Scan(&id)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return id, nil
// }

// func (r *EquationAttemptsRepositoryStruct) GetStudentAttempts(ctx context.Context, studentId int) ([]model.EquationAttempts, error) {
// 	query := `
// 		SELECT id, student_id, equation_id, given_answer, correct, attempted_at
// 		FROM equation_attempts
// 		WHERE student_id = $1
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

// func (r *EquationAttemptsRepositoryStruct) GetErrorStats(ctx context.Context, studentId int) (int, error) {
// 	query := `
// 		select count(*) from equation_attempts where correct = false and student_id = $1
// 	`

// 	var count int
// 	err := r.db.QueryRowContext(ctx, query, studentId).Scan(&count)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return count, nil
// }

// func (r *EquationAttemptsRepositoryStruct) GetAllStats(ctx context.Context, studentId int) (int, error) {
// 	query := `
// 		select count(*) from equation_attempts where student_id = $1
// 	`

// 	var count int
// 	err := r.db.QueryRowContext(ctx, query, studentId).Scan(&count)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return count, nil
// }

// func (r *EquationAttemptsRepositoryStruct) GetTotalAttemptsBySchoolId(ctx context.Context, schoolId int) (int, error) {
// 	query := `
// 		select count(*)
// 		from equation_attempts
// 		join users on users.id = equation_attempts.student_id
// 		join classes on classes = users.classes_id
// 		where classes.school_id = $1
// 	`

// 	var count int
// 	err := r.db.QueryRowContext(ctx, query, schoolId).Scan(&count)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return count, nil
// }

// func (r *EquationAttemptsRepositoryStruct) GetWrongAnswersBySchoolId(ctx context.Context, schoolId int) (int, error) {
// 	query := `
// 		select count(*)
// 		from equation_attempts
// 		join users on users.id = equation_attempts.student_id
// 		join classes on classes = users.classes_id
// 		where classes.school_id = $1 and correct = false
// 	`

// 	var count int
// 	err := r.db.QueryRowContext(ctx, query, schoolId).Scan(&count)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return count, nil
// }

// func (r *EquationAttemptsRepositoryStruct) GetClassesAccuracyBySchoolId(ctx context.Context, schoolId int) ([]model.ClassShortStats, error) {
// 	query := `
// 		SELECT classes.name, COALESCE(SUM(CASE WHEN equation_attempts.correct THEN 1 ELSE 0 END) * 100.0 / NULLIF(COUNT(equation_attempts.id), 0), 0)
// 		from equation_attempts
// 		join users on users.id = equation_attempts.student_id
// 		join classes on classes.id = users.classes_id
// 		where classes.school_id = $1
// 	`

// 	rows, err := r.db.QueryContext(ctx, query, schoolId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	stats := make([]model.ClassShortStats, 0)
// 	for rows.Next() {
// 		var stat model.ClassShortStats

// 		err := rows.Scan(
// 			&stat.Name,
// 			&stat.Accuracy,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		stats = append(stats, stat)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return stats, nil
// }

// func (r *EquationAttemptsRepositoryStruct) GetEquationTypeAccuracyBySchoolId(ctx context.Context, schoolId int) ([]model.EquationTypeStats, error) {
// 	query := `
// 		SELECT equation_types.name, COALESCE(SUM(CASE WHEN equation_attempts.correct THEN 1 ELSE 0 END) * 100.0 / NULLIF(COUNT(equation_attempts.id), 0), 0)
// 		from equation_attempts
// 		join equations on equations.id = equation_attempts.equation_id
// 		join equation_types on equations.equaition_type_id = equation_types.id
// 		join users on equation_attempts.student_id = users.id
// 		join classes on users.class_id = classes.id
// 		where classes.school_id = $1
// 		group by equation_types.name
// 	`

// 	rows, err := r.db.QueryContext(ctx, query, schoolId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	stats := make([]model.EquationTypeStats, 0)
// 	for rows.Next() {
// 		var stat model.EquationTypeStats

// 		err := rows.Scan(
// 			&stat.Type,
// 			&stat.Accuracy,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		stats = append(stats, stat)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return stats, nil
// }
// func (r *EquationAttemptsRepositoryStruct) GetStudentsShortStatsByClassId(ctx context.Context, classId int) ([]model.StudentShortStats, error) {
// 	query := `
// 		SELECT
// 			users.id,
// 			users.fullname,
// 			COALESCE(ROUND(SUM(CASE WHEN equation_attempts.correct THEN 1 ELSE 0 END) * 100.0 / NULLIF(COUNT(equation_attempts.id), 0), 2), 0) as accuracy,
// 			COUNT(DISTINCT equation_attempts.level_id) as levels_completed
// 		FROM users
// 		LEFT JOIN equation_attempts ON equation_attempts.student_id = users.id
// 		WHERE users.class_id = $1
// 		GROUP BY users.id, users.fullname
// 	`

// 	rows, err := r.db.QueryContext(ctx, query, classId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	stats := make([]model.StudentShortStats, 0)
// 	for rows.Next() {
// 		var stat model.StudentShortStats

// 		err := rows.Scan(
// 			&stat.StudentId,
// 			&stat.Name,
// 			&stat.Accuracy,
// 			&stat.LevelsComplited,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		stats = append(stats, stat)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return stats, nil
// }

// func (r *EquationAttemptsRepositoryStruct) GetEquationTypeAccuracyByClassId(ctx context.Context, classId int) ([]model.EquationTypeStats, error) {
// 	query := `
// 		SELECT equation_types.name, COALESCE(SUM(CASE WHEN equation_attempts.correct THEN 1 ELSE 0 END) * 100.0 / NULLIF(COUNT(equation_attempts.id), 0), 0)
// 		from equation_attempts
// 		join equations on equations.id = equation_attempts.equation_id
// 		join equation_types on equations.equaition_type_id = equation_types.id
// 		join users on users.id = equation_attempts.student_id
// 		where users.class_id = $1
// 		group by equation_types.name
// 	`

// 	rows, err := r.db.QueryContext(ctx, query, classId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	stats := make([]model.EquationTypeStats, 0)
// 	for rows.Next() {
// 		var stat model.EquationTypeStats

// 		err := rows.Scan(
// 			&stat.Type,
// 			&stat.Accuracy,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		stats = append(stats, stat)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return stats, nil
// }

// func (r *EquationAttemptsRepositoryStruct) GetTotalAttemptsByClassId(ctx context.Context, classId int) (int, error) {
// 	query := `
// 		select count(*)
// 		from equation_attempts
// 		join users on users.id = equation_attempts.student_id
// 		where users.class_id = $1
// 	`

// 	var count int
// 	err := r.db.QueryRowContext(ctx, query, classId).Scan(&count)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return count, nil
// }

// func (r *EquationAttemptsRepositoryStruct) GetWrongAnswersByClassId(ctx context.Context, classId int) (int, error) {
// 	query := `
// 		select count(*)
// 		from equation_attempts
// 		join users on users.id = equation_attempts.student_id
// 		where users.class_id = $1 and correct = false
// 	`

// 	var count int
// 	err := r.db.QueryRowContext(ctx, query, classId).Scan(&count)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return count, nil
// }

// func (r *EquationAttemptsRepositoryStruct) GetExtendedEquationTypeStats(ctx context.Context, studentId int) ([]model.ExtendedEquationTypeStats, error) {
// 	query := `
// 		SELECT equation_types.name, COUNT(*) as total, SUM(CASE WHEN equation_attempts.correct THEN 1 ELSE 0) as right,
// 		SUM(CASE WHEN equation_attempts.correct THEN 0 ELSE 1) as wrong,
// 		(right * 100.0 / NULLIF(total, 0)) as accuracy
// 		from equation_attempts
// 		join equations on equations.id = equation_attempts.equation_id
// 		join equation_types on equations.equaition_type_id = equation_types.id
// 		where student_id = $1
// 		group by equation_types.name
// 	`

// 	rows, err := r.db.QueryContext(ctx, query, studentId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	stats := make([]model.ExtendedEquationTypeStats, 0)
// 	for rows.Next() {
// 		var stat model.ExtendedEquationTypeStats

// 		err := rows.Scan(
// 			&stat.Type,
// 			&stat.Attempts,
// 			&stat.Correct,
// 			&stat.Wrong,
// 			&stat.Accuracy,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		stats = append(stats, stat)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return stats, nil
// }
