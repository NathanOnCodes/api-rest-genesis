package models

import (
	"github.com/NathanCavalcanteFerreira/api-rest-genesis/db"
	"gorm.io/gorm"
)

type Conversion struct {
	gorm.Model
	Value    float64 `gorm:"not null"`
	Symbol string `gorm:"not null"`
}


func PopulateConversion(value float64, symbol string) {
	vc := Conversion{
		Value: value,
		Symbol: symbol,
	}	
	db.DB.Create(&vc)
}

func FindSymbolByCoinName(name string) string {
	var symbol string
	db.DB.Model(&Coin{}).Where("name = ?", name).Pluck("symbol", &symbol)
	return symbol
}