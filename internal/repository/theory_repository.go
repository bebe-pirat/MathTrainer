package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type TheoryRepository interface {
	GetByEquationType(ctx context.Context, typeId int) (*model.Theory, error)
}

type TheoryRepositoryStruct struct {
	db *sql.DB
}

func NewTheoryRepositoryStruct(db *sql.DB) *TheoryRepositoryStruct {
	return &TheoryRepositoryStruct{
		db: db,
	}
}

func (r *TheoryRepositoryStruct) GetByEquationType(ctx context.Context, typeId int) (*model.Theory, error) {
	query := `
		SELECT id, equation_type_id, name, content
		FROM theory 
		WHERE equation_type_id = $1
	`

	var theory model.Theory
	err := r.db.QueryRowContext(ctx, query).Scan(&theory.Id, &theory.EquationTypeId, &theory.Name, &theory.Content)
	if err != nil {
		return nil, err
	}

	return &theory, nil
}
