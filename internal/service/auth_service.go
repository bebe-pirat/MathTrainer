package service

import (
	"MathTrainer/internal/repository"
	"context"
	"errors"
	"time"
)

type AuthService interface {
	Login(ctx context.Context, login string, password string) (int, error)
	Logout(ctx context.Context, session_id int) error
	UpdateLastLoginUser(ctx context.Context, id int, lastLogin time.Time) error
}

type AuthServiceStruct struct {
	userRepo repository.UserRepository
}

func NewAuthServiceStruct(userRepo repository.UserRepository) *AuthServiceStruct {
	return &AuthServiceStruct{
		userRepo: userRepo,
	}
}

func (s *AuthServiceStruct) Login(ctx context.Context, login string, password string) (int, error)

// пока что оставим это
// разберись с JWT и созданием сессии + куки

func (s *AuthServiceStruct) Logout(ctx context.Context)

// подумать, что тут сделать
/* сделать 3 миграцию, где создаться таблица session
при logout надо удалять сессию из БД, из куков и что-то сделать с JWT(?)
тоже снести скорее всего
*/

func (s *AuthServiceStruct) UpdateLastLoginUser(ctx context.Context, id int, lastLogin time.Time) error {
	if id <= 0 {
		return errors.New("bad request")
	}

	return s.userRepo.UpdateLastLoginUser(ctx, id, lastLogin)
}
