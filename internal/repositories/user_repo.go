package repositories

import (
	"furniture-shop-backend/internal/database"
	"furniture-shop-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *fiber.Ctx
}