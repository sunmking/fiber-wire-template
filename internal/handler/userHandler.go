package handler

import (
	"fiber-wire-template/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	GetUserList(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
}
type userHandler struct {
	*Handler
	userService service.UserService
}

func (u userHandler) GetUserList(ctx *fiber.Ctx) error {
	//TODO implement me
	users := u.userService.GetList(ctx)
	if users != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": users.Error(),
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"users": users,
		})
	}
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
