package controller

import (
	"log"
	"strconv"

	"github.com/NathanCavalcanteFerreira/api-rest-genesis/models"
	"github.com/gofiber/fiber/v2"
)

func Convert(c *fiber.Ctx) error {
	amount, amountErr := strconv.ParseFloat(c.Params("amount"), 64)
	rate, rateErr := strconv.ParseFloat(c.Params("rate"), 64)
	to := c.Params("to")

	if amountErr != nil || rateErr != nil{
		log.Println("valor errado")
	}

	converted := Calc(amount, rate)

	symbol, symbolErr := models.FindSymbolByCoinName(to)
	if symbolErr != nil {
		log.Println("moeda nao encontrada")
	}


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
