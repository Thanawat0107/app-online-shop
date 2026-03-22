package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	GO_ENV      string
	APP_PORT    string
	APP_HOST    string
	APP_NAME    string
	APP_VERSION string
	APP_CODE    string

	DB_PORT     string
	DB_HOST     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading file env.")
	}
	
	return &Env{
		GO_ENV: getEnv("GO_ENV", "development", false),
		APP_PORT: getEnv("APP_PORT", "8000", false),
		APP_HOST: getEnv("APP_HOST", "", false),
		APP_NAME: getEnv("APP_NAME", "", false),
		APP_VERSION: getEnv("APP_VERSION", "", false),
		APP_CODE: getEnv("APP_CODE", "", false),

		DB_PORT: getEnv("DB_PORT", "", true),
		DB_HOST: getEnv("DB_HOST", "", true),
		DB_USERNAME: getEnv("DB_USERNAME", "", true),
		DB_PASSWORD: getEnv("DB_PASSWORD", "", true),
		DB_NAME: getEnv("DB_NAME", "", true),
	}
}

func getEnv(key string, defaultValue string, required bool) string {
	if val, exist := os.LookupEnv(key); exist && val != "" {
		return val
	}
	if required {
		log.Fatalf("Error Missing required value env key: %s", key)
	}

	return defaultValue
}
