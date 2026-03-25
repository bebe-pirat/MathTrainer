package service

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/repository"
	"context"
	"errors"
	"time"
)

type AuthService interface {
	Login(ctx context.Context, login string, password string) (*model.SessionData, error)
	Logout(ctx context.Context, session_id int) error
	UpdateLastLoginUser(ctx context.Context, id int, lastLogin time.Time) error
}

type AuthServiceStruct struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
}

func NewAuthServiceStruct(userRepo repository.UserRepository, sessionRepo repository.SessionRepository) *AuthServiceStruct {
	return &AuthServiceStruct{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

// пока что так
func (s *AuthServiceStruct) Login(ctx context.Context, login string, password string) (*model.SessionData, error) {
	userId, err := s.userRepo.UserExists(ctx, login, password)
	if err != nil {
		return nil, err
	}

	sessionId, err := s.sessionRepo.CreateSession(ctx, userId, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, err
	}

	roleId, err := s.userRepo.GetRoleById(ctx, userId)
	if err != nil {
		return nil, err
	}

	sessionData := model.SessionData{
		SessionID: sessionId,
		UserID:    userId,
		Role:      roleId,
	}

	return &sessionData, err
}

func (s *AuthServiceStruct) Logout(ctx context.Context, sessionId int) error {
	if sessionId <= 0 {
		return errors.New("require session id")
	}

	return s.sessionRepo.DeleteSession(ctx, sessionId)
}

func (s *AuthServiceStruct) UpdateLastLoginUser(ctx context.Context, id int, lastLogin time.Time) error {
	if id <= 0 {
		return errors.New("bad request")
	}

	return s.userRepo.UpdateLastLoginUser(ctx, id, lastLogin)
}
