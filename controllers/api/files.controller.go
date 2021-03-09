package api

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	services "github.com/vivify-ideas/fiber_boilerplate/services"
)

// UploadFiles -> upload multiple files action
func UploadFiles(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	user, err := services.GetUserFromJWT(*token)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Parse the multipart form:
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	files := form.File["files"]

	storedFiles, errors := services.UploadFiles(c, files, &user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": errors})
	}
	return c.Status(fiber.StatusOK).JSON(storedFiles)
}
