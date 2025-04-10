package handlers

import (
	"go-boilerplate/internal/middlewares"
	"go-boilerplate/internal/services"
	"go-boilerplate/internal/utils"
	"net/http"
)

type DashboardHandlerInterface interface {
	Dashboard(res http.ResponseWriter, req *http.Request)
}

type DashboardHandler struct {
	dashboardService services.DashboardServiceInterface
}

func NewDashboardHandler(dashboardService services.DashboardServiceInterface) DashboardHandlerInterface {
	return &DashboardHandler{dashboardService: dashboardService}
}

func (h *DashboardHandler) Dashboard(res http.ResponseWriter, req *http.Request) {
	userClaims := req.Context().Value(middlewares.UserClaimsKey).(*utils.UserClaims)

	data, err := h.dashboardService.GetDashboardData(req.Context(), userClaims)
	if err != nil {
			utils.SendResponse(res, "Failed to fetch dashboard data", http.StatusInternalServerError, nil)
			return
	}

	utils.SendResponse(res, "Welcome to the dashboard", http.StatusOK, data)
}