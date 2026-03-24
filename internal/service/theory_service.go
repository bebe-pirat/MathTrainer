package service

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
	"errors"
)

type TheoryService interface {
	GetTheoryByEquationType(ctx context.Context, typeId int) (*model.Theory, error)
}

type TheoryServiceStruct struct {
	theoryRepo repository.TheoryRepository
}

func NewTheoryServiceStruct(theoryRepo repository.TheoryRepository) *TheoryServiceStruct {
	return &TheoryServiceStruct{
		theoryRepo: theoryRepo,
	}
}

func (s *TheoryServiceStruct) GetTheoryByEquationType(ctx context.Context, typeId int) (*model.Theory, error) {
	if typeId <= 0 {
		return nil, errors.New("bad request")
	}

	return s.theoryRepo.GetByEquationType(ctx, typeId)
}
