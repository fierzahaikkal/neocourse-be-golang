package handler

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/model/user"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/usecase"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type AuthHandler struct {
	AuthUC   *usecase.AuthUseCase
	JWTSecret string
	log      *log.Logger
}

func NewAuthHandler(authUC *usecase.AuthUseCase, jwtSecret string, log *log.Logger) *AuthHandler {
	return &AuthHandler{AuthUC: authUC, JWTSecret: jwtSecret, log: log}
}

func (ah *AuthHandler) SignUp(c *fiber.Ctx) error {
	var req user.SignUpRequest
	var user entity.User

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err := ah.AuthUC.SignUp(&req)
	if err != nil {
		if err == utils.ErrUserExists {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	token, err := utils.GenerateJWT(&user, ah.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return utils.SuccessResponse(c, token, fiber.StatusOK)
}

func (ah *AuthHandler) SignIn(c *fiber.Ctx) error {
	var req user.SignInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	userFromDb, err := ah.AuthUC.SignIn(&req)
	if err != nil {
		if err == utils.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "An error occurred while processing your request"})
	}

	token, err := utils.GenerateJWT(userFromDb, ah.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": userFromDb,"token": token})
}

func (ah *AuthHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("userID")

	user, err := ah.AuthUC.GetUser(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})
}
