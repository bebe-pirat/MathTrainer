package service

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type TeacherService interface {
	GetClassByTeacherId(ctx context.Context, teacherId int) (int, error)
	GetClassStudents(ctx context.Context, classId int) ([]model.User, error)
	GetStudentAttempts(ctx context.Context, studentId int, equation_type_id int) ([]model.AttemptForTeacher, error)
	GetEquationTypesByStudentId(ctx context.Context, studetnId int) ([]model.ShortEquationType, error) 

	CreateStudent(ctx context.Context, classId int, fullName, email, login string) (*model.UserCredentials, error)
	UpdateStudent(ctx context.Context, studentId int, fullName, email string) error
	DeleteStudent(ctx context.Context, studentId int) error
}

type TeacherServiceStruct struct {
	userRepo    repository.UserRepository
	attemptRepo repository.EquationAttemptsRepository
	equationTypeRepo repository.EquationTypeRepository
}

func NewTeacherServiceStruct(userRepo repository.UserRepository, attemptRepo repository.EquationAttemptsRepository, equationTypeRepo repository.EquationTypeRepository) *TeacherServiceStruct {
	return &TeacherServiceStruct{
		userRepo:    userRepo,
		attemptRepo: attemptRepo,
		equationTypeRepo: equationTypeRepo,
	}
}

func (s *TeacherServiceStruct) GetClassByTeacherId(ctx context.Context, teacherId int) (int, error) {
	if teacherId <= 0 {
		return 0, errors.New("invalid id")
	}

	return s.userRepo.GetClassByUserId(ctx, teacherId)
}

func (s *TeacherServiceStruct) GetStudentAttempts(ctx context.Context, studentId int, equation_type_id int) ([]model.AttemptForTeacher, error) {
	if studentId <= 0 {
		return nil, errors.New("invalid student id")
	}

	attempts, err := s.attemptRepo.GetStudentAttempts(ctx, studentId, equation_type_id)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return attempts, nil
}

func (s *TeacherServiceStruct) GetClassStudents(ctx context.Context, classId int) ([]model.User, error) {
	if classId <= 0 {
		return nil, errors.New("invalid id")
	}

	return s.userRepo.GetStudentsByClass(ctx, classId)	
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

	student := model.User{
		Email:        email,
		Login:        login,
		FullName:     fullName,
		PasswordHash: passwordHash,
		RoleId:       1, // тут убрать эту херню, заменить на что-то адекватное
		Blocked:      false,
		ClassId:      &classId,
		CreatedAt:    time.Now(),
	}

	login, err = s.userRepo.CreateUser(ctx, student)
	if err != nil {
		return nil, err
	}

	return &model.UserCredentials{Login: login, Password: password}, nil
}

func (s *TeacherServiceStruct) GetEquationTypesByStudentId(ctx context.Context, studentId int) ([]model.ShortEquationType, error)  {
	if studentId <= 0 {
		return nil, errors.New("bad requst")
	}

	return s.equationTypeRepo.GetEquationTypesByStudentId(ctx, studentId)
}

// TODO: метод не раюотает)))) исправить 
func (s *TeacherServiceStruct) UpdateStudent(ctx context.Context, studentId int, fullName, email string) error { 
	fullName = strings.TrimSpace(fullName)
	if fullName == "" {
		return errors.New("bad request")
	}

	email = strings.TrimSpace(email)
	if email == "" {
		return errors.New("bad request")
	}

	student := model.User{
		Email:    email,
		FullName: fullName,
	}
	_, err := s.userRepo.UpdateUser(ctx, student)

	return err
}

func (s *TeacherServiceStruct) DeleteStudent(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("bad request")
	}

	return s.userRepo.DeleteUser(ctx, id)
}
