package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type ClassRepository interface {
	CreateClass(ctx context.Context, e model.Class) (int, error)
	UpdateClass(ctx context.Context, e model.Class) (*model.Class, error)
	DeleteClass(ctx context.Context, id int) error

	GetClassById(ctx context.Context, id int) (*model.Class, error)
	GetAllClasses(ctx context.Context) ([]model.Class, error)
	GetClassesBySchoolId(ctx context.Context, schoolId int) ([]model.Class, error)
}

type ClassRepositoryStruct struct {
	db *sql.DB
}

func NewClassRepositoryStruct(db *sql.DB) *ClassRepositoryStruct {
	return &ClassRepositoryStruct{
		db: db,
	}
}

func (r *ClassRepositoryStruct) CreateClass(ctx context.Context, e model.Class) (int, error) {
	query := `
			INSERT INTO classes(name, grade, school_id, created_at) 
			VALUES ($1, $2, $3, $4)
			RETURNING id`

	var id int
	err := r.db.QueryRowContext(ctx, query, e.Name, e.Grade, e.SchoolId, e.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ClassRepositoryStruct) UpdateClass(ctx context.Context, e model.Class) (*model.Class, error) {
	query := `
			UPDATE classes 
			SET name = $1, grade = $2, school_id = $3
			WHERE id = $4
			RETURNING *`

	var class model.Class
	err := r.db.QueryRowContext(ctx, query, e.Name, e.Grade, e.SchoolId).Scan(&class.Id, &class.Name, &class.Grade, &class.SchoolId, e.CreatedAt)
	if err != nil {
		return nil, nil
	}

	return &class, nil
}

func (r *ClassRepositoryStruct) DeleteClass(ctx context.Context, id int) error {
	query := `
			DELETE FROM classes 
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

func (r *ClassRepositoryStruct) GetClassById(ctx context.Context, id int) (*model.Class, error) {
	query := `
		SELECT id, name, grade, school_id, created_at
		FROM classes
		WHERE id = $1
	`

	var class model.Class
	err := r.db.QueryRowContext(ctx, query, id).Scan(&class.Id, &class.Name, &class.Grade, &class.SchoolId, &class.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &class, nil
}

func (r *ClassRepositoryStruct) GetAllClasses(ctx context.Context) ([]model.Class, error) {
	query := `
		SELECT id , name, grade, school_id, created_at
		FROM classes
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []model.Class
	for rows.Next() {
		var class model.Class
		err := rows.Scan(
			&class.Id,
			&class.Name,
			&class.Grade,
			&class.SchoolId,
			&class.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		classes = append(classes, class)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return classes, nil
}

func (r *ClassRepositoryStruct) GetClassesBySchoolId(ctx context.Context, schoolId int) ([]model.Class, error) {
	query := `
		SELECT id , name, grade, school_id, created_at
		FROM classes
		WHERE school_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, schoolId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []model.Class
	for rows.Next() {
		var class model.Class
		err := rows.Scan(
			&class.Id,
			&class.Name,
			&class.Grade,
			&class.SchoolId,
			&class.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		classes = append(classes, class)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return classes, nil
}
