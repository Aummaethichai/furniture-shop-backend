package database

import (
	"furniture-shop-backend/config"
	models "furniture-shop-backend/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) {
	var err error
	cfg := config.LoadConfig()
	DB, err = gorm.Open(postgres.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ failed to connect database: ", err)
	}
	log.Println("✅ Database connected")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ migration failed: ", err)
	}
	log.Println("✅ Migration success")
}
