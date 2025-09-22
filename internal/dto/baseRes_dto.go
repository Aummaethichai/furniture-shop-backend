package dto

import "github.com/gofiber/fiber/v2"

type BaseResponse struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func ResponseOK(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(BaseResponse{
		StatusCode: fiber.StatusOK,
		Status:     "success",
		Message:    message,
		Data:       data,
	})
}

func ResponseBadRequest(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(BaseResponse{
		StatusCode: fiber.StatusBadRequest,
		Status:     "Fail",
		Message:    message,
		Data:       data,
	})
}

func ResponseCreated(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(BaseResponse{
		StatusCode: fiber.StatusCreated,
		Status:     "success",
		Message:    message,
		Data:       data,
	})
}

func ResponseInternalServerError(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusInternalServerError).JSON(BaseResponse{
		StatusCode: fiber.StatusInternalServerError,
		Status:     "Fail",
		Message:    message,
		Data:       data,
	})
}

func ResponseNotFound(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusNotFound).JSON(BaseResponse{
		StatusCode: fiber.StatusNotFound,
		Status:     "Fail",
		Message:    message,
		Data:       data,
	})
}

func ResponseUnauthorized(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusUnauthorized).JSON(BaseResponse{
		StatusCode: fiber.StatusUnauthorized,
		Status:     "Fail",
		Message:    message, // unauthorized
		Data:       data,
	})
}

func ResponseForbidden(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusForbidden).JSON(BaseResponse{
		StatusCode: fiber.StatusForbidden,
		Status:     "Fail",
		Message:    message, // Forbidden
		Data:       data,
	})
}
