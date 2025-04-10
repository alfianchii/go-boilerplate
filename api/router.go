package api

import (
	"go-boilerplate/internal/app"
	"go-boilerplate/internal/middlewares"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Post("/login", app.AuthHandler.Login)

		r.Group(func(r chi.Router) {
			r.Use(middlewares.AuthMiddleware("admin", app.AuthService))
			r.Get("/dashboard", app.DashboardHandler.Dashboard)
		})
	})

	return r
}