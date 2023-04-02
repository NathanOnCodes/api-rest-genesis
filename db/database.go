package db

import (
	"log"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dns := "admin:admin@tcp(db:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	time.Sleep(2 * time.Second)

	for i := 1; i <= 3; i++ {
		data, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
		if err != nil {
			log.Printf("Erro ao conectar ao banco de dados: %v. Tentativa %d de 3.", err, i)
			time.Sleep(3 * time.Second)
		} else {
			DB = data
			return DB
		}
	}
	log.Fatal("Não foi possível conectar ao banco de dados após 3 tentativas.")
	return nil
}
