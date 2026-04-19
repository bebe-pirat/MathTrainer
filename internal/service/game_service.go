package service

import (
	"MathTrainer/internal"
	"MathTrainer/internal/generator"
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
	"errors"
	"math/rand/v2"
	"sort"
	"time"
)

type weightedType struct {
	Type   model.EquationTypeWithOperands
	Weight float64
	Count  int
}

type GameService interface {
	GenerateAdaptiveEquationSet(ctx context.Context, sectionId int, studentId int) ([]model.Equation, error)                                     // получение сета уравнений
	CheckEquations(ctx context.Context, answers []model.Answer) ([]model.EquationFeedback, error)                                                // проверка
	FinishLevel(ctx context.Context, feedback []model.EquationFeedback, sectionId int, levelOrder int, studentId int) (*model.StarsAndXP, error) // завершение xp and stars

	CreateAttempts(ctx context.Context, answers []model.Answer, studentId int) error
	CreateStudentLevelProgress(ctx context.Context, countStars int, studentId int, sectionId int, levelOrder int) error
}

type GameServiceStruct struct {
	equationRepo repository.EquationTypeRepository
	attemptRepo  repository.EquationAttemptsRepository
	progressRepo repository.StudentProgressRepository
	userRepo     repository.UserRepository
}

func NewGameServiceStruct(equationRepo repository.EquationTypeRepository, attemptRepo repository.EquationAttemptsRepository, progressRepo repository.StudentProgressRepository, userRepo repository.UserRepository) *GameServiceStruct {
	return &GameServiceStruct{
		equationRepo: equationRepo,
		attemptRepo:  attemptRepo,
		progressRepo: progressRepo,
		userRepo:     userRepo,
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

func (s *GameServiceStruct) CreateAttempts(ctx context.Context, answers []model.Answer, studentId int) error {
	for i := 0; i < len(answers); i++ {
		attempt := model.Attempt{
			StudentId:      studentId,
			EquationTypeId: answers[i].EquationTypeId,
			GivenAnswer:    answers[i].UserAnswer,
			CorrectAnswer:  answers[i].CorrectAnswer,
			EquationText:   answers[i].Text,
			AnsweredAt:     time.Now(),
		}

		err := s.attemptRepo.CreateAttempt(ctx, attempt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *GameServiceStruct) CreateStudentLevelProgress(ctx context.Context, countStars int, studentId int, sectionId int, levelOrder int) error {
	if countStars < internal.MinStars || countStars > internal.MaxStars {
		return errors.New("bad request")
	}

	return s.progressRepo.CreateStudentProgressLevel(ctx, model.StudentProgress{
		StudentId:   studentId,
		SectionId:   sectionId,
		LevelOrder:  levelOrder,
		CountStarts: countStars,
		FinishedAt:  time.Now(),
	})
}

func (s *GameServiceStruct) CheckEquations(ctx context.Context, answers []model.Answer) ([]model.EquationFeedback, error) {
	feedback := make([]model.EquationFeedback, len(answers))

	for i := 0; i < len(answers); i++ {
		feedback[i] = model.EquationFeedback{
			EquationId:    answers[i].EquationId,
			IsCorrect:     answers[i].CorrectAnswer == answers[i].UserAnswer,
			CorrectAnswer: answers[i].CorrectAnswer,
			Feedback:      answers[i].Text,
		}
	}

	return feedback, nil
}

func (s *GameServiceStruct) FinishLevel(ctx context.Context, feedback []model.EquationFeedback, sectionId int, levelOrder int, studentId int) (*model.StarsAndXP, error) {
	correctCount, xp := 0, 0

	for _, item := range feedback {
		if item.IsCorrect {
			xp += internal.RightAnswerRegularExpressionXP
			correctCount++
		}

		if !item.IsCorrect {
			xp += internal.WrongAnswer
		}
	}

	if xp%internal.CountEquationInSet == internal.RightAnswerRegularExpressionXP {
		xp += internal.LevelWithoutMistakes
	}

	err := s.userRepo.AddXP(ctx, studentId, xp)
	if err != nil {
		return nil, err
	}

	stars := 0

	accuracy := float64(correctCount) / float64(internal.CountEquationInSet)

	if accuracy >= internal.OneStarsPercent && accuracy < internal.TwoStarsPercent {
		stars = 1
	} else if accuracy >= internal.TwoStarsPercent && accuracy < internal.ThreeStarsPercent {
		stars = 2
	} else if accuracy >= internal.ThreeStarsPercent && accuracy < 1 {
		stars = 3
	}

	err = s.CreateStudentLevelProgress(ctx, stars, studentId, sectionId, levelOrder)
	if err != nil {
		return nil, err
	}

	return &model.StarsAndXP{
		SectionId:  sectionId,
		LevelOrder: levelOrder,
		Stars:      stars,
		CommonXP:   xp,
	}, nil
}
