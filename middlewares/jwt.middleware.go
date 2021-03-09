package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/vivify-ideas/fiber_boilerplate/config"
)

// Auth -> protect routes with JWT authentication
func Auth() fiber.Handler {
	config := config.App.Env
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config["APP_SECRET"]),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"message": err.Error()})
}
