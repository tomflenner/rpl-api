package config

import (
	"fmt"
	"log"
	"os"

	"github.com/b4cktr4ck5r3/rpl-api/models"
	"github.com/joho/godotenv"
)

func loadDotEnv() {
	var err error = godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getConfigValue(key string) string {
	var value string
	value = os.Getenv(key)

	if value == "" {
		log.Fatal(fmt.Sprintf("Error loading %s from env", key))
	}

	return value
}

var Config models.Config

func InitializeConfig() {
	loadDotEnv()

	Config = models.Config{}

	var port string
	port = getConfigValue("PORT")

	Config.PORT = port
}
