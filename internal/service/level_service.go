package service

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
	"database/sql"
	"errors"
)

type LevelService interface {
	GetLevels(ctx context.Context) ([]model.Level, error)
	GetLevelById(ctx context.Context, id int) (*model.Level, error)
	GetTestLevel(ctx context.Context) (*model.Level, error)

	// возвращает id студент прогресс который начал этот уровень
	StartLevel(ctx context.Context, studentId, levelId int) error
	FinishLevel(ctx context.Context, e model.StudentProgress) (*model.StudentProgress, error)
}

type LevelServiceStruct struct {
	levelRepo    repository.LevelRepository
	progressRepo repository.StudentProgressRepository
}

func NewLevelServiceStruct(levelRepo repository.LevelRepository, progressRepo repository.StudentProgressRepository) *LevelServiceStruct {
	return &LevelServiceStruct{
		levelRepo:    levelRepo,
		progressRepo: progressRepo,
	}
}

func (s *LevelServiceStruct) GetLevels(ctx context.Context) ([]model.Level, error) {
	return s.levelRepo.GetAllLevels(ctx)
}

func (s *LevelServiceStruct) GetLevelById(ctx context.Context, id int) (*model.Level, error) {
	if id < 0 {
		return nil, errors.New("id is required")
	}

	return s.levelRepo.GetById(ctx, id)
}

func (s *LevelServiceStruct) GetTestLevel(ctx context.Context) (*model.Level, error) {
	return s.levelRepo.GetTestLevel(ctx)
}

func (s *LevelServiceStruct) StartLevel(ctx context.Context, studentId, levelId int) error {
	if studentId <= 0 {
		return model.NewBadRequestError("id of student is required")
	}

	if levelId <= 0 {
		return model.NewBadRequestError("id of student is required")
	}

	err := s.progressRepo.StartLevel(ctx, studentId, levelId)
	if err == sql.ErrNoRows {
		return model.NewNotFoundError("level is not found")
	}

	if err != nil {
		return model.NewInternalServerError(err.Error())
	}

	return nil
}

func (s *LevelServiceStruct) FinishLevel(ctx context.Context, e model.StudentProgress) (*model.StudentProgress, error) {
	if e.Id <= 0 {
		return nil, errors.New("id is required")
	}

	if e.CountStarts > 3 || e.CountStarts < 0 {
		return nil, errors.New("count of stars should be between 0 and 3")
	}

	return s.progressRepo.FinishLevel(ctx, e)
}
