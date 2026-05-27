package service

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
)

type DirectorService interface {
	GetSchoolIdByDirectorId(ctx context.Context, directorId int) (int, error)
}

type directorService struct {
	schoolRepo repository.SchoolRepository
}

func NewDirectorService(schoolRepo repository.SchoolRepository) DirectorService {
	return &directorService{
		schoolRepo: schoolRepo,
	}
}

func (s *directorService) GetSchoolIdByDirectorId(ctx context.Context, directorId int) (int, error) {
	if directorId < 0 {
		return 0, model.BadRequest("wrong director id")
	}

	return s.schoolRepo.GetSchoolIdByDirectorId(ctx, directorId)
}
