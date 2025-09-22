package router

import (
	"furniture-shop-backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	UserHandler *handlers.UserHandler
	// ProductHandler *handlers.ProductHandler
	// เพิ่ม OrderHandler, PaymentHandler ได้เรื่อยๆ
}

func SetupRoutes(app *fiber.App, h *Handlers) {
	api := app.Group("/api/v1")

	// User routes
	userGroup := api.Group("/user")
	userGroup.Get("/", h.UserHandler.FindAllUser)
	userGroup.Get("/find/:id", h.UserHandler.FindUserByID)
	userGroup.Post("/register", h.UserHandler.CreateUser)
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
