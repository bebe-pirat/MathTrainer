package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type RoleRepository interface {
	GetAllRoles(ctx context.Context) ([]model.Role, error)
}

type RoleRepositoryStruct struct {
	db *sql.DB
}

func NewRoleRepositoryStruct(db *sql.DB) *RoleRepositoryStruct {
	return &RoleRepositoryStruct{
		db: db,
	}
}

func (r *RoleRepositoryStruct) GetAllRoles(ctx context.Context) ([]model.Role, error) {
	query := `
		SELECT id, name 
		FROM roles
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := make([]model.Role, 0)
	for rows.Next() {
		var role model.Role

		err := rows.Scan(
			&role.Id,
			&role.Name,
		)

		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, err
}
