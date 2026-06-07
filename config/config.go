package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	APPport          string
	DSN              string
	JWTSecretKey     string
	JWTExpiry        time.Duration
    JWTRefreshExpiry time.Duration
	REDISaddr        string
	REDISpassword    string
	FileServiceURL   string
}

func NewConfig() AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found!")
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
        log.Fatal("APP_PORT environment variable is required!")
    }

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environtment variable is required!")
	}

	accessExpiryStr := os.Getenv("JWT_EXPIRY")
	accessExpiry, err := time.ParseDuration(accessExpiryStr)
	if err != nil {
		log.Fatal("invalid JWT_EXPIRY format!")
	}

    refreshExpiryStr := os.Getenv("JWT_REFRESH_EXPIRY")
	refreshExpiry, err := time.ParseDuration(refreshExpiryStr)
	if err != nil {
		log.Fatal("invalid JWT_REFRESH_EXPIRY format!")
	}
	
    secretKey := os.Getenv("JWT_SECRET_KEY")
    if secretKey == "" {
        log.Fatal("JWT_SECRET_KEY environment variable is required!")
    }

	redisAddr := os.Getenv("REDIS_ADDR")
    if redisAddr == "" {
        log.Fatal("REDIS_ADDR environment variable is required!")
    }

	redisPassword := os.Getenv("REDIS_PASSWORD")
	if redisPassword == "" {
        log.Fatal("REDIS_PASSWORD environment variable is required!")
    }

	fileServiceURL := os.Getenv("FILE_SERVICE_URL")
	if fileServiceURL == "" {
		log.Fatal("FILE_SERVICE_URL environment variable is required!")
	}

	return AppConfig{
		APPport:          appPort,
		DSN:              dsn,
		JWTSecretKey:     secretKey,
        JWTExpiry:        accessExpiry,
        JWTRefreshExpiry: refreshExpiry,
		REDISaddr:        redisAddr,
		REDISpassword:    redisPassword,
		FileServiceURL:   fileServiceURL,
	}
}