package handlers

import "github.com/gofiber/fiber/v2"

func GetVersion(c *fiber.Ctx) error {
	response := &fiber.Map{
		"version": 1,
	}

	return c.Status(200).JSON(response)
}
