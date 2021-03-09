package routes

import (
	"github.com/gofiber/fiber/v2"
	Controllers "github.com/vivify-ideas/fiber_boilerplate/controllers/api"
	"github.com/vivify-ideas/fiber_boilerplate/middlewares"
)

// RegisterAPI - registering API routes
func RegisterAPI(api fiber.Router) {
	registerAuth(api)
	registerFiles(api)
}

func registerAuth(api fiber.Router) {
	auth := api.Group("/auth")

	auth.Post("/register", Controllers.Register)
	auth.Post("/login", Controllers.Login)
	auth.Get("/me", middlewares.Auth(), Controllers.GetAuthenticatedUser)
	auth.Post("/reset-password", Controllers.ResetPassword)
	auth.Post("/reset-password-complete", Controllers.ResetPasswordComplete)
}

func registerFiles(api fiber.Router) {
	files := api.Group("/files")

	files.Post("/", middlewares.Auth(), Controllers.UploadFiles)
}
