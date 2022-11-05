package handlers

import (
	"github.com/b4cktr4ck5r3/rpl-api/database"
	"github.com/gofiber/fiber/v2"
)

func GetVersion(c *fiber.Ctx) error {
	response := &fiber.Map{
		"version": 1,
	}

	return c.Status(200).JSON(response)
}

func GetPlayers(c *fiber.Ctx) error {
	result, err := database.SelectPlayers()

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err,
		})
	}

	return c.Status(200).JSON(result)
}
