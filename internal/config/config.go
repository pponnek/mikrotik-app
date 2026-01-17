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
	EncryptionKEY string 
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		fmt.Println("failed to load env")
	}

	appPort := os.Getenv("APP_PORT")
	dbURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	encryKey := os.Getenv("ENCRYPTION_KEY")

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

	if encryKey == "" {
		log.Fatal("ENCRYPTION_KEY is required")
	}

	return &Config{
		AppPort:   appPort,
		DBurl:     dbURL,
		JWTSecret: jwtSecret,
		EncryptionKEY: encryKey,
	}
}
