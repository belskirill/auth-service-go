package auth

import (
	"context"
	"proto-auth/internal/domain/auth"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	logger       *zap.Logger
	userSaver    UserSaver
	UserProvider UserProvider
}

func NewService(logger *zap.Logger, userSaver UserSaver, userProvider UserProvider) *Service {
	return &Service{
		logger:       logger,
		userSaver:    userSaver,
		UserProvider: userProvider,
	}
}

func (auth *Service) Login(ctx context.Context, login auth.Login) (string, error) {
	panic("implement me")
}

func (auth *Service) Register(ctx context.Context, register auth.Register) (int64, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	id, err := auth.userSaver.SaveUser(ctx, register.Email, string(hashPass))
	if err != nil {
		return 0, err
	}

	return id, nil

}
