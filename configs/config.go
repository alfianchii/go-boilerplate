package configs

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppURL string
	AppPort string

	DBDriver string
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string

	JWTSecret string
}

var (
	Address = fmt.Sprintf("%s:%s", GetENV("APP_URL"), GetENV("APP_PORT"))
)

func InitENV() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		AppName: GetENV("APP_NAME"),
		AppURL: GetENV("APP_URL"),
		AppPort: GetENV("APP_PORT"),
		DBDriver: GetENV("DB_DRIVER"),
		DBHost: GetENV("DB_HOST"),
		DBPort: GetENV("DB_PORT"),
		DBName: GetENV("DB_DATABASE"),
		DBUser: GetENV("DB_USERNAME"),
		DBPass: GetENV("DB_PASSWORD"),
		JWTSecret: GetENV("JWT_SECRET"),
	}
}

func GetENV(key string) string {
	dotEnv, err := godotenv.Read()
	if err != nil {
		log.Fatalf("Error reading .env file")
	}

	return dotEnv[key]
}