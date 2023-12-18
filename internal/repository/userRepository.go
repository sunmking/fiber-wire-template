package repository

import (
	"fiber-wire-template/internal/model"
	"fiber-wire-template/pkg/util/table"
	"github.com/gofiber/fiber/v2"
)

type userRepository struct {
	*Repository
}
type UserRepository interface {
	Create(ctx *fiber.Ctx) error
	GetList(ctx *fiber.Ctx) ([]model.User, error)
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

func (u userRepository) GetList(ctx *fiber.Ctx) ([]model.User, error) {
	var users []model.User
	//TODO implement me
	var q = u.db.Select("*").From(table.TbaUser).OrderBy("id DESC")
	err := q.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
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
