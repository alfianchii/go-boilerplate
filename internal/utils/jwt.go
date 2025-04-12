package utils

import (
	"errors"
	"go-boilerplate/configs"
	"go-boilerplate/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID int `json:"user_id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Roles []models.Role `json:"roles"`
	jwt.RegisteredClaims
}

type GeneratedJWT struct {
	Token string `json:"token"`
	ExpiresAt *jwt.NumericDate `json:"expires_at"`
}

func GenerateJWT(user *models.User, secret string) (GeneratedJWT, error) {
	tokenExp := jwt.NewNumericDate(time.Now().Add(configs.TokenDuration))
	
	claims := UserClaims{
		UserID: user.ID,
		Name: user.Name,
		Username: user.Username,
		Roles:  user.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.Name,
			Subject:   user.Username,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: tokenExp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return GeneratedJWT{}, errors.New("failed to sign token")
	}

	return GeneratedJWT{
		Token: signedToken,
		ExpiresAt: tokenExp,
	}, nil
}

func ValidateJWT(tokenString string, secret string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, jwt.ErrTokenInvalidId
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func GetBearerToken(authHeader string) (string, error) {
	unique := len("Bearer ")
	
	if authHeader == "" {
		return "", errors.New("unauthorized; authorization header is empty")
	}

	if len(authHeader) < unique || authHeader[:unique] != "Bearer " {
		return "", errors.New("unauthorized; uthorization header is not a Bearer token")
	}

	return authHeader[unique:], nil
}