package service

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
	"database/sql"
)

type StatsService interface {
	GetSchoolStats(ctx context.Context, schoolId int) (*model.SchoolStats, error)
	GetClassStats(ctx context.Context, classId int) (*model.ClassStats, error)
	GetStudentStats(ctx context.Context, studentId int) (*model.StudentStats, error)
}

type StatsServiceStruct struct {
	classRepo    repository.ClassRepository
	schoolRepo   repository.SchoolRepository
	studentRepo  repository.UserRepository
	attemptRepo  repository.EquationAttemptsRepository
	progressRepo repository.StudentProgressRepository
	achievRepo   repository.AchievementOfStudentRepository
}

func NewStatStatsServiceStruct(classRepo repository.ClassRepository,
	schoolRepo repository.SchoolRepository,
	studentRepo repository.UserRepository,
	attemptsRepo repository.EquationAttemptsRepository,
	progressRepo repository.StudentProgressRepository,
	achievRepo repository.AchievementOfStudentRepository) *StatsServiceStruct {
	return &StatsServiceStruct{
		classRepo:    classRepo,
		schoolRepo:   schoolRepo,
		studentRepo:  studentRepo,
		attemptRepo:  attemptsRepo,
		progressRepo: progressRepo,
		achievRepo:   achievRepo,
	}
}

func (s *StatsServiceStruct) GetSchoolStats(ctx context.Context, schoolId int) (*model.SchoolStats, error) {
	countStudents, err := s.studentRepo.GetTotalStudentBySchoolId(ctx, schoolId)
	if err != nil {
		return nil, err
	}

	classes, err := s.classRepo.GetClassesBySchoolId(ctx, schoolId)
	if err != nil {
		return nil, err
	}

	// totalCount, err := s.attemptRepo.GetTotalAttemptsBySchoolId(ctx, schoolId)
	// if err != nil {
	// 	return nil, err
	// }

	// wrongCount, err := s.attemptRepo.GetWrongAnswersBySchoolId(ctx, schoolId)
	// if err != nil {
	// 	return nil, err
	// }

	// accuracy := float32(totalCount) / float32(wrongCount) * 100.0

	// equationTypes, err := s.attemptRepo.GetEquationTypeAccuracyBySchoolId(ctx, schoolId)
	// if err != nil {
	// 	return nil, err
	// }

	// classStats, err := s.attemptRepo.GetClassesAccuracyBySchoolId(ctx, schoolId)
	// if err != nil {
	// 	return nil, err
	// }

	// return &model.SchoolStats{
	// 	StudentsCount:       countStudents,
	// 	ClassesCount:        len(classes),
	// 	TotalEquationSolved: totalCount,
	// 	WrongAnswers:        wrongCount,
	// 	Accuracy:            accuracy,
	// 	EquationTypes:       equationTypes,
	// 	Classes:             classStats,
	// }, nil

	return &model.SchoolStats{
		StudentsCount:       countStudents,
		ClassesCount:        len(classes),
		TotalEquationSolved: 0,
		WrongAnswers:        0,
		Accuracy:            0,
		EquationTypes:       nil,
		Classes:             nil,
	}, nil
}

func (s *StatsServiceStruct) GetClassStats(ctx context.Context, classId int) (*model.ClassStats, error) {
	students, err := s.attemptRepo.GetStudentsShortStatsByClassId(ctx, classId)
	if err != nil {
		return nil, err
	}

	// TODO: удалить, сделано исключительно для отладки
	// slog.Info("student short stats", students)

	studentCount := len(students)

	totalCount, err := s.attemptRepo.GetTotalAttemptsByClassId(ctx, classId)
	if err == sql.ErrNoRows {
		totalCount = 0
	} else if err != nil {
		return nil, err
	}

	wrongCount, err := s.attemptRepo.GetWrongAnswersByClassId(ctx, classId)
	if err == sql.ErrNoRows {
		wrongCount = 0
	} else if err != nil {
		return nil, err
	}

	correctCount := totalCount - wrongCount
	accuracy := float32(0.0)
	if totalCount != 0 {
		accuracy = float32(correctCount) / float32(totalCount) * 100.0
	}

	equationTypes, err := s.attemptRepo.GetEquationTypeAccuracyByClassId(ctx, classId)
	if err != nil {
		return nil, err
	}

	return &model.ClassStats{
		StudentsCount:  studentCount,
		TotalAttempts:  totalCount,
		WrongAnswers:   wrongCount,
		CorrectAnswers: correctCount,
		Accuracy:       accuracy,
		EquationTypes:  equationTypes,
		Students:       students,
	}, nil
}

func (s *StatsServiceStruct) GetStudentStats(ctx context.Context, studentId int) (*model.StudentStats, error) {
	totalCount, err := s.attemptRepo.GetTotalCountAttempts(ctx, studentId)
	if err != nil {
		return nil, err
	}

	wrongCount, err := s.attemptRepo.GetCountErrorAttempts(ctx, studentId)
	if err != nil {
		return nil, err
	}

	correctCount := totalCount - wrongCount
	accuracy := float32(correctCount) / float32(totalCount) * 100.0

	// TODO: удали, нужно было для отладки
	// slog.Info("accuracy info", "accuracy", accuracy, "correct_count", correctCount, "total_count", totalCount, "wrong_count", wrongCount)

	complitedLevels, err := s.progressRepo.GetCountComplitedLevels(ctx, studentId)
	if err != nil {
		return nil, err
	}

	starsCount, err := s.progressRepo.GetTotalStars(ctx, studentId)
	if err != nil {
		return nil, err
	}

	xp, err := s.studentRepo.GetStudentXP(ctx, studentId)
	if err != nil {
		return nil, err
	}

	achievements, err := s.achievRepo.GetAchievementOfStudentsByStudentId(ctx, studentId)
	if err != nil {
		return nil, err
	}

	equationTypes, err := s.attemptRepo.GetExtendedEquationTypeStats(ctx, studentId)
	if err != nil {
		return nil, err
	}

	// TODO: удали, нужно для отладки
	// slog.Info("equation_types", "equation type name", equationTypes[0].Type, "accuracy", equationTypes[0].Accuracy)

	weakTopics := make([]string, 0)
	for _, value := range equationTypes {
		if value.Accuracy < 50.0 {
			weakTopics = append(weakTopics, value.Type)
		}
	}

	return &model.StudentStats{
		TotalAttempts:   totalCount,
		CorrectAnswers:  correctCount,
		WrongAnswers:    wrongCount,
		Accuracy:        accuracy,
		LevelsCompleted: complitedLevels,
		StarsTotal:      starsCount,
		XP:              xp,
		EquationTypes:   equationTypes,
		Achievements:    achievements,
		WeakTopics:      weakTopics,
	}, nil
}
