package utils

import (
	"errors"
	"go-boilerplate/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID string   `json:"user_id"`
	Username string   `json:"username"`
	Roles  []models.Role `json:"roles"`
	jwt.RegisteredClaims
}

func GenerateJWT(user *models.User, secret string) (string, error) {
	claims := UserClaims{
		UserID: user.ID,
		Username: user.Username,
		Roles:  user.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.Username,
			Subject:   user.ID,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
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