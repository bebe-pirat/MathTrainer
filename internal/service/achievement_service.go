package service

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
	"errors"
	"time"
)

type AchievementService interface {
	CheckAchievements(ctx context.Context, studentId int) ([]model.AchievementOfStudent, error)
	GiveAchievement(ctx context.Context, studentId int, achievementId int) error
}

type AchievementServiceStruct struct {
	achRepo repository.AchievementOfStudentRepository
}

func NewAchievementServiceStruct(achRepo repository.AchievementOfStudentRepository) *AchievementServiceStruct {
	return &AchievementServiceStruct{
		achRepo: achRepo,
	}
}

func (s *AchievementServiceStruct) CheckAchievements(ctx context.Context, studentId int) ([]model.AchievementOfStudent, error) {
	if studentId <= 0 {
		return nil, errors.New("bad request")
	}

	return s.achRepo.GetAchievementOfStudentsByStudentId(ctx, studentId)
}

func (s *AchievementServiceStruct) GiveAchievement(ctx context.Context, studentId int, achievementId int) error {
	if studentId <= 0 {
		return errors.New("bad request")
	}

	if achievementId <= 0 {
		return errors.New("bad request")
	}

	ach := model.AchievementOfStudent{
		StudentId:     studentId,
		AchievementId: achievementId,
		GotAt:         time.Now(),
	}

	return s.achRepo.CreateAchieveOfStud(ctx, ach)
}
