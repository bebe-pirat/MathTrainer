package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUser(ctx context.Context, e model.User) (string, error)

	UpdateUser(ctx context.Context, e model.User) (*model.User, error)
	UpdateLastLoginUser(ctx context.Context, id int, lastLogin time.Time) error
	BlockUser(ctx context.Context, id int, blocked bool) error

	DeleteUser(ctx context.Context, id int) error

	GetUserById(ctx context.Context, id int) (*model.User, error)
	GetRoleById(ctx context.Context, id int) (int, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	UserExists(ctx context.Context, login string, password string) (int, error)
	GetStudentsByClass(ctx context.Context, classID int) ([]model.User, error)
	GetTeachersBySchool(ctx context.Context, schoolID int) ([]model.User, error)
	GetClassByUserId(ctx context.Context, userId int) (int, error)
	GetGradeByStudentId(ctx context.Context, studentId int) (int, error)

	GetTotalStudentBySchoolId(ctx context.Context, schoolId int) (int, error)

	// student page
	GetStudentProfileById(ctx context.Context, id int) (*model.StudentProfile, error)

	AddXP(ctx context.Context, studentId int, xp int) error
	GetStudentXP(ctx context.Context, studentId int) (int, error)
}

type UserRepositoryStruct struct {
	db *sql.DB
}

func NewUserRepositoryStruct(db *sql.DB) *UserRepositoryStruct {
	return &UserRepositoryStruct{
		db: db,
	}
}

func (r *UserRepositoryStruct) GetGradeByStudentId(ctx context.Context, studentId int) (int, error) {
	query := `
		SELECT grade 
		FROM users
		JOIN classes ON users.class_id = classes.id
		WHERE users.id = $1;
	`

	var grade int
	err := r.db.QueryRowContext(ctx, query, studentId).Scan(&grade)
	if err != nil {
		return 0, err
	}

	return grade, nil
}

func (r *UserRepositoryStruct) GetClassByUserId(ctx context.Context, userId int) (int, error) {
	query := `
		SELECT class_id FROM users
		WHERE id = $1;
	`

	var classId int
	err := r.db.QueryRowContext(ctx, query, userId).Scan(&classId)
	if err != nil {
		return 0, err
	}

	return classId, nil
}

func (r *UserRepositoryStruct) GetRoleById(ctx context.Context, id int) (int, error) {
	query := ` 
		SELECT role_id FROM users
		WHERE id = $1;
	`

	var roleId int
	err := r.db.QueryRowContext(ctx, query, id).Scan(&roleId)
	if err != nil {
		return 0, err
	}

	return roleId, nil
}

func (r *UserRepositoryStruct) CreateUser(ctx context.Context, e model.User) (string, error) {
	query := `
		INSERT INTO users (email, login, password_hash, role_id, blocked, fullname, class_id, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING login
	`

	var login string
	err := r.db.QueryRowContext(ctx, query, e.Email, e.Login, e.PasswordHash, e.RoleId, e.Blocked, e.FullName, e.ClassId, e.CreatedAt).Scan(&login)
	if err != nil {
		return "", err
	}

	return login, err
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

func (r *UserRepositoryStruct) BlockUser(ctx context.Context, id int, blocked bool) error {
	query := `
		UPDATE users 
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
		SET last_login = $1
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
		SELECT id, email, login, Password_hash, Role_Id, Blocked, FullName, Class_Id, Created_at, last_login
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
		SELECT id, email, login, password_hash, role_id, blocked, fullName, class_id, created_at, last_login
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

		var classId sql.NullInt64
		err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Login,
			&user.PasswordHash,
			&user.RoleId,
			&user.Blocked,
			&user.FullName,
			&classId,
			&user.CreatedAt,
			&user.LastLogin,
		)

		if !classId.Valid {
			user.ClassId = nil
		} else {
			i := int(classId.Int64)
			user.ClassId = &i
		}

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

func (r *UserRepositoryStruct) UserExists(ctx context.Context, login string, password string) (int, error) {
	query := `
		SELECT id, password_hash 
		FROM users
		WHERE login = $1
	`

	var id int
	var hashpassword []byte
	err := r.db.QueryRowContext(ctx, query, login).Scan(&id, &hashpassword)
	if err != nil {
		return 0, err
	}

	if err = bcrypt.CompareHashAndPassword(hashpassword, []byte(password)); err != nil {
		return 0, errors.New("user or password is wrong")
	}

	return id, nil
}

func (r *UserRepositoryStruct) GetStudentsByClass(ctx context.Context, classId int) ([]model.User, error) {
	query := `
		SELECT id, email, login, Password_hash, Role_Id, Blocked, FullName, Class_Id, Created_at, last_login
		FROM users 
		WHERE class_id = $1 AND role_id = 1
	`

	rows, err := r.db.QueryContext(ctx, query, classId)
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

func (r *UserRepositoryStruct) GetTeachersBySchool(ctx context.Context, schoolId int) ([]model.User, error) {
	query := `
		SELECT users.id, email, login, password_hash, role_id, blocked, fullName, class_id, users.created_at, last_login
		FROM users 
		JOIN classes ON class_id = classes.id 
		WHERE role_id = 3
	`

	args := make([]interface{}, 0)

	if schoolId > 0 {
		query += "AND classes.school_id = $1"
		args = append(args, schoolId)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]model.User, 0)
	for rows.Next() {
		var lastLogin sql.NullTime
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
			&lastLogin,
		)

		if err != nil {
			return nil, err
		}

		if lastLogin.Valid {
			user.LastLogin = &lastLogin.Time
		} else {
			user.LastLogin = nil
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryStruct) GetTotalStudentBySchoolId(ctx context.Context, schoolId int) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM users JOIN classes ON users.class_id = classes.id
		WHERE classes.school_id = $1
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, schoolId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *UserRepositoryStruct) GetStudentProfileById(ctx context.Context, id int) (*model.StudentProfile, error) {
	query := `
		SELECT id, fullname, classes.name, schools.name, xp 
		FROM users 
		JOIN classes ON users.class_id = classes.id
		JOIN schools ON classes.school_id = schools.id
		WHERE id = $1
	`

	var profile model.StudentProfile
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&profile.ID,
		&profile.FullName,
		&profile.ClassName,
		&profile.SchoolName,
		&profile.XP,
	)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *UserRepositoryStruct) AddXP(ctx context.Context, studentId int, xp int) error {
	query := `
		UPDATE users
		SET xp = xp + $1
		WHERE id = $2;
	`

	res, err := r.db.ExecContext(ctx, query, xp, studentId)
	if err != nil {
		return err
	}

	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *StudentProgressRepositoryStruct) GetStudentXP(ctx context.Context, studentId int) (int, error) {
	query := `
		SELECT xp
		FROM users
		WHERE users.id = $1;
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, studentId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
