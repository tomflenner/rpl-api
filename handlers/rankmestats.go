package handlers

import (
	"strconv"

	"github.com/b4cktr4ck5r3/rpl-api/database"
	"github.com/b4cktr4ck5r3/rpl-api/models"
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
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(result)
}

func GetPlayerBySteamID(c *fiber.Ctx) error {
	result, err := database.SelectPlayerBySteamId(c.Params("steam_id"))

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(result)
}

func GetPlayersTop(c *fiber.Ctx) error {
	by := c.Query("by")

	if by == "" {
		return c.Status(400).JSON("Missing query params by=, use for example /top?by=hs")
	}

	limit := c.Query("limit", "10")
	intLimit, err := strconv.Atoi(limit)

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	var result []models.Player

	if by == "hs" {
		result, err = database.SelectPlayersTopByHs(intLimit);
	} else if by == "kd" {
		result, err = database.SelectPlayersTopByKd(intLimit);
	} else if by == "rank" {
		result, err = database.SelectPlayersTopByRank(intLimit);
	} else {
		return c.Status(400).JSON("Unknown query params by=, use for example /top?by=hs")
	}

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(result)
}
