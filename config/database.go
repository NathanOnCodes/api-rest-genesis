package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var UseDatabase *gorm.DB

func ConnectDB() *gorm.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("HOST"), os.Getenv("DB_NAME"))
	time.Sleep(5 * time.Second)
	
	conn, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Printf("Erro ao conectar, mas espere um pouco, provavelmente houve conflito com sua porta, caso demore tente alterar na variavel dns api-rest-genesis-database para 127.0.0.1:")
			
	}
	UseDatabase = conn
	return UseDatabase
}
