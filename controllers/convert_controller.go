package controller

import (
	"strconv"
	"github.com/NathanCavalcanteFerreira/api-rest-genesis/models"
	"github.com/gofiber/fiber/v2"
)

func Convert(c *fiber.Ctx) error {
	amount, _ := strconv.ParseFloat(c.Params("amount"), 64)
	rate, _ := strconv.ParseFloat(c.Params("rate"), 64)
	to := c.Params("to")

	converted := Calc(amount, rate)

	symbol := models.FindSymbolByCoinName(to)

	models.PopulateConversion(converted, symbol)

	return c.JSON(fiber.Map{
		"moeda": symbol,
		"valor": converted,
	})
}

func Calc(amount float64, rate float64) float64 {
	converted := amount * rate
	return converted
}
