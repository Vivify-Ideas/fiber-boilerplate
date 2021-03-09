package routes

import (
	"github.com/gofiber/fiber/v2"
	Controllers "github.com/vivify-ideas/fiber_boilerplate/controllers/web"
)

// RegisterWeb - registering web routes
func RegisterWeb(api fiber.Router) {
	registerHome(api)
}

func registerHome(api fiber.Router) {
	auth := api.Group("/")

	auth.Get("/", Controllers.Home)
}
