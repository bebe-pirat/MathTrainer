package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
	"fmt"
)

type EquationTypeRepository interface {
	GetEquationTypesBySection(ctx context.Context, sectionId int) ([]model.EquationTypeWithOperands, error)
	GetOperandsByEquationType(ctx context.Context, equationTypeId int) ([]model.Operand, error)
	GetEquationTypesByStudentId(ctx context.Context, studentId int) ([]model.ShortEquationType, error)

	// TODO: добавить метод для создания типа уравнения + внесение запписи в таблицу операндов, можно еще обновление и удаление
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
		JOIN section_equation_types ON equation_types.id = section_equation_types.eqaution_type_id
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
		return nil, fmt.Errorf("error here: %w", err)
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
		return nil, fmt.Errorf("poerand get error: %w", err)
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
			return nil, fmt.Errorf("poerand get error: %w", err)
		}

		operands = append(operands, operand)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("poerand get error: %w", err)
	}

	return operands, nil
}

func (r *EquationTypeRepositoryStruct) GetEquationTypesByStudentId(ctx context.Context, studentId int) ([]model.ShortEquationType, error) {
	query := `
		SELECT equation_types.id, equation_types.name
		FROM equation_types
		JOIN classes ON classes.grade = equation_types.class
		JOIN users ON classes.id = users.class_id
		WHERE users.id = $1;
	`

	rows, err := r.db.QueryContext(ctx, query, studentId)
	if err != nil {
		return nil, fmt.Errorf("poerand get error: %w", err)
	}
	defer rows.Close()

	var types []model.ShortEquationType
	for rows.Next() {
		var eqType model.ShortEquationType

		err := rows.Scan(
			&eqType.Id,
			&eqType.Name,
		)

		if err != nil {
			return nil, fmt.Errorf("types get error: %w", err)
		}

		types = append(types, eqType)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("types get error: %w", err)
	}

	return types, nil
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
