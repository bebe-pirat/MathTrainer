package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type EquationTypeRepository interface {
	GetAllEquationTypes(ctx context.Context) ([]model.EquationType, error)
	GetEquationTypeById(ctx context.Context) (*model.EquationType, error)
	GetEquationTypesByLevelId(ctx context.Context, levelId int) ([]model.EquationType, error)
}

type EquationTypeRepositoryStruct struct {
	db *sql.DB
}

func NewEquationTypeRepositoryStruct(db *sql.DB) *EquationTypeRepositoryStruct {
	return &EquationTypeRepositoryStruct{
		db: db,
	}
}

func (r *EquationTypeRepositoryStruct) GetAllEquationTypes(ctx context.Context) ([]model.EquationType, error) {
	query := `
	SELECT name, description
	FROM equation_types
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	equationTypes := make([]model.EquationType, 0)
	for rows.Next() {
		var equationType model.EquationType
		err := rows.Scan(
			&equationType.Id,
			&equationType.Name,
			&equationType.Description,
		)

		if err != nil {
			return nil, err
		}

		equationTypes = append(equationTypes, equationType)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return equationTypes, nil
}

func (r *EquationTypeRepositoryStruct) GetEquationTypeById(ctx context.Context, id int) (*model.EquationType, error) {
	query := `
	SELECT name, description
	FROM equation_types
	WHERE id = $1
	`

	var equationType model.EquationType
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&equationType.Id,
		&equationType.Name,
		&equationType.Description,
	)
	if err != nil {
		return nil, err
	}

	return &equationType, nil
}

func (r *EquationAttemptsRepositoryStruct) GetEquationTypesByLevelId(ctx context.Context, levelId int) ([]model.EquationType, error) {
	query := `
		SELECT id, name, description
		FROM equation_types JOIN level_equation_type ON equation_types.id = level_equation_types.equation_type_id
		WHERE level_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, levelId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	equationTypes := make([]model.EquationType, 0)
	for rows.Next() {
		var equationType model.EquationType
		err := rows.Scan(
			&equationType.Id,
			&equationType.Name,
			&equationType.Description,
		)

		if err != nil {
			return nil, err
		}

		equationTypes = append(equationTypes, equationType)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return equationTypes, nil
}
