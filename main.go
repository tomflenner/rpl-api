package main

import (
	"fmt"
	"log"

	"github.com/b4cktr4ck5r3/rpl-api/config"
	"github.com/b4cktr4ck5r3/rpl-api/router"
	"github.com/gofiber/fiber/v2"
)

var port, Port string

func initializeEnv() {
	port = config.GetConfigValue("PORT")
	if port == "" {
		log.Fatal("Port is missing from env")
	}
}

func main() {
	initializeEnv()

	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(fmt.Sprintf(":%s", port))
}
