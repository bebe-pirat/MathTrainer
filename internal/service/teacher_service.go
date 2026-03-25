package service

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	// "strings"
)

type TeacherService interface {
	GetClassByTeacherId(ctx context.Context, teacherId int) (int, error)
	GetClassStudents(ctx context.Context, classId int) ([]model.StudentShortStats, error)

	CreateStudent(ctx context.Context, classId int, fullName, email, login string) (*model.UserCredentials, error)
	UpdateStudent(ctx context.Context, studentId int, fullName, email string) error
	DeleteStudent(ctx context.Context, studentId int) error
}

type TeacherServiceStruct struct {
	userRepo    repository.UserRepository
	attemptRepo repository.EquationAttemptsRepository
}

func NewTeacherServiceStruct(userRepo repository.UserRepository, attemptRepo repository.EquationAttemptsRepository) *TeacherServiceStruct {
	return &TeacherServiceStruct{
		userRepo:    userRepo,
		attemptRepo: attemptRepo,
	}
}

func (s *TeacherServiceStruct) GetClassByUserId(ctx context.Context, teacherId int) (int, error) {
	if teacherId <=0 {
		return 0, errors.New("invalid id")
	}

	return s.userRepo.GetClassByUserId(ctx, teacherId)
}

func (s *TeacherServiceStruct) GetClassStudents(ctx context.Context, classId int) ([]model.StudentShortStats, error) {
	if classId <= 0 {
		return nil, errors.New("invalid id")
	}

	return s.attemptRepo.GetStudentsShortStatsByClassId(ctx, classId)
}

func (s *TeacherServiceStruct) CreateStudent(ctx context.Context, classId int, fullName, email, login string) (*model.UserCredentials, error) {
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
		RoleId:       1, // тут убрать эту херню, заменить на что-то адекватное
		Blocked:      false,
		ClassId:      classId,
		CreatedAt:    time.Now(),
	}

	login, err = s.userRepo.CreateUser(ctx, teacher)
	if err != nil {
		return nil, err
	}

	return &model.UserCredentials{Login: login, Password: password}, nil
}

func (s *AdminServiceStruct) UpdateStudent(ctx context.Context, studentId int, fullName, email string) (*model.User, error) { // метод не раотает))))
	fullName = strings.TrimSpace(fullName)
	if fullName == "" {
		return nil, errors.New("bad request")
	}

	email = strings.TrimSpace(email)
	if email == "" {
		return nil, errors.New("bad request")
	}

	student := model.User{
		Email:    email,
		FullName: fullName,
	}

	return s.userRepo.UpdateUser(ctx, student)
}

func (s *TeacherServiceStruct) DeleteUser(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("bad request")
	}

	return s.userRepo.DeleteUser(ctx, id)
}
