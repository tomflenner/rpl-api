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
		log.Fatal("Error loading .env file: " + err.Error())
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

var Cfg models.Config

func InitializeConfig() {
	loadDotEnv()

	Cfg = models.Config{}

	var apiPort string = getConfigValue("API_PORT")
	Cfg.ApiPort = apiPort

	var dbUser string = getConfigValue("DB_USER")
	Cfg.DbUser = dbUser

	var dbPassword string = getConfigValue("DB_PASSWORD")
	Cfg.DbPassword = dbPassword

	var dbHost string = getConfigValue("DB_HOST")
	Cfg.DbHost = dbHost

	var dbPort string = getConfigValue("DB_PORT")
	Cfg.DbPort = dbPort

	var dbName string = getConfigValue("DB_NAME")
	Cfg.DbName = dbName
}
