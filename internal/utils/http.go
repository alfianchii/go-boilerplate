package utils

import (
	"encoding/json"
	"go-boilerplate/internal/models"
	"net/http"
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