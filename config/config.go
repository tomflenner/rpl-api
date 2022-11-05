package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfigValue(key string) string {
	var err error = godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading env")
	}

	return os.Getenv(key)
}
