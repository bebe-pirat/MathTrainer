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

	CreateEquationType(ctx context.Context, equationType model.EquationType) error
	UpdateEquationType(ctx context.Context, equationType model.EquationType) error
	DeleteEquationType(ctx context.Context, equationTypeId int) error
	GetEquationTypes(ctx context.Context) ([]model.EquationType, error)

	JoinEquationTypeToSection(ctx context.Context, equationTypeId, sectionId int) error
	CreateOperandForEquationType(ctx context.Context, operand model.Operand) error
	UpdateOperandForEquationType(ctx context.Context, operand model.Operand) error
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

func (r *EquationTypeRepositoryStruct) CreateEquationType(ctx context.Context, equationType model.EquationType) error {
	query := `
		INSERT INTO equation_types (class, name, description, operations, num_operands, no_remainder, max_result)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	res, err := r.db.ExecContext(ctx, query, equationType.Class, equationType.Name, equationType.Description, equationType.Operations, equationType.NumOperands, equationType.NoRemainder, equationType.MaxResult)
	if err != nil {
		return err
	}

	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *EquationTypeRepositoryStruct) UpdateEquationType(ctx context.Context, equationType model.EquationType) error {
	query := `
		UPDATE equation_types SET class = $1, 
									name = $2, 
									description = $3, 
									operations = $4, 
									num_operands = $5, 
									no_remainder = $6, 
									max_result = $7
		WHERE id = $8;
	`

	res, err := r.db.ExecContext(ctx, query, equationType.Class, equationType.Name, equationType.Description, equationType.Operations, equationType.NumOperands, equationType.NoRemainder, equationType.MaxResult, equationType.Id)
	if err != nil {
		return err
	}

	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *EquationTypeRepositoryStruct) DeleteEquationType(ctx context.Context, equationTypeId int) error {
	query := `
		DELETE FROM	equation_types 
		WHERE id = $1;
	`

	res, err := r.db.ExecContext(ctx, query, equationTypeId)
	if err != nil {
		return err
	}

	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *EquationTypeRepositoryStruct) GetEquationTypes(ctx context.Context) ([]model.EquationType, error) {
	query := `
		SELECT id, class, name, description, operations, num_operands, no_remainder, max_result;
		FROM equation_types
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []model.EquationType
	for rows.Next() {
		var eqType model.EquationType

		err := rows.Scan(
			&eqType.Id,
			&eqType.Class,
			&eqType.Name,
			&eqType.Description,
			&eqType.Operations,
			&eqType.NumOperands,
			&eqType.NoRemainder,
			&eqType.MaxResult,
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

func (r *EquationTypeRepositoryStruct) JoinEquationTypeToSection(ctx context.Context, equationTypeId, sectionId int) error {
	query := `
		INSERT INTO section_equation_types(section_id, eqaution_type_id) 
		VALUES ($1, $2);
	`
	res, err := r.db.ExecContext(ctx, query, sectionId, equationTypeId)
	if err != nil {
		return err
	}

	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *EquationTypeRepositoryStruct) CreateOperandForEquationType(ctx context.Context, operand model.Operand) error {
	query := `
		INSERT INTO operands(equation_type_id, operand_order, min_value, max_value) 
		VALUES ($1, $2, $3, $4);
	`
	res, err := r.db.ExecContext(ctx, query, operand.EquationTypeId, operand.OperandOrder, operand.MinValue, operand.MaxValue)
	if err != nil {
		return err
	}
	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *EquationTypeRepositoryStruct) UpdateOperandForEquationType(ctx context.Context, operand model.Operand) error {
	query := `
		UPDATE operands SET operand_order = $1, min_value = $4, 
							max_value = $3
		WHERE id = $4; 
	`
	res, err := r.db.ExecContext(ctx, query, operand.OperandOrder, operand.MinValue, operand.MaxValue, operand.Id)
	if err != nil {
		return err
	}
	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
