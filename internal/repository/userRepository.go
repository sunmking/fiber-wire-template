package repository

import (
	"github.com/gofiber/fiber/v2"
)

type userRepository struct {
	*Repository
}
type UserRepository interface {
	Create(ctx *fiber.Ctx) error
	GetList(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

func NewUserRepository(r *Repository) UserRepository {
	return &userRepository{r}
}

func (u userRepository) Create(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetList(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetOne(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Update(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Delete(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
