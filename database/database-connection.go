package database

import (
	"go-project/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./database/test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados:", err)
	}

	// Chamando modelo
	err = DB.AutoMigrate(&models.People{})
	if err != nil {
		log.Fatal("Erro ao migrar o modelo:", err)
	}
}
