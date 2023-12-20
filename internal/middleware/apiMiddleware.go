package middleware

import (
	"fiber-wire-template/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func ApiMiddleware(jwt *jwt.JWT) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Set("Access-Control-Allow-Origin", "*")
		ctx.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if ctx.Method() == "OPTIONS" {
			ctx.Status(204)
		}
		var token string
		if ctx.Get("Authorization") != "" {
			token = ctx.Get("Authorization")
		} else if ctx.Get("access_token") != "" {
			token = ctx.Get("access_token")
		}
		if token == "" {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		if x, err := jwt.ParseToken(token); err != nil {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		} else {
			ctx.Locals("user", x)
		}
		return ctx.Next()
	}
}
