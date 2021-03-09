package api

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/vivify-ideas/fiber_boilerplate/requests"
	services "github.com/vivify-ideas/fiber_boilerplate/services"
)

// Register -> registration action
func Register(c *fiber.Ctx) error {
	var data requests.RegisterRequest
	// user := new(models.User)
	var err = c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	if err := data.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}
	authResult, err := services.CreateUser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(201).JSON(authResult)
}

// Login -> login action
func Login(c *fiber.Ctx) error {
	var data requests.LoginRequest
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	if err := data.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	authResult, err := services.Login(data.Email, data.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(authResult)
}

// GetAuthenticatedUser -> retrieve auth user
func GetAuthenticatedUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	user, err := services.GetUserFromJWT(*token)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(user)
}

// ResetPassword -> start reset password process
func ResetPassword(c *fiber.Ctx) error {
	var data requests.ResetPasswordRequest
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	if err := data.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	ok, _ := services.ResetPassword(data.Email)

	return c.JSON(fiber.Map{
		"success": ok,
	})
}

// ResetPasswordComplete -> finish reset password process
func ResetPasswordComplete(c *fiber.Ctx) error {
	var data requests.ResetPasswordCompleteRequest
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	if err := data.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	ok, _ := services.ResetPasswordComplete(data.Token, data.Password)

	return c.JSON(fiber.Map{
		"success": ok,
	})
}
