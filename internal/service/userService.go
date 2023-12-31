package service

import (
	"fiber-wire-template/internal/repository"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	GetList(ctx *fiber.Ctx) interface{}
	GetOne(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func NewUserService(service *Service, userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

func (u userService) Create(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) Update(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) GetList(ctx *fiber.Ctx) interface{} {
	if users, err := u.userRepo.GetList(ctx); err != nil {
		return err
	} else {
		fmt.Println(users)
		return users
	}
}

func (u userService) GetOne(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) Delete(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
