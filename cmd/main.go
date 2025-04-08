package main

import (
	"fmt"
	"go-boilerplate/api"
	"go-boilerplate/configs"
	"go-boilerplate/internal/app"
	"log"
	"net/http"
)

func main() {
	app := app.InitApp()
	router := api.SetupRouter(app)

	fmt.Printf("Server is running on http://%s\n", configs.Address)
	log.Fatal(http.ListenAndServe(configs.Address, router))
}