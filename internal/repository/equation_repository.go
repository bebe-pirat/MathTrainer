package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type EquationRepository interface {
	GetRandomEquation(ctx context.Context, levelId int) (*model.Equation, error)
	GetRandomEquationByEquationTypeId(ctx context.Context, levelId int, equation_type_id int) (*model.Equation, error)
	GetEquationById(ctx context.Context, id int) (*model.Equation, error)
}

type EquationRepositoryStruct struct {
	db *sql.DB
}

func NewEquationRepositoryStruct(db *sql.DB) *EquationRepositoryStruct {
	return &EquationRepositoryStruct{
		db: db,
	}
}

func (r *EquationRepositoryStruct) GetRandomEquation(ctx context.Context, levelId int) (*model.Equation, error) {
	query := `
		SELECT Id, expression, correct_answer, equation_type_id, difficulty
		FROM equations 
		JOIN level_equation_type ON equations.equation_type_id = level_equation_type.equation_type_id
		WHERE level_equation_type.level_id = $1
		ORDER BY RANDOM()
		LIMIT 1;
	`

	var equation model.Equation
	err := r.db.QueryRowContext(ctx, query, levelId).Scan(
		&equation.Id,
		&equation.Expression,
		&equation.CorrectAnswer,
		&equation.EquationTypeId,
		equation.Difficulty,
	)
	if err != nil {
		return nil, err
	}

	return &equation, nil
}

func (r *EquationRepositoryStruct) GetRandomEquationByEquationTypeId(ctx context.Context, levelId int, equation_type_id int) (*model.Equation, error) {
	query := `
		SELECT Id, expression, correct_answer, equation_type_id, difficulty
		FROM equations 
		JOIN level_equation_type ON equations.equation_type_id = level_equation_type.equation_type_id
		WHERE level_equation_type.level_id = $1 AND equation_type_id = $2
		ORDER BY RANDOM()
		LIMIT 1;
	`

	var equation model.Equation
	err := r.db.QueryRowContext(ctx, query, levelId, equation_type_id).Scan(
		&equation.Id,
		&equation.Expression,
		&equation.CorrectAnswer,
		&equation.EquationTypeId,
		equation.Difficulty,
	)
	if err != nil {
		return nil, err
	}

	return &equation, nil
}

func (r *EquationRepositoryStruct) GetEquationById(ctx context.Context, id int) (*model.Equation, error) {
	query :=
		`SELECT id, expression, correct_answer, equation_type_id, difficulty
		FROM equation
		WHERE id = $1
		`

	var equation model.Equation
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&equation.Id,
		&equation.Expression,
		&equation.CorrectAnswer,
		&equation.EquationTypeId,
		equation.Difficulty,
	)
	if err != nil {
		return nil, err
	}

	return &equation, nil
}
