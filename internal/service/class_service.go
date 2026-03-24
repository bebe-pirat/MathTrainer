package service

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
	"errors"
	"strings"
	"time"
)

type ClassService interface {
	GetClassesBySchool(ctx context.Context, schoolId int) ([]model.Class, error)

	CreateClass(ctx context.Context, name string, grade int, schoolId int) (int, error)
	UpdateClass(ctx context.Context, classId int, name string) error
	DeleteClass(ctx context.Context, classId int) error
}

type ClassServiceStruct struct {
	classRepo repository.ClassRepository
}

func NewClassServiceStruct(classRepo repository.ClassRepository) *ClassServiceStruct {
	return &ClassServiceStruct{
		classRepo: classRepo,
	}
}

func (s *ClassServiceStruct) GetClassesBySchool(ctx context.Context, schoolId int) ([]model.Class, error) {
	return s.classRepo.GetClassesBySchoolId(ctx, schoolId)
}

func (s *ClassServiceStruct) CreateClass(ctx context.Context, name string, grade int, schoolId int) (int, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return 0, errors.New("bad request")
	}

	if grade <= 0 || grade > 4 {
		return 0, errors.New("grade is over borders")
	}

	if schoolId <= 0 {
		return 0, errors.New("bad request")
	}

	class := model.Class{
		Name:      name,
		Grade:     grade,
		SchoolId:  schoolId,
		CreatedAt: time.Now(),
	}

	return s.classRepo.CreateClass(ctx, class)
}

func (s *ClassServiceStruct) UpdateClass(ctx context.Context, classId int, name string) (*model.Class, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("bad request")
	}

	class := model.Class{
		Id:   classId,
		Name: name,
	}

	return s.classRepo.UpdateClass(ctx, class)
}

func (s *ClassServiceStruct) DeleteClass(ctx context.Context, classId int) error {
	if classId <= 0 {
		return errors.New("bad request")
	}

	return s.classRepo.DeleteClass(ctx, classId)
}	
