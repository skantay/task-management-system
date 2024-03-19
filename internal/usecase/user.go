package usecase

import (
	"context"
	"fmt"

	"github.com/skantay/task-management-system/internal/entities"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(ctx context.Context, user entities.User, password []byte) (entities.User, error)

	Exists(ctx context.Context, username, email string) (bool, error)
}

type UserUsecase interface {
	Create(ctx context.Context, user entities.User, password string) (entities.User, error)
}

type userUsecase struct {
	repo UserRepository
}

func NewUserUsecase(a UserRepository) UserUsecase {
	return userUsecase{a}
}

func (u userUsecase) Create(ctx context.Context, user entities.User, password string) (entities.User, error) {
	if exists, err := u.repo.Exists(ctx, user.Username, user.Email); err != nil {
		return entities.User{}, err
	} else if exists {
		return entities.User{}, entities.ErrExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to hash a password: %w", err)
	}

	created, err := u.repo.Create(ctx, user, hashedPassword)
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to create a user: %w", err)
	}

	return created, nil
}
