package fileStorageService

import (
	"log"
	"os"
)

var (
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
)

func LoadConfig() {
	DBHost = os.Getenv("DB_HOST")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	DBPort = os.Getenv("DB_PORT")

	DBSSLMode = os.Getenv("DB_SSLMODE")
	if DBSSLMode == "" {
		DBSSLMode = "disable" // Set default value if not provided
	}

	if DBHost == "" || DBUser == "" || DBPassword == "" || DBName == "" || DBPort == "" {
		log.Fatal("Missing required environment variables")
	}

	log.Println("Configuration loaded successfully")
}
