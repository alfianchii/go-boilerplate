package services

import (
	"context"
	"errors"
	"go-boilerplate/configs"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	GenerateJWT(ctx context.Context, creds models.LoginRequest) (string, error)
}

type AuthService struct {
	userRepo repositories.UserRepositoryInterface
}

func NewAuthService(userRepo repositories.UserRepositoryInterface) AuthServiceInterface {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) GenerateJWT(ctx context.Context, creds models.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByUsernameWithRoles(ctx, creds.Username)
	if err != nil {
		return "", errors.New(err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWT(user, configs.GetENV("JWT_SECRET"))
	if err != nil {
			return "", errors.New("failed to generate token")
	}
	
	return token, nil
}