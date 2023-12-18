package handler

import (
	"fiber-wire-template/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
}
type userHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) UserHandler {
	return &userHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (u userHandler) CreateUser(ctx *fiber.Ctx) error {
	//TODO implement me
	return ctx.SendString("CreateUser")
}

func (u userHandler) GetUser(ctx *fiber.Ctx) error {
	//TODO implement me
	u.userService.GetList(ctx)
	return ctx.SendString("Hello, World! GetUser")
}

func (u userHandler) UpdateUser(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u userHandler) DeleteUser(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
