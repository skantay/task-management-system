package usecase

import (
	"context"
	"fmt"

	"github.com/skantay/task-management-system/internal/entities"
)

type TaskRepository interface {
	Create(ctx context.Context, task entities.Task) (entities.Task, error)
	Delete(ctx context.Context, id int64) error
}

type TaskUsecase interface {
	Create(ctx context.Context, task entities.Task) (entities.Task, error)
	Delete(ctx context.Context, id int64) error

	AddParticipant(ctx context.Context, usernames []string) ([]string, error)
	DeleteParticipants(ctx context.Context, usernames []string) ([]string, error)
	GetParticipants(ctx context.Context) []entities.User
}

type taskUsecase struct {
	repo TaskRepository
}

func NewTaskUsecase(a UserRepository) UserUsecase {
	return userUsecase{a}
}

func (t taskUsecase) Create(ctx context.Context, task entities.Task) (entities.Task, error) {
	task, err := t.repo.Create(ctx, task)
	if err != nil {
		return entities.Task{}, fmt.Errorf("failed to create a task: %w", err)
	}

	return task, nil
}

func (t taskUsecase) Delete(ctx context.Context, id int64) error {
	if err := t.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete a task: %w", err)
	}

	return nil
}
