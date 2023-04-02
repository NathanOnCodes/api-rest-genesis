package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/NathanCavalcanteFerreira/api-rest-genesis/db"
	"github.com/NathanCavalcanteFerreira/api-rest-genesis/models"
)

func GetConversions(c *fiber.Ctx) error{
	var conversions []models.Conversion
	result := db.DB.Find(&conversions)
	
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ocorreu um erro ao buscar as conversões",
		})
	}
	
	if len(conversions) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Ainda não houve conversões",
		})
	}

	conversionsJson := make([]fiber.Map, len(conversions))

	for i, conversion := range conversions {
		conversionsJson[i] = fiber.Map{
			"simbolo": conversion.Symbol,
			"valor": conversion.Value,
		}
	}

	return c.Status(fiber.StatusOK).JSON(conversionsJson)
}