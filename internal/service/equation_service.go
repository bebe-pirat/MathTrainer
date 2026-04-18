package service

// import (
// 	"MathTrainer/internal/model"
// 	"MathTrainer/internal/repository"
// 	"context"
// 	"errors"
// )

// type EquationService interface {
// 	GenerateRandomEquation(ctx context.Context, studentID int, levelID int) (*model.Equation, error)
// 	GenerateRandomEquationByEquationTypeId(ctx context.Context, level_id, equation_type_id int) (*model.Equation, error)
// 	GenerateEquationById(ctx context.Context, id int) (*model.Equation, error)
// }

// type EquationServiceStruct struct {
// 	equationRepo repository.EquationRepository
// }

// func NewEquationServiceStruct(equationRepo repository.EquationRepository) *EquationServiceStruct {
// 	return &EquationServiceStruct{
// 		equationRepo: equationRepo,
// 	}
// }

// func (s *EquationServiceStruct) GenerateRandomEquation(ctx context.Context, studentID int, levelID int) (*model.Equation, error) {
// 	if levelID <= 0 {
// 		return nil, errors.New("invalid id")
// 	}

// 	return s.equationRepo.GetRandomEquation(ctx, levelID)
// }

// func (s *EquationServiceStruct) GenerateRandomEquationByEquationTypeId(ctx context.Context, levelId, equationTypeId int) (*model.Equation, error) {
// 	if levelId <= 0 {
// 		return nil, errors.New("invalid id")
// 	}

// 	if equationTypeId <= 0 {
// 		return nil, errors.New("invalid id")
// 	}

// 	return s.equationRepo.GetRandomEquationByEquationTypeId(ctx, levelId, equationTypeId)
// }
// func (s *EquationServiceStruct) GenerateEquationById(ctx context.Context, id int) (*model.Equation, error) {
// 	if id <= 0 {
// 		return nil, errors.New("invalid id")
// 	}

// 	return s.equationRepo.GetEquationById(ctx, id)
// }
