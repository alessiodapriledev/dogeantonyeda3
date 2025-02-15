package config

import (
	"carabinieri-backend/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// PostgreSQL connection string: user:password@host:port/dbname
	dsn := "user=user password=password dbname=dbname host=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate your models
	db.AutoMigrate(&models.Alert{})

	return db
}
