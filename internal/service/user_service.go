package service

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, e model.User) (int, error)

	UpdateUser(ctx context.Context, e model.User) (*model.User, error)
	BlockUser(ctx context.Context, id int, blocked bool) error
	UpdateLastLoginUser(ctx context.Context, id int, lastLogin time.Time) error

	DeleteUser(ctx context.Context, id int) error

	GetAllUsers(ctx context.Context) ([]model.User, error)

	Login(ctx context.Context, login string, password string) error
	Logout(ctx context.Context) // подумать, что тут сделать
}

type UserServiceStruct struct {
	userRepo repository.UserRepository
}

func NewUserServiceStruct(userRepo repository.UserRepository) *UserServiceStruct {
	return &UserServiceStruct{
		userRepo: userRepo,
	}
}

func (s *UserServiceStruct) CreateUser(ctx context.Context, e model.User) (int, error) {
	e.PasswordHash = strings.TrimSpace(e.PasswordHash)
	if e.PasswordHash == "" {
		return 0, errors.New("passwords is required")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(e.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	e.PasswordHash = string(passwordHash)

	e.Email = strings.TrimSpace(e.Email)
	if e.Email == "" {
		return 0, errors.New("email is required")
	}

	e.Login = strings.TrimSpace(e.Login)
	if e.Login == "" {
		return 0, errors.New("login is required")
	}

	e.FullName = strings.TrimSpace(e.FullName)
	if e.FullName == "" {
		return 0, errors.New("fullname is required")
	}

	return s.userRepo.CreateUser(ctx, e)
}

func (s *UserServiceStruct) UpdateUser(ctx context.Context, e model.User) (*model.User, error) {
	e.PasswordHash = strings.TrimSpace(e.PasswordHash)
	if e.PasswordHash == "" {
		return nil, errors.New("passwords is required")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(e.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	e.PasswordHash = string(passwordHash)

	e.Email = strings.TrimSpace(e.Email)
	if e.Email == "" {
		return nil, errors.New("email is required")
	}

	e.Login = strings.TrimSpace(e.Login)
	if e.Login == "" {
		return nil, errors.New("login is required")
	}

	e.FullName = strings.TrimSpace(e.FullName)
	if e.FullName == "" {
		return nil, errors.New("fullname is required")
	}

	return s.userRepo.UpdateUser(ctx, e)
}

func (s *UserServiceStruct) BlockUser(ctx context.Context, id int, blocked bool) error {
	if id <= 0 {
		return errors.New("bad request")
	}

	return s.userRepo.BlockUser(ctx, id, blocked)
}

func (s *UserServiceStruct) UpdateLastLoginUser(ctx context.Context, id int, lastLogin time.Time) error {
	if id <= 0 {
		return errors.New("bad request")
	}

	return s.userRepo.UpdateLastLoginUser(ctx, id, lastLogin)
}

func (s *UserServiceStruct) DeleteUser(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("bad request")
	}

	return s.userRepo.DeleteUser(ctx, id)
}

func (s *UserServiceStruct) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.userRepo.GetAllUsers(ctx)
}

func (s *UserServiceStruct) Login(ctx context.Context, login string, password string) error

// пока что оставим это
// разберись с JWT и созданием сессии + куки

func (s *UserServiceStruct) Logout(ctx context.Context)

// подумать, что тут сделать
/* сделать 3 миграцию, где создаться таблица session
при logout надо удалять сессию из БД, из куков и что-то сделать с JWT(?)
тоже снести скорее всего
*/
