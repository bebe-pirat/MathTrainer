package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUser(ctx context.Context, e model.User) (int, error)

	UpdateUser(ctx context.Context, e model.User) (*model.User, error)
	BlockAndUnblockUser(ctx context.Context, id int) error
	UpdateLastLoginUser(ctx context.Context, id int, lastLogin time.Time) error

	DeleteUser(ctx context.Context, id int) error

	GetUserById(ctx context.Context, id int) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	UserExists(ctx context.Context, login string, password string) (bool, error)
}

type UserRepositoryStruct struct {
	db *sql.DB
}

func NewUserRepositoryStruct(db *sql.DB) *UserRepositoryStruct {
	return &UserRepositoryStruct{
		db: db,
	}
}

func (r *UserRepositoryStruct) CreateUser(ctx context.Context, e model.User) (int, error) {
	query := `
		INSERT INTO users (email, login, password_hash, role_id, blocked, fullname, class_id, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	var id int
	err := r.db.QueryRowContext(ctx, query).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (r *UserRepositoryStruct) UpdateUser(ctx context.Context, e model.User) (*model.User, error) {
	query := `
		UPDATE users 
		SET email = $1, login = $2, password_hash = $3, role_id = $4, blocked = $5, fullname = $6, class_id = $7
		WHERE id = $8
		RETURNING *
	`

	var user model.User
	err := r.db.QueryRowContext(ctx, query, e.Email, e.Login, e.PasswordHash, e.RoleId, e.Blocked, e.FullName, e.ClassId, e.Id).Scan(
		user.Id,
		user.Email,
		e.Login,
		e.PasswordHash,
		e.RoleId,
		e.Blocked,
		e.FullName,
		e.ClassId,
		e.CreatedAt,
		e.LastLogin,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryStruct) BlockAndUnblockUser(ctx context.Context, id int, blocked bool) error {
	query := `
		UPDATE classes 
		SET blocked = $1
		WHERE id = $2
	`

	res, err := r.db.ExecContext(ctx, query, blocked, id)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *UserRepositoryStruct) UpdateLastLoginUser(ctx context.Context, id int, lastLogin time.Time) error {
	query := `
		UPDATE users
		SET lastlogin = $1
		WHERE id = $2
	`

	res, err := r.db.ExecContext(ctx, query, lastLogin, id)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *UserRepositoryStruct) DeleteUser(ctx context.Context, id int) error {
	query := `
		DELETE FROM users 
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

func (r *UserRepositoryStruct) GetUserById(ctx context.Context, id int) (*model.User, error) {
	query := `
		SELECT id, email, login, Password_hash, Role_Id, Blocked, FullName, Class_Id, Created_at, lastlogin
		FROM users
		WHERE id = $1
	`

	var user model.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.Id,
		&user.Email,
		&user.Login,
		&user.PasswordHash,
		&user.RoleId,
		&user.Blocked,
		&user.FullName,
		&user.ClassId,
		&user.CreatedAt,
		&user.LastLogin,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryStruct) GetAllUsers(ctx context.Context) ([]model.User, error) {
	query := `
		SELECT id, email, login, Password_hash, Role_Id, Blocked, FullName, Class_Id, Created_at, lastlogin
		FROM users
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]model.User, 0)
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Login,
			&user.PasswordHash,
			&user.RoleId,
			&user.Blocked,
			&user.FullName,
			&user.ClassId,
			&user.CreatedAt,
			&user.LastLogin,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryStruct) UserExists(ctx context.Context, login string, password string) (bool, error) {
	query := `
		SELECT id, password 
		FROM users
		WHERE email = $1 OR login = $1
	`

	var id int
	var hashpassword []byte
	err := r.db.QueryRowContext(ctx, query, login).Scan(&id, &password)
	if err != nil {
		return false, err
	}

	if err = bcrypt.CompareHashAndPassword(hashpassword, []byte(password)); err != nil {
		return false, err
	}

	return true, nil
}
