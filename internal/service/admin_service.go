package service

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
	"errors"
	"log/slog"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	CreateSchool(ctx context.Context, name string, address string) error
	GetSchools(ctx context.Context) ([]model.School, error)
	GetTeachersBySchoolId(ctx context.Context, schoolId int) ([]model.User, error)

	CreateTeacher(ctx context.Context, fullName, login, email string, classId int) (*model.UserCredentials, error)
	ChangeBlockingUser(ctx context.Context, userId int, blocked bool) error
	GetAllUsers(ctx context.Context) ([]model.User, error)
}

type AdminServiceStruct struct {
	userRepo   repository.UserRepository
	schoolRepo repository.SchoolRepository
}

func NewAdminServiceStruct(userRepo repository.UserRepository, schoolRepo repository.SchoolRepository) *AdminServiceStruct {
	return &AdminServiceStruct{
		userRepo:   userRepo,
		schoolRepo: schoolRepo,
	}
}

func (s *AdminServiceStruct) CreateSchool(ctx context.Context, name string, address string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("bad request")
	}

	address = strings.TrimSpace(address)
	if address == "" {
		return errors.New("bad request")
	}

	school := model.School{
		Name:       name,
		Address:    address,
		Created_at: time.Now(),
	}

	return s.schoolRepo.CreateSchool(ctx, school)
}

func (s *AdminServiceStruct) GetSchools(ctx context.Context) ([]model.School, error) {
	return s.schoolRepo.GetAllSchools(ctx)
}

func (s *AdminServiceStruct) GetTeachersBySchoolId(ctx context.Context, schoolId int) ([]model.User, error) {
	return s.userRepo.GetTeachersBySchool(ctx, schoolId)
}

func (s *AdminServiceStruct) CreateTeacher(ctx context.Context, fullName, login, email string, classId int) (*model.UserCredentials, error) {
	slog.Info("parameters", "fullname", fullName, "email", email, "login", login, "class_id", classId)

	fullName = strings.TrimSpace(fullName)
	if fullName == "" {
		return nil, errors.New("bad request")
	}

	email = strings.TrimSpace(email)
	if email == "" {
		return nil, errors.New("bad request")
	}

	login = strings.TrimSpace(login)
	if login == "" {
		return nil, errors.New("bad request")
	}

	password, err := GenerateRandomPassword()
	if err != nil {
		return nil, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	teacher := model.User{
		Email:        email,
		Login:        login,
		FullName:     fullName,
		PasswordHash: passwordHash,
		RoleId:       3, // тут убрать эту херню, заменить на что-то адекватное, тупая я 2 роль - алмин, а 3 - учитель
		Blocked:      false,
		ClassId:      &classId,
		CreatedAt:    time.Now(),
	}

	login, err = s.userRepo.CreateUser(ctx, teacher)
	if err != nil {
		return nil, err
	}

	return &model.UserCredentials{Login: login, Password: password}, nil
}

func (s *AdminServiceStruct) ChangeBlockingUser(ctx context.Context, userId int, blocked bool) error {
	if userId <= 0 {
		return errors.New("bad request")
	}

	return s.userRepo.BlockUser(ctx, userId, blocked)
}

func (s *AdminServiceStruct) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.userRepo.GetAllUsers(ctx)
}

// func (s *UserServiceStruct) CreateUser(ctx context.Context, e model.User) (int, error) {
// 	e.PasswordHash = strings.TrimSpace(e.PasswordHash)
// 	if e.PasswordHash == "" {
// 		return 0, errors.New("passwords is required")
// 	}
// 	passwordHash, err := bcrypt.GenerateFromPassword([]byte(e.PasswordHash), bcrypt.DefaultCost)
// 	if err != nil {
// 		return 0, err
// 	}
// 	e.PasswordHash = string(passwordHash)

// 	e.Email = strings.TrimSpace(e.Email)
// 	if e.Email == "" {
// 		return 0, errors.New("email is required")
// 	}

// 	e.Login = strings.TrimSpace(e.Login)
// 	if e.Login == "" {
// 		return 0, errors.New("login is required")
// 	}

// 	e.FullName = strings.TrimSpace(e.FullName)
// 	if e.FullName == "" {
// 		return 0, errors.New("fullname is required")
// 	}

// 	return s.userRepo.CreateUser(ctx, e)
// }

// func (s *UserServiceStruct) UpdateUser(ctx context.Context, e model.User) (*model.User, error) {
// 	e.PasswordHash = strings.TrimSpace(e.PasswordHash)
// 	if e.PasswordHash == "" {
// 		return nil, errors.New("passwords is required")
// 	}
// 	passwordHash, err := bcrypt.GenerateFromPassword([]byte(e.PasswordHash), bcrypt.DefaultCost)
// 	if err != nil {
// 		return nil, err
// 	}
// 	e.PasswordHash = string(passwordHash)

// 	e.Email = strings.TrimSpace(e.Email)
// 	if e.Email == "" {
// 		return nil, errors.New("email is required")
// 	}

// 	e.Login = strings.TrimSpace(e.Login)
// 	if e.Login == "" {
// 		return nil, errors.New("login is required")
// 	}

// 	e.FullName = strings.TrimSpace(e.FullName)
// 	if e.FullName == "" {
// 		return nil, errors.New("fullname is required")
// 	}

// 	return s.userRepo.UpdateUser(ctx, e)
// }

// func (s *UserServiceStruct) DeleteUser(ctx context.Context, id int) error {
// 	if id <= 0 {
// 		return errors.New("bad request")
// 	}

// 	return s.userRepo.DeleteUser(ctx, id)
// }

// func (s *UserServiceStruct) GetAllUsers(ctx context.Context) ([]model.User, error) {
// 	return s.userRepo.GetAllUsers(ctx)
// }
