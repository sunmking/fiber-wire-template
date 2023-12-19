package route

import (
	"fiber-wire-template/internal/handler"
	"fiber-wire-template/pkg/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
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

	var ConfigDefault = requestid.Config{
		Next:       nil,
		Header:     fiber.HeaderXRequestID,
		Generator:  utils.UUID,
		ContextKey: "requestid",
	}
	ser.App.Use(requestid.New(ConfigDefault))
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
