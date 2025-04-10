package services

import (
	"context"
	"go-boilerplate/internal/utils"
)

type DashboardServiceInterface interface {
	GetDashboardData(ctx context.Context, userClaims *utils.UserClaims) (*utils.UserClaims, error)
}

type DashboardService struct {}

func NewDashboardService() DashboardServiceInterface {
	return &DashboardService{}
}

func (s *DashboardService) GetDashboardData(ctx context.Context, userClaims *utils.UserClaims) (*utils.UserClaims, error) {
	return userClaims, nil
}