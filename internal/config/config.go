package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort   string
	DBurl     string
	JWTSecret string
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		fmt.Println("kONTOL")
	}

	appPort := os.Getenv("APP_PORT")
	dbURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")

	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	if appPort == "" {
		log.Println("Warning: APP_PORT not set, defaulting to 8080")
		appPort = "8080"
	}
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	return &Config{
		AppPort:   appPort,
		DBurl:     dbURL,
		JWTSecret: jwtSecret,
	}
}
