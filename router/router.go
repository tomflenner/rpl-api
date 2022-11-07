package router

import (
	"github.com/b4cktr4ck5r3/rpl-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/version", handlers.GetVersion)
	api.Get("/players", handlers.GetPlayers)
	api.Get("/players/:steam_id", handlers.GetPlayerBySteamID)
	api.Get("/top10kd", handlers.GetPlayersTop10ByKd)
	api.Get("/top10hs", handlers.GetPlayersTop10ByHs)

}
