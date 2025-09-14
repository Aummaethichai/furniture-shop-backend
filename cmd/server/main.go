package main

import (
	"furniture-shop-backend/config"
	"furniture-shop-backend/internal/database"
	"furniture-shop-backend/internal/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// connect DB
	cfg := config.LoadConfig()
	pg := database.ConnectPostgres(&cfg.Database)
	sqlDB, err := pg.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}
	defer sqlDB.Close()


	app := fiber.New()

	router.SetupRoutes(app)

	log.Println("🚀 Server running at :8080")
	log.Fatal(app.Listen(":8080"))
}
