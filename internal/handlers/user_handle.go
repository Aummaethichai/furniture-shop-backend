package handlers

import (
	"furniture-shop-backend/internal/database"
	models "furniture-shop-backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Register handles user registration
func Register(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Role     string `json:"role"` // admin, staff, customer
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	user := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Name:     req.Name,
		Role:     req.Role,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status_code": fiber.StatusCreated,
		"status_message":	"Success",
		"status":	"OK",
		"message":	"User registered successfully",
		"data":	user,
	})
}

// func (h *UserHandler) Register(c *fiber.Ctx) error {
//     var body struct {
//         Email    string `json:"email"`
//         Name     string `json:"name"`
//         Password string `json:"password"`
//     }
//     if err := c.BodyParser(&body); err != nil {
//         return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
//     }

//     user, err := h.service.Register(body.Email, body.Name, body.Password)
//     if err != nil {
//         return c.Status(400).JSON(fiber.Map{"error": err.Error()})
//     }

//     return c.Status(201).JSON(user)
// }
