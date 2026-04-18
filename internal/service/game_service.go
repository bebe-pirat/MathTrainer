package service

import (
	"MathTrainer/internal"
	"MathTrainer/internal/generator"
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
	"math/rand/v2"
	"sort"
)

type weightedType struct {
	Type   model.EquationTypeWithOperands
	Weight float64
	Count  int
}

type GameService interface {
	GenerateAdaptiveEquationSet(ctx context.Context, sectionId int, studentId int) ([]model.Equation, error)
	CheckEquations(ctx context.Context, answers []model.Answer) ([]model.EquationFeedback, error)

	CreateAttempts(ctx context.Context, answers []model.Answer, studentId int) error
	CreateStudentLevelProgress(ctx context.Context, feedback []model.EquationFeedback, studentId int, sectionId int) error
}

type GameServiceStruct struct {
	equationRepo repository.EquationTypeRepository
	attemptRepo  repository.EquationAttemptsRepository
}

func NewGameServiceStruct(equationRepo repository.EquationTypeRepository, attemptRepo repository.EquationAttemptsRepository) *GameServiceStruct {
	return &GameServiceStruct{
		equationRepo: equationRepo,
		attemptRepo:  attemptRepo,
	}
}

func (s *GameServiceStruct) GenerateAdaptiveEquationSet(ctx context.Context, sectionId int, studentId int) ([]model.Equation, error) {
	typeStats, err := s.attemptRepo.GetStudentSectionStats(ctx, studentId, sectionId)
	if err != nil {
		return nil, err
	}

	types, err := s.equationRepo.GetEquationTypesBySection(ctx, sectionId)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(types); i++ {
		operands, err := s.equationRepo.GetOperandsByEquationType(ctx, types[i].Id)
		if err != nil {
			return nil, err
		}

		types[i].Operands = operands
	}

	weightedTypes, totalWeight := s.generateWeightedTypes(typeStats, types)

	equations, err := s.generateEquationsBasedOnWeightedTypes(weightedTypes, totalWeight)
	if err != nil {
		return nil, err
	}

	return equations, nil
}

func (s *GameServiceStruct) generateWeightedTypes(typeStats map[int]float32, types []model.EquationTypeWithOperands) ([]weightedType, float64) {
	weightedTypes := make([]weightedType, 0)

	weakLen := 0
	mediumLen := 0
	strongLen := 0
	newLen := 0

	for i := 0; i < len(types); i++ {
		accuracy, exists := typeStats[types[i].Id]
		if !exists {
			weightedTypes = append(weightedTypes, weightedType{Type: types[i],
				Weight: internal.NewWeight,
				Count:  0})
			newLen++
			continue
		}

		if accuracy < internal.WeakCategoryAccurary {
			weightedTypes = append(weightedTypes, weightedType{Type: types[i],
				Weight: internal.WeakWeight,
				Count:  0})
			weakLen++
		} else if accuracy < internal.MegiumCategoryAccuracy {
			weightedTypes = append(weightedTypes, weightedType{Type: types[i],
				Weight: internal.MediumWeight,
				Count:  0})
			mediumLen++
		} else {
			weightedTypes = append(weightedTypes, weightedType{Type: types[i],
				Weight: internal.StrongWeight,
				Count:  0})
			strongLen++
		}
	}

	totalWeight := 0.0
	totalWeight += float64(weakLen) * internal.WeakWeight
	totalWeight += float64(mediumLen) * internal.MediumWeight
	totalWeight += float64(strongLen) * internal.StrongWeight
	totalWeight += float64(newLen) * internal.NewWeight

	return weightedTypes, totalWeight
}

func (s *GameServiceStruct) generateEquationsBasedOnWeightedTypes(weightedTypes []weightedType, totalWeight float64) ([]model.Equation, error) {
	equations := make([]model.Equation, internal.CountEquationInSet)

	for i := 0; i < internal.CountEquationInSet; i++ {
		sort.Slice(weightedTypes, func(i, j int) bool {
			expectedI := float64(i+1) * weightedTypes[i].Weight / totalWeight
			expectedJ := float64(i+1) * weightedTypes[j].Weight / totalWeight
			ratioI := float64(weightedTypes[i].Count) / expectedI
			ratioJ := float64(weightedTypes[j].Count) / expectedJ
			return ratioI < ratioJ
		})

		candidates := weightedTypes
		if len(weightedTypes) > 3 {
			candidates = weightedTypes[:3]
		}

		selectedIdx := 0
		if len(candidates) > 1 {
			selectedIdx = rand.IntN(len(candidates))
		}

		selectedType := candidates[selectedIdx].Type

		eq, err := generator.GenerateEquation(selectedType)
		if err != nil {
			return nil, err
		}

		eq.Id = i

		equations = append(equations, eq)

		for i := range weightedTypes {
			if weightedTypes[i].Type.Id == selectedType.Id {
				weightedTypes[i].Count++
				break
			}
		}
	}

	return equations, nil
}
