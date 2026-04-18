package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type EquationTypeRepository interface {
	GetEquationTypesBySection(ctx context.Context, sectionId int) ([]model.EquationTypeWithOperands, error)
	GetOperandsByEquationType(ctx context.Context, equationTypeId int) ([]model.Operand, error)

	// old and maybe unnessesary
	// TODO: при ненадобности удалиьт нахер
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

func (r *EquationTypeRepositoryStruct) GetEquationTypesBySection(ctx context.Context, sectionId int) ([]model.EquationTypeWithOperands, error) {
	query := `
		SELECT equation_types.id, operations, num_operands, no_remainder, max_result
		FROM equation_types
		JOIN section_equation_types ON equation_types.id = section_equation_types.equation_type_id
		WHERE section_id = $1;
	`

	rows, err := r.db.QueryContext(ctx, query, sectionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equationTypes []model.EquationTypeWithOperands
	for rows.Next() {
		var equationType model.EquationTypeWithOperands

		err := rows.Scan(
			&equationType.Id,
			&equationType.Operations,
			&equationType.NumOperands,
			&equationType.NoRemainder,
			&equationType.MaxResult,
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

func (r *EquationTypeRepositoryStruct) GetOperandsByEquationType(ctx context.Context, equationTypeId int) ([]model.Operand, error) {
	query := `
		SELECT id, operand_order, min_value, max_value
		FROM operands 
		WHERE equation_type_id = $1;
	`

	rows, err := r.db.QueryContext(ctx, query, equationTypeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var operands []model.Operand
	for rows.Next() {
		var operand model.Operand

		err := rows.Scan(
			&operand.Id,
			&operand.OperandOrder,
			&operand.MinValue,
			&operand.MaxValue,
		)

		if err != nil {
			return nil, err
		}

		operands = append(operands, operand)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return operands, nil
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
