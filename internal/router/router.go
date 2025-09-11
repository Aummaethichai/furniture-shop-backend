package router

import (
	"furniture-shop-backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// User routes
	userGroup := api.Group("/users")
	userGroup.Post("/register", handlers.Register)
	// userGroup.Post("/login", handlers.Login)
	// userGroup.Get("/", handlers.GetAllUsers)
	// userGroup.Get("/:id", handlers.GetUserByID)
	// userGroup.Put("/:id", handlers.UpdateUser)
	// userGroup.Delete("/:id", handlers.DeleteUser)

	// // Product routes
	// productGroup := api.Group("/products")
	// productGroup.Post("/", handlers.CreateProduct)
	// productGroup.Get("/", handlers.GetAllProducts)
	// productGroup.Get("/:id", handlers.GetProductByID)
	// productGroup.Put("/:id", handlers.UpdateProduct)
	// productGroup.Delete("/:id", handlers.DeleteProduct)

	// // Order routes
	// orderGroup := api.Group("/orders")
	// orderGroup.Post("/", handlers.CreateOrder)
	// orderGroup.Get("/", handlers.GetAllOrders)
	// orderGroup.Get("/:id", handlers.GetOrderByID)
	// orderGroup.Put("/:id", handlers.UpdateOrder)
	// orderGroup.Delete("/:id", handlers.DeleteOrder)
}