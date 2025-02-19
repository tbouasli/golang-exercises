package database

import (
	"go-project/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("./database/test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados:", err)
	}

	// Chamando modelo
	err = db.AutoMigrate(&models.People{})
	if err != nil {
		log.Fatal("Erro ao migrar o modelo:", err)
	}

	// Inserindo modelo
	db.Create(&models.People{Name: "Kaique", Age: 20})
}
