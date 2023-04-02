package main

import (
	
	"github.com/NathanCavalcanteFerreira/api-rest-genesis/controllers"
	"github.com/NathanCavalcanteFerreira/api-rest-genesis/db"
	"github.com/NathanCavalcanteFerreira/api-rest-genesis/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db.InitDB()
	db.DB.AutoMigrate(&models.Coin{})
	models.PopulateCoinTable(nil)
	db.DB.AutoMigrate(&models.Conversion{})

	app.Get("/", func(c *fiber.Ctx) error {
		message := "Bem vindo, para converter as moedas siga esse exemplo:\nhttp://localhost:8000/exchange/valor/moeda-atual/moeda-destino/cotação\nvoce tambem pode conferir o seu histórico de conversões na rota /logs"
		return c.SendString(message)
	})

	app.Get("/exchange/:amount/:from/:to/:rate", controller.Convert)

	app.Get("/logs", controller.GetConversions)

	app.Listen(":8000")
}
