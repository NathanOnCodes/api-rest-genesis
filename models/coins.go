package models

import (
	"github.com/NathanCavalcanteFerreira/api-rest-genesis/db"
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
    db.DB.Find(&coins)
    if len(coins) > 0 {
        return nil
    }

    coins = []Coin{
        {Name: "BRL", Symbol: "R$"},
        {Name: "USD", Symbol: "US$"},
        {Name: "BTC", Symbol: "₿"},
        {Name: "EUR", Symbol: "€"},
    }

    if err := db.DB.Create(&coins).Error; err != nil {
        return err
    }

    return nil
}
