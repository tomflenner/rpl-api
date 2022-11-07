package handlers

import (
	"github.com/MrWaggel/gosteamconv"
	"github.com/b4cktr4ck5r3/rpl-api/models"
	"github.com/b4cktr4ck5r3/rpl-api/utils"
	"github.com/gofiber/fiber/v2"
)

func PostPayloadToMakeToken(c *fiber.Ctx) error {
	p := new(models.TokensPayload)

	if err := c.BodyParser(p); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	_, err := gosteamconv.SteamStringToInt64(p.SteamID)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := utils.GenerateJWT(p.SteamID)

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	if token == "" {
		return c.Status(500).JSON(&fiber.Map{
			"error": "empty token",
		})
	}

	return c.Status(200).JSON(token)
}

func GetTokenValidity(c *fiber.Ctx) error {
	tokenString := c.Params("tokenString")

	isValid, err := utils.VerifyJWT(tokenString)

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	if !isValid {
		return c.Status(500).JSON(&fiber.Map{
			"error": "token not valid",
		})
	}

	return c.Status(200).JSON(true)
}
