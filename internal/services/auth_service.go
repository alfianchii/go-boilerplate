package services

import (
	"context"
	"errors"
	"go-boilerplate/configs"
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/internal/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	GenerateJWT(ctx context.Context, creds models.LoginRequest, ipAddress string) (string, error)
}

type AuthService struct {
	userRepo repositories.UserRepositoryInterface
	sessionRepo repositories.SessionRepositoryInterface
}

func NewAuthService(userRepo repositories.UserRepositoryInterface, tokenRepo repositories.SessionRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		userRepo: userRepo,
		sessionRepo: tokenRepo,
	}
}

func (s *AuthService) GenerateJWT(ctx context.Context, creds models.LoginRequest, ipAddress string) (string, error) {
	user, err := s.userRepo.FindByUsernameWithRoles(ctx, creds.Username)
	if err != nil {
		return "", errors.New(err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	jwt, err := utils.GenerateJWT(user, configs.GetENV("JWT_SECRET"))
	if err != nil {
			return "", err
	}

	rowID, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("failed to generate UUID for session")
	}

	session := models.Session{
		RowID: rowID.String(),
		UserID: user.ID,
		Token: jwt.Token,
		ExpiresAt: jwt.ExpiresAt.Time,
		CreatedAt: time.Now().Local(),
		IPAddress: ipAddress,
		IsBlacklisted: false,
	}

	s.sessionRepo.StoreSession(ctx, session)

	return jwt.Token, nil
}