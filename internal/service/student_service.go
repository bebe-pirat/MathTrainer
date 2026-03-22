package service

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
	"errors"
)

type StudentService interface {
	GetProfile(ctx context.Context, studentID int) (*model.StudentProfile, error)
	GetAchievements(ctx context.Context, studentID int) ([]model.AchievementOfStudent, error)
}

type StudentServiceStruct struct {
	studentRepo repository.UserRepository
	achRepo     repository.AchievementOfStudentRepository
}

func NewStudentServiceStruct(userRepo repository.UserRepository, achRepo repository.AchievementOfStudentRepository) *StudentServiceStruct {
	return &StudentServiceStruct{
		studentRepo: userRepo,
		achRepo:     achRepo,
	}
}

func (s *StudentServiceStruct) GetProfile(ctx context.Context, studentID int) (*model.StudentProfile, error) {
	if studentID <= 0 {
		return nil, errors.New("invalid id")
	}
	return s.studentRepo.GetStudentProfileById(ctx, studentID)
}

func (s *StudentServiceStruct) GetAchievements(ctx context.Context, studentID int) ([]model.AchievementOfStudent, error) {
	if studentID <= 0 {
		return nil, errors.New("invalid id")
	}

	return s.achRepo.GetAchievementOfStudentsByStudentId(ctx, studentID)
}
