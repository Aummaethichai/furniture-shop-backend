package database

import (
	"furniture-shop-backend/config"
	models "furniture-shop-backend/internal/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectPostgres(cfg *config.DatabaseConfig) *gorm.DB {
	var err error

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			// IgnoreRecordNotFoundError: true,
			ParameterizedQueries: true,
			Colorful:             true,
		},
	)
	cfg.BuildPostgresDSN()
	DB, err = gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		panic("❌ failed to connect database")
		// log.Fatal("❌ failed to connect database: ", err)
	}
	log.Println("✅ Database connected")

	err = DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{})
	if err != nil {
		panic("❌ migration failed")
		// log.Fatal("❌ migration failed: ", err)
	}
	log.Println("✅ Migration success")

	return DB
}
