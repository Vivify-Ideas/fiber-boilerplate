package middlewares

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/vivify-ideas/fiber_boilerplate/config"
)

// RegisterMiddlewares -> register middlewares globally
func RegisterMiddlewares(app *fiber.App, web *fiber.Router, api *fiber.Router) {
	registerGlobalMiddlewares(app)
	registerAPIMiddlewares(api)
	registerWebMiddlewares(web)
}

func registerGlobalMiddlewares(app *fiber.App) {
	// Register fiber built-in middlewares
	registerBuiltInMiddlewares(app)

	// 404 page returned for not found URLs
	app.Use(func(c *fiber.Ctx) error {
		if err := c.SendStatus(fiber.StatusNotFound); err != nil {
			panic(err)
		}
		if err := c.Render("errors/404", fiber.Map{}); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return nil
	})
}

func registerBuiltInMiddlewares(app *fiber.App) {
	env := config.App.Env

	// Recover middleware for Fiber that recovers from panics anywhere
	// in the stack chain and handles the control to the centralized ErrorHandler.
	// https://docs.gofiber.io/api/middleware/recover
	app.Use(recover.New())

	// ETag middleware for Fiber that lets caches be more efficient
	// and save bandwidth, as a web server does not need to resend
	// a full response if the content has not changed.
	// https://docs.gofiber.io/api/middleware/etag
	app.Use(etag.New())

	// CORS middleware for Fiber that that can be used to enable Cross-Origin
	// Resource Sharing with various options.
	// https://docs.gofiber.io/api/middleware/cors
	corsAllowCredentials, _ := strconv.ParseBool(env["CORS_ALLOW_CREDENTIALS"])
	app.Use(cors.New(cors.Config{
		AllowOrigins:     env["CORS_ALLOW_ORIGINS"],
		AllowHeaders:     env["CORS_ALLOW_HEADERS"],
		AllowMethods:     env["CORS_ALLOW_METHODS"],
		AllowCredentials: corsAllowCredentials,
	}))
}

func registerWebMiddlewares(web *fiber.Router) {}

func registerAPIMiddlewares(api *fiber.Router) {}
