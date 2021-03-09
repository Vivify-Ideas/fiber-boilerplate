package api

import (
	"github.com/gofiber/fiber/v2"
)

// Home - render landing route
func Home(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{})
}
