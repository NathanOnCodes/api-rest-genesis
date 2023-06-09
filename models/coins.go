package models

import (
	"github.com/NathanCavalcanteFerreira/api-rest-genesis/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Coin struct {
    gorm.Model
    Name string `gorm:"unique;not null"`
    Symbol string `gorm:"not null"`
}

func PopulateCoinTable(c *fiber.Ctx) error {
    var coins []Coin
    db.UseDatabase.Find(&coins)
    if len(coins) > 0 {
        return nil
    }

    coins = []Coin{
        {Name: "BRL", Symbol: "R$"},
        {Name: "USD", Symbol: "$"},
        {Name: "BTC", Symbol: "₿"},
        {Name: "EUR", Symbol: "€"},
    }

    if err := db.UseDatabase.Create(&coins).Error; err != nil {
        return err
    }

    return nil
}

func FindSymbolByCoinName(name string) (string, error) {
	var symbol string
	db.UseDatabase.Model(&Coin{}).Where("name = ?", name).Pluck("symbol", &symbol)
	return symbol, nil
}