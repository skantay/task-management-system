package repository

import (
	"context"

	"github.com/skantay/task-management-system/internal/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user entities.User) error
	Get(ctx context.Context, id int64) (entities.User, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, user entities.User) (entities.User, error)

	Exists(ctx context.Context, user entities.User) (bool, error)
}
