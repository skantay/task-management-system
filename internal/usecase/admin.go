package usecase

import (
	"context"
	"fmt"
	"unicode/utf8"

	"github.com/skantay/task-management-system/internal/entities"
)

type AdminRepository interface {
	List(ctx context.Context, limit, offset uint) ([]entities.User, error)
	GetBannedUsers(ctx context.Context, limit, offset uint) ([]entities.User, error)
	GetUser(ctx context.Context, id int64) (entities.User, error)
	DeleteUser(ctx context.Context, id int64) error
	IsAdmin(ctx context.Context, id int64) (bool, error)
	SetAdmin(ctx context.Context, id int64) error
	RevokeAdmin(ctx context.Context, id int64) error
	IsBanned(ctx context.Context, id int64) (bool, error)
	BanUser(ctx context.Context, id int64, description string) error
	UnbanUser(ctx context.Context, id int64) error
}

type AdminUsecase interface {
	List(ctx context.Context, limit, offset uint) ([]entities.User, error)
	GetBannedUsers(ctx context.Context, limit, offset uint) ([]entities.User, error)
	GetUser(ctx context.Context, id int64) (entities.User, error)
	DeleteUser(ctx context.Context, id int64) error
	SetAdmin(ctx context.Context, id int64) error
	RevokeAdmin(ctx context.Context, id int64) error
	BanUser(ctx context.Context, id int64, description string) error
	UnbanUser(ctx context.Context, id int64) error
}

type adminUsecase struct {
	repo AdminRepository
}

func NewAdminUsecase(adminRepository AdminRepository) AdminUsecase {
	return adminUsecase{adminRepository}
}

func (a adminUsecase) List(ctx context.Context, limit, offset uint) ([]entities.User, error) {
	users, err := a.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	return users, nil
}

func (a adminUsecase) GetBannedUsers(ctx context.Context, limit, offset uint) ([]entities.User, error) {
	users, err := a.repo.GetBannedUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get banned users: %w", err)
	}

	return users, nil
}

func (a adminUsecase) GetUser(ctx context.Context, id int64) (entities.User, error) {
	user, err := a.repo.GetUser(ctx, id)
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (a adminUsecase) DeleteUser(ctx context.Context, id int64) error {
	if err := a.repo.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func (a adminUsecase) SetAdmin(ctx context.Context, id int64) error {
	if isAdmin, err := a.repo.IsAdmin(ctx, id); err != nil {
		return fmt.Errorf("failed to check admin status: %w", err)
	} else if isAdmin {
		return nil
	}

	if err := a.SetAdmin(ctx, id); err != nil {
		return fmt.Errorf("failed to set admin: %w", err)
	}

	return nil
}

func (a adminUsecase) RevokeAdmin(ctx context.Context, id int64) error {
	if isAdmin, err := a.repo.IsAdmin(ctx, id); err != nil {
		return fmt.Errorf("failed to check admin status: %w", err)
	} else if !isAdmin {
		return nil
	}

	if err := a.RevokeAdmin(ctx, id); err != nil {
		return fmt.Errorf("failed to revoke admin: %w", err)
	}

	return nil
}

func (a adminUsecase) BanUser(ctx context.Context, id int64, description string) error {
	if isBanned, err := a.repo.IsBanned(ctx, id); err != nil {
		return fmt.Errorf("failed to check ban status: %w", err)
	} else if isBanned {
		return nil
	}

	if letters := utf8.RuneCountInString(description); letters < entities.MaxBanDescriptionLen {
		return entities.ErrLongInput
	}

	if err := a.repo.BanUser(ctx, id, description); err != nil {
		return fmt.Errorf("failed to ban user: %w", err)
	}

	return nil
}

func (a adminUsecase) UnbanUser(ctx context.Context, id int64) error {
	if isBanned, err := a.repo.IsBanned(ctx, id); err != nil {
		return fmt.Errorf("failed to check ban status: %w", err)
	} else if !isBanned {
		return nil
	}

	if err := a.repo.UnbanUser(ctx, id); err != nil {
		return fmt.Errorf("failed to unban user: %w", err)
	}

	return nil
}
