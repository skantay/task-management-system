package repository

import (
	"context"

	"github.com/skantay/task-management-system/internal/entities"
)

type UserTaskRepository interface {
	ListByUserID(ctx context.Context, id int64) ([]entities.Task, error)
	ListByTaskID(ctx context.Context, id int64) ([]entities.Task, error)
	ListByID(ctx context.Context, userID, taskID int64) ([]entities.Task, error)
}
