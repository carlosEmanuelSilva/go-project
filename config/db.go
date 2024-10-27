package config

import (
  "log"
  "tp/internal/models"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=carlos password=152535 dbname=books_db port=5432 sslmode=disable"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Falha ao conectar ao banco de dados:", err)
    }
    // Executa a migração para criar as tabelas
    database.AutoMigrate(&models.User{}, &models.Book{}, &models.Review{})
    
    DB = database
}
