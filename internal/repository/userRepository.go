package repository

import (
	"context"
	"fiber-wire-template/internal/model"
	"fiber-wire-template/pkg/util/table"
	// "fmt" // Removed as debug prints are removed
	// "github.com/gofiber/fiber/v2" // Removed as fiber.Ctx is no longer used
)

type userRepository struct {
	*Repository
}
type UserRepository interface {
	// Create should accept context.Context and a user model.
	// Example: Create(ctx context.Context, user *model.User) error
	Create(ctx context.Context, user *model.User) error // Placeholder, to be implemented
	GetList(ctx context.Context, page, pageSize int) ([]model.User, error)
	// GetOne should accept context.Context and a user ID.
	// Example: GetOne(ctx context.Context, id uint) (*model.User, error)
	GetOne(ctx context.Context, id uint) (*model.User, error) // Placeholder, to be implemented
	// Update should accept context.Context, user ID, and a user model or update fields.
	// Example: Update(ctx context.Context, id uint, user *model.User) error
	Update(ctx context.Context, id uint, user *model.User) error // Placeholder, to be implemented
	// Delete should accept context.Context and a user ID.
	// Example: Delete(ctx context.Context, id uint) error
	Delete(ctx context.Context, id uint) error // Placeholder, to be implemented
}

func NewUserRepository(r *Repository) UserRepository {
	return &userRepository{r}
}

// Create implements UserRepository.
// It should accept context.Context and a user model.
// Example: func (u *userRepository) Create(ctx context.Context, user *model.User) error
func (u *userRepository) Create(ctx context.Context, user *model.User) error {
	//TODO implement me:
	// - Accept context.Context.
	// - Accept specific model or parameters instead of *fiber.Ctx.
	// - Implement the actual database logic.
	panic("implement me")
}

func (u *userRepository) GetList(ctx context.Context, page, pageSize int) ([]model.User, error) {
	var users []model.User
	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	q := u.Db.Select("*").
		From(table.TbaUser).
		OrderBy("id DESC").
		Limit(int64(pageSize)).
		Offset(int64(offset))

	// Pass context to the query if ozzo-dbx supports it with `With()` or similar
	// For now, assuming q.All(&users) uses a default context or the one from u.Db
	err := q.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetOne implements UserRepository.
// It should accept context.Context and a user ID.
// Example: func (u *userRepository) GetOne(ctx context.Context, id uint) (*model.User, error)
func (u *userRepository) GetOne(ctx context.Context, id uint) (*model.User, error) {
	//TODO implement me:
	// - Accept context.Context.
	// - Accept specific model or parameters instead of *fiber.Ctx.
	// - Implement the actual database logic.
	panic("implement me")
}

// Update implements UserRepository.
// It should accept context.Context, user ID, and a user model or update fields.
// Example: func (u *userRepository) Update(ctx context.Context, id uint, user *model.User) error
func (u *userRepository) Update(ctx context.Context, id uint, user *model.User) error {
	//TODO implement me:
	// - Accept context.Context.
	// - Accept specific model or parameters instead of *fiber.Ctx.
	// - Implement the actual database logic.
	panic("implement me")
}

// Delete implements UserRepository.
// It should accept context.Context and a user ID.
// Example: func (u *userRepository) Delete(ctx context.Context, id uint) error
func (u *userRepository) Delete(ctx context.Context, id uint) error {
	//TODO implement me:
	// - Accept context.Context.
	// - Accept specific model or parameters instead of *fiber.Ctx.
	// - Implement the actual database logic.
	panic("implement me")
}
