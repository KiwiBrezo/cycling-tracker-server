package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetENVByKey(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("(GetENVByKey) Error loading .env file")
	}

	return os.Getenv(key)
}
