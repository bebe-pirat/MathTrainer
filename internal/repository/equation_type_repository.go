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

	CreateEquationType(ctx context.Context, tx *sql.Tx, equationType model.EquationType) (int, error)
	UpdateEquationType(ctx context.Context, tx *sql.Tx, equationType model.EquationType) error
	DeleteEquationType(ctx context.Context, tx *sql.Tx, equationTypeId int) error
	GetEquationTypes(ctx context.Context) ([]model.EquationType, error)

	JoinEquationTypeToSection(ctx context.Context, equationTypeId, sectionId int) error
	UnJoinEquationTypeToSection(ctx context.Context, equationTypeId, sectionId int) error
	DeleteEquationTypeFromSection(ctx context.Context, tx *sql.Tx, equationTypeId int) error
	GetSectionsAndEquationTypes(ctx context.Context) ([]model.SectionAndEquationType, error)

	CreateOperandForEquationType(ctx context.Context, tx *sql.Tx, operand model.Operand) error
	UpdateOperandForEquationType(ctx context.Context, tx *sql.Tx, operand model.Operand) error
	DeleteOperandsForEquationType(ctx context.Context, tx *sql.Tx, equationTypeId int) error
	GetOperandsForEquationTypeId(ctx context.Context, equationTypeId int) ([]model.OperandResponse, error)

	CreateFullEquationType(ctx context.Context, equationType model.EquationType, operands []model.Operand) error
	DeleteFullEquationType(ctx context.Context, equationTypeId int) error
	UpdateFullEquationType(ctx context.Context, equationType model.EquationType, operands []model.Operand, equationTypeId int) error
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

func (r *EquationTypeRepositoryStruct) CreateEquationType(ctx context.Context, tx *sql.Tx, equationType model.EquationType) (int, error) {
	query := `
		INSERT INTO equation_types (class, name, description, operations, num_operands, no_remainder, max_result)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`

	var id int
	err := tx.QueryRowContext(ctx, query, equationType.Class, equationType.Name, equationType.Description, equationType.Operations, equationType.NumOperands, equationType.NoRemainder, equationType.MaxResult).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *EquationTypeRepositoryStruct) UpdateEquationType(ctx context.Context, tx *sql.Tx, equationType model.EquationType) error {
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

	res, err := tx.ExecContext(ctx, query, equationType.Class, equationType.Name, equationType.Description, equationType.Operations, equationType.NumOperands, equationType.NoRemainder, equationType.MaxResult, equationType.Id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *EquationTypeRepositoryStruct) DeleteEquationType(ctx context.Context, tx *sql.Tx, equationTypeId int) error {
	query := `
		DELETE FROM	equation_types 
		WHERE id = $1;
	`

	res, err := tx.ExecContext(ctx, query, equationTypeId)
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
		SELECT equation_types.id, equation_types.class, equation_types.name, description, operations, num_operands, no_remainder, max_result
		FROM equation_types;
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

func (r *EquationTypeRepositoryStruct) UnJoinEquationTypeToSection(ctx context.Context, equationTypeId, sectionId int) error {
	query := `
		DELETE FROM section_equation_types
		WHERE section_id = $1 AND eqaution_type_id = $2;
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

func (r *EquationTypeRepositoryStruct) DeleteEquationTypeFromSection(ctx context.Context, tx *sql.Tx, equationTypeId int) error {
	query := `
		DELETE FROM section_equation_types
		WHERE eqaution_type_id = $1;
	`
	res, err := tx.ExecContext(ctx, query, equationTypeId)
	if err != nil {
		return err
	}

	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *EquationTypeRepositoryStruct) CreateOperandForEquationType(ctx context.Context, tx *sql.Tx, operand model.Operand) error {
	query := `
		INSERT INTO operands(equation_type_id, operand_order, min_value, max_value) 
		VALUES ($1, $2, $3, $4);
	`
	res, err := tx.ExecContext(ctx, query, operand.EquationTypeId, operand.OperandOrder, operand.MinValue, operand.MaxValue)
	if err != nil {
		return err
	}
	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *EquationTypeRepositoryStruct) UpdateOperandForEquationType(ctx context.Context, tx *sql.Tx, operand model.Operand) error {
	query := `
		UPDATE operands SET operand_order = $1, min_value = $2, 
							max_value = $3
		WHERE id = $4; 
	`
	res, err := tx.ExecContext(ctx, query, operand.OperandOrder, operand.MinValue, operand.MaxValue, operand.Id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return nil
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *EquationTypeRepositoryStruct) DeleteOperandsForEquationType(ctx context.Context, tx *sql.Tx, equationTypeId int) error {
	query := `
		DELETE FROM operands 
		WHERE equation_type_id = $1;
	`

	res, err := tx.ExecContext(ctx, query, equationTypeId)
	if err != nil {
		return err
	}
	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *EquationTypeRepositoryStruct) GetOperandsForEquationTypeId(ctx context.Context, equationTypeId int) ([]model.OperandResponse, error) {
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

	var types []model.OperandResponse
	for rows.Next() {
		var eqType model.OperandResponse

		err := rows.Scan(
			&eqType.Id,
			&eqType.OperandOrder,
			&eqType.MinValue,
			&eqType.MaxValue,
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

func (r *EquationTypeRepositoryStruct) CreateFullEquationType(ctx context.Context, equationType model.EquationType, operands []model.Operand) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	equationTypeId, err := r.CreateEquationType(ctx, tx, equationType)
	if err != nil {
		return fmt.Errorf("create equation type: %w", err)
	}

	for _, operand := range operands {
		operand.EquationTypeId = equationTypeId
		if err := r.CreateOperandForEquationType(ctx, tx, operand); err != nil {
			return fmt.Errorf("create operand (order %d): %w", operand.OperandOrder, err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func (r *EquationTypeRepositoryStruct) DeleteFullEquationType(ctx context.Context, equationTypeId int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	err = r.DeleteOperandsForEquationType(ctx, tx, equationTypeId)
	if err != nil {
		return fmt.Errorf("delete operands: %w", err)
	}

	if err := r.DeleteEquationTypeFromSection(ctx, tx, equationTypeId); err != nil {
		return fmt.Errorf("delete equation type from section: %w", err)
	}

	if err := r.DeleteEquationType(ctx, tx, equationTypeId); err != nil {
		return fmt.Errorf("delete equation type: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func (r *EquationTypeRepositoryStruct) UpdateFullEquationType(ctx context.Context, equationType model.EquationType, operands []model.Operand, equationTypeId int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	err = r.UpdateEquationType(ctx, tx, equationType)
	if err != nil {
		return fmt.Errorf("update equation type: %w", err)
	}

	for _, operand := range operands {
		operand.EquationTypeId = equationTypeId
		if err := r.UpdateOperandForEquationType(ctx, tx, operand); err != nil {
			return fmt.Errorf("update operand (order %d, id: %d): %w", operand.OperandOrder, operand.Id, err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func (r *EquationTypeRepositoryStruct) GetSectionsAndEquationTypes(ctx context.Context) ([]model.SectionAndEquationType, error) {
	query := `
		SELECT section_id, sections.name, eqaution_type_id, equation_types.name, sections.class
		FROM sections 
		JOIN section_equation_types ON sections.id = section_equation_types.section_id
		JOIN equation_types ON equation_types.id = section_equation_types.eqaution_type_id;
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	secAndEqs := make([]model.SectionAndEquationType, 0)
	for rows.Next() {
		var secAndEq model.SectionAndEquationType

		err := rows.Scan(
			&secAndEq.SectionId,
			&secAndEq.SectionName,
			&secAndEq.EquationTypeId,
			&secAndEq.EquationTypeName,
			&secAndEq.Class,
		)
		if err != nil {
			return nil, err
		}

		secAndEqs = append(secAndEqs, secAndEq)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return secAndEqs, nil
}
