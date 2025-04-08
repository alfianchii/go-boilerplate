package app

import (
	"go-boilerplate/configs"
	"go-boilerplate/internal/database"
	"go-boilerplate/internal/handlers"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/internal/services"
)

type App struct {
	DB *database.DB
	UserRepo repositories.UserRepositoryInterface
	AuthService services.AuthServiceInterface
	AuthHandler handlers.AuthHandlerInterface
}

func InitApp() *App {
	cfg := configs.InitENV()
	db := database.InitDB(cfg)

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	return &App{
		DB:          db,
		UserRepo:    userRepo,
		AuthService: authService,
		AuthHandler: authHandler,
	}
}