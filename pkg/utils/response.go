package utils

import "github.com/gofiber/fiber/v2"

// SuccessResponse sends a JSON response with a success status and data
func SuccessResponse(c *fiber.Ctx, data interface{}, statusCode int) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}

// ErrorResponse sends a JSON response with an error status and message
func ErrorResponse(c *fiber.Ctx, err string, statusCode int) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "error",
		"message": err,
	})
}

