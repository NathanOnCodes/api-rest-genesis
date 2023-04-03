package db

import (
	"log"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var UseDatabase *gorm.DB

func ConnectDB() *gorm.DB {
	dns := "admin:admin@tcp(api-rest-genesis-database:3306)/mysqldatabase?charset=utf8mb4&parseTime=True&loc=Local"
	time.Sleep(5 * time.Second)
	
	conn, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Printf("Erro ao conectar, mas espere um pouco, provavelmente houve conflito com sua porta, caso demore tente alterar na variavel dns api-rest-genesis-database para 127.0.0.1:")
			
	}
	UseDatabase = conn
	return UseDatabase
}
