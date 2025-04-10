package utils

import (
	"encoding/json"
	"go-boilerplate/internal/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SetHeaderJson(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
}

func SendResponse(res http.ResponseWriter, msg string, status int, data interface{}) {
	SetHeaderJson(res)

	response := models.Response{
		Message: msg,
		Status: status,
		Data: data,
	}

	res.WriteHeader(status)
	json.NewEncoder(res).Encode(response)
}

func GenerateJWT(user *models.User, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"username": user.Username,
		"roles": user.Roles,
		"exp": jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	return token.SignedString([]byte(secret))
}