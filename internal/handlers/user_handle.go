package handlers

import (
	"furniture-shop-backend/internal/database"
	"furniture-shop-backend/internal/dto"
	models "furniture-shop-backend/internal/models"
	"furniture-shop-backend/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	validator *validator.Validate
	service   *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{
		service:   s,
		validator: validator.New(),
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest

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
		"status_code":    fiber.StatusCreated,
		"status_message": "Success",
		"status":         "OK",
		"message":        "User registered successfully",
		"data":           user,
	})
}

func (h *UserHandler) FindUserByID(c *fiber.Ctx) error {
	req := dto.FindUserRequest{
		ID: c.Params("id"),
	}

	// ✅ validate dto
	if err := h.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			// "status_message": "Bad Request",
			"status": "Fail",
			"error":  err.Error(),
		})
	}

	user, err := h.service.FindUserByID(req.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			// "status_message": err.Error(),
			"status": "Fail",
			"error":  err.Error(),
		})
	}

	return c.JSON(user)
}

func (h *UserHandler) FindAllUser(c *fiber.Ctx) error {
	users, err := h.service.FindAllUser()
	if err != nil {
		return dto.ResponseInternalServerError(c, "failed to fetch users", err.Error())
	}

	// ถ้าอยาก handle response แบบ users ไม่พบข้อมูลใช้แบบนี้ได้
	// if len(users) == 0 {
	// 	return c.SendString("ไม่พบข้อมูล")
	// }

	return dto.ResponseOK(c, "users fetched successfully", users)
}
