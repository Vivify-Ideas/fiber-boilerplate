package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vivify-ideas/fiber_boilerplate/config"
	"github.com/vivify-ideas/fiber_boilerplate/database"
	"github.com/vivify-ideas/fiber_boilerplate/middlewares"
	"github.com/vivify-ideas/fiber_boilerplate/routes"
)

func main() {
	app := fiber.New(
		fiber.Config{
			Views:        config.App.Views,
			ErrorHandler: config.App.ErrorHandler,
		},
	)

	// Initialize database
	database.Init()

	// Load static
	config.App.LoadStatic(app)

	// Register web routes
	web := app.Group("")
	routes.RegisterWeb(web)

	// Register application API routes (using the /api/v1 group)
	api := app.Group("/api")
	apiv1 := api.Group("/v1")
	routes.RegisterAPI(apiv1)

	middlewares.RegisterMiddlewares(app, &web, &api)

	// Default is :3000, but you can override the value in .env
	config := config.App.Env
	app.Listen(config["APP_LISTEN_PORT"])
}
