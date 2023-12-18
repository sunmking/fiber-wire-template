package route

import (
	"fiber-wire-template/internal/handler"
	"fiber-wire-template/pkg/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(
	userHandler handler.UserHandler,
) *server.FiberServer {
	ser := &server.FiberServer{
		App: fiber.New(fiber.Config{
			CaseSensitive: true,
			StrictRouting: true,
			ServerHeader:  "Fiber",
			BodyLimit:     1024 * 1024 * 10,
		}),
	}

	ser.App.Use(cors.New())
	// logger
	// cors
	ser.App.Use(cors.New())
	// static
	ser.App.Static("/", "resources/")
	// ping
	ser.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	ser.App.Get("/user", userHandler.GetUser)
	ser.App.Get("/users", userHandler.GetUserList)

	return ser
}
