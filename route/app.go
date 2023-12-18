package route

import (
	"fiber-wire-template/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(
	userHandler handler.UserHandler,
) *fiber.App {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		BodyLimit:     1024 * 1024 * 10,
	})
	app.Use(cors.New())
	// logger
	// cors
	app.Use(cors.New())
	// static
	app.Static("/", "resources/")
	// ping
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	app.Get("/user", userHandler.GetUser)

	return app
}
