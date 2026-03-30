package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type SchoolRepository interface {
	CreateSchool(ctx context.Context, e model.School) error
	UpdateSchool(ctx context.Context, e model.School) (*model.School, error)
	DeleteSchool(ctx context.Context, id int) error
	GetAllSchools(ctx context.Context) ([]model.School, error)
	GetSchoolById(ctx context.Context, id int) (*model.School, error)
}

type SchoolRepositoryStruct struct {
	db *sql.DB
}

func NewSchoolRepositoryStruct(db *sql.DB) *SchoolRepositoryStruct {
	return &SchoolRepositoryStruct{
		db: db,
	}
}

func (r *SchoolRepositoryStruct) CreateSchool(ctx context.Context, e model.School) error {
	query := `
			INSERT INTO schools(name, address, created_at)
			VALUES($1, $2, $3)`

	res, err := r.db.ExecContext(ctx, query, e.Name, e.Address, e.Created_at)
	if err != nil {
		return nil
	}
	if rows, err := res.RowsAffected(); err != nil && rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *SchoolRepositoryStruct) UpdateSchool(ctx context.Context, e model.School) (*model.School, error) {
	query := `
			UPDATE schools 
			SET name = $1, address = $2
			RETURNING *`

	var school model.School
	err := r.db.QueryRowContext(ctx, query, e.Name, e.Address).Scan(&school.Id, &school.Name, &school.Address, &school.Created_at)
	if err != nil {
		return nil, nil
	}

	return &school, nil
}

func (r *SchoolRepositoryStruct) DeleteSchool(ctx context.Context, id int) error {
	query := `
			DELETE FROM schools
			WHERE id = $1`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *SchoolRepositoryStruct) GetSchoolById(ctx context.Context, id int) (*model.School, error) {
	query := `
			SELECT id, name, address, created_at 
			FROM schools
			WHERE id = $1`

	var school model.School
	err := r.db.QueryRowContext(ctx, query, id).Scan(&school.Id, &school.Name, &school.Address, &school.Created_at)
	if err != nil {
		return nil, err
	}

	return &school, nil
}

func (r *SchoolRepositoryStruct) GetAllSchools(ctx context.Context) ([]model.School, error) {
	query := `
			SELECT id, name, address, created_at 
			FROM schools`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schools []model.School
	for rows.Next() {
		var school model.School
		err := rows.Scan(
			&school.Id,
			&school.Name,
			&school.Address,
			&school.Created_at,
		)

		if err != nil {
			return nil, err
		}

		schools = append(schools, school)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schools, nil
}
