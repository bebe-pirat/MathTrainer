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
	GetStudentLevelsMap(ctx context.Context, studentId int) (*model.LevelsMap, error)
}

type StudentServiceStruct struct {
	studentRepo repository.UserRepository
	achRepo     repository.AchievementOfStudentRepository
	sectionRepo repository.SectionRepository
}

func NewStudentServiceStruct(userRepo repository.UserRepository, achRepo repository.AchievementOfStudentRepository, sectionRepo repository.SectionRepository) *StudentServiceStruct {
	return &StudentServiceStruct{
		studentRepo: userRepo,
		achRepo:     achRepo,
		sectionRepo: sectionRepo,
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

func (s *StudentServiceStruct) GetStudentLevelsMap(ctx context.Context, studentId int) (*model.LevelsMap, error) {
	if studentId <= 0 {
		return nil, errors.New("invalid id")
	}

	position, err := s.sectionRepo.GetStudentPosition(ctx, studentId)
	if err != nil {
		return nil, err
	}

	class, err := s.studentRepo.GetGradeByStudentId(ctx, studentId)
	if err != nil {
		return nil, err
	}

	sections, err := s.sectionRepo.GetSectionsByClass(ctx, class, position.SectionId)
	if err != nil {
		return nil, err
	}

	levelsMap := &model.LevelsMap{
		Sections: sections,
		Position: *position,
	}

	return levelsMap, nil
}
