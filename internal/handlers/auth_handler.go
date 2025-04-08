package handlers

import (
	"go-boilerplate/internal/models"
	"go-boilerplate/internal/services"
	"go-boilerplate/internal/utils"
	"net/http"
)

type AuthHandlerInterface interface {
	Login(res http.ResponseWriter, req *http.Request)
}

type AuthHandler struct {
	authService services.AuthServiceInterface
}

func NewAuthHandler(authService services.AuthServiceInterface) AuthHandlerInterface {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		utils.SendResponse(res, "Failed to parse form data", http.StatusBadRequest, nil)
		return
	}
	
	creds := models.LoginRequest{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	token, err := h.authService.Login(req.Context(), creds)
	if err != nil {
		utils.SendResponse(res, "Failed to authenticate", http.StatusUnauthorized, nil)
		return
	}

	utils.SendResponse(res, "Login successful", http.StatusOK, map[string]string{"token": token})
}