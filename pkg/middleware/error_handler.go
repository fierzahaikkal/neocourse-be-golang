package middleware

import (
	"fmt"
	"log"

	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func RecoveryMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				log.Printf("Recovered from panic: %v", err)
				_ = utils.ErrorResponse(c, "Internal Server Error", fiber.StatusInternalServerError)
			}
		}()

		return c.Next()
	}
}
