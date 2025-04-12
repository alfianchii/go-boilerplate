package utils

import (
	"encoding/json"
	"go-boilerplate/internal/models"
	"net/http"
	"strings"
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

func GetClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}

	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	ip := r.RemoteAddr
	ip = strings.Split(ip, ":")[0]
	return ip
}