package main

import (
	"furniture-shop-backend/config"
	"furniture-shop-backend/internal/database"
	"furniture-shop-backend/internal/handlers"
	repos "furniture-shop-backend/internal/repositories"
	"furniture-shop-backend/internal/router"
	"furniture-shop-backend/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// connect DB
	cfg := config.LoadConfig()
	pg := database.ConnectPostgres(&cfg.PostgresDB)
	sqlDB, err := pg.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}
	defer sqlDB.Close()

	// repo
	userRepo := repos.NewUserRepository(pg)
	// productRepo := repos.NewProductRepo(db)

	// service
	userService := services.NewUserService(userRepo)
	// productService := services.NewProductService(productRepo)

	// handler
	userHandler := handlers.NewUserHandler(userService)
	// productHandler := handlers.NewProductHandler(productService)

	// setup fiber
	app := fiber.New()

	// setup router
	router.SetupRoutes(app, &router.Handlers{
		UserHandler: userHandler,
		// ProductHandler: productHandler,
	})

	log.Println("🚀 Server running at :8080")
	log.Fatal(app.Listen(":8080"))
}
