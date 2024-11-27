package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser		 string
	DBPassword string
	DBHost		 string
	DBPort		 string
	DBName	 	 string
	DBSSLMode	 string
}

func LoadConfig() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		DBUser: 		os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost: 	 	os.Getenv("DB_HOST"),
		DBPort: 		os.Getenv("DB_PORT"),
		DBName: 		os.Getenv("DB_NAME"),
		DBSSLMode: 	os.Getenv("DB_SSLMODE"),
	}
}