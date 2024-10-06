package usecase

import (
	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/repository"
	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	"github.com/gofiber/fiber/v2"

	"github.com/fierzahaikkal/neocourse-be-golang/internal/model/user"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	UserRepo *repository.UserRepository
	log      *log.Logger
}

func NewAuthUseCase(userRepo *repository.UserRepository, log *log.Logger) *AuthUseCase {
	return &AuthUseCase{UserRepo: userRepo, log: log}
}

func (uc *AuthUseCase) SignUp(jwtSecret string) fiber.Handler {
    return func(c *fiber.Ctx) error {
		var req user.SignUpRequest

        if err := c.BodyParser(req); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
        }

		var user entity.User
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
        }

        user.ID = utils.GenUUID()
        user.Password = string(hashedPassword)

        if err := uc.UserRepo.Register(&user); err != nil {
            if err == utils.ErrUserExists {
                return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
            }
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
        }

        token, err := utils.GenerateJWT(&user, jwtSecret)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
        }

        return c.Status(fiber.StatusCreated).JSON(fiber.Map{"token": token})
    }
}

func (uc *AuthUseCase) SignIn(jwtSecret string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var req user.SignInRequest

        if err := c.BodyParser(&req); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Invalid request body",
            })
        }

        var user entity.User
        userFromDb, err := uc.UserRepo.FindByEmail(req.Email, &user); 
		if err != nil {
            if err == utils.ErrRecordNotFound {
                return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                    "error": "Invalid credentials",
                })
            }
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "An error occurred while processing your request",
            })
        }

        if err := bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(req.Password)); err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Invalid credentials",
            })
        }

        token, err := utils.GenerateJWT(&user, jwtSecret)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to generate token",
            })
        }

        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "token": token,
        })
    }
}
