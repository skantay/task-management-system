package repository

import (
	"context"

	"github.com/skantay/task-management-system/internal/entities"
)

type TaskRepository interface {
	Create(ctx context.Context, task entities.Task) (entities.Task, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, task entities.Task) (entities.Task, error)
	Get(ctx context.Context, id int64) (entities.Task, error)

	List(ctx context.Context, limit, offset uint) ([]entities.Task, error)
}
