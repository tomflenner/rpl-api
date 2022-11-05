package main

import (
	"github.com/b4cktr4ck5r3/rpl-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(":3000")
}
