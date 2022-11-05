package main

import (
	"fmt"

	"github.com/b4cktr4ck5r3/rpl-api/config"
	"github.com/b4cktr4ck5r3/rpl-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitializeConfig()

	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(fmt.Sprintf(":%s", config.Config.PORT))
}
