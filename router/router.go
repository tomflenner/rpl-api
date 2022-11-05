package router

import (
	"github.com/b4cktr4ck5r3/rpl-api/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/version", handler.GetVersion)
}
