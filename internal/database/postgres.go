package database

import (
	"furniture-shop-backend/config"
	models "furniture-shop-backend/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres(cfg *config.DatabaseConfig) *gorm.DB {
	var err error
	cfg.BuildPostgresDSN()
	// cfg := config.LoadConfig()
	DB, err = gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
	if err != nil {
		log.Fatal("❌ failed to connect database: ", err)
	}
	log.Println("✅ Database connected")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ migration failed: ", err)
	}
	log.Println("✅ Migration success")

	return DB
}
