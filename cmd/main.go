package main

import (
	"fmt"
	"go-boilerplate/configs"
	"go-boilerplate/internal/database"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := configs.InitENV()
	database.InitDB(cfg)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fmt.Printf("Server is running on http://%s\n", configs.Address)
	log.Fatal(http.ListenAndServe(configs.Address, r))
}