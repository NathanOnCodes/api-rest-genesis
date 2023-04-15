package models

import (
	"github.com/NathanCavalcanteFerreira/api-rest-genesis/config"
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
	db.UseDatabase.Create(&vc)
}

