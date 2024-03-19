package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/skantay/task-management-system/internal/entities"
)

type adminRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) adminRepository {
	return adminRepository{db}
}

func (a adminRepository) GetUsers(ctx context.Context, limit, offset uint) ([]entities.User, error) {
	tx, err := a.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin a transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	users := make([]entities.User, 0, limit)

	query := `SELECT id, name, username, email, is_admin
				FROM users
				LIMIT ?
				OFFSET ?`

	if err := tx.SelectContext(
		ctx,
		&users,
		query,
		limit,
		offset,
	); err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return users, nil
}

func (a adminRepository) GetUser(ctx context.Context, id int64) (entities.User, error) {
	tx, err := a.db.BeginTxx(ctx, nil)
	if err != nil {
		return entities.User{}, fmt.Errorf("failed to begin a transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var user entities.User

	query := `SELECT id, name, username, email, is_admin
				FROM users
				WHERE id = ?`

	if err := tx.SelectContext(
		ctx,
		&user,
		query,
		id,
	); err != nil {
		return entities.User{}, fmt.Errorf("failed to get a user: %w", err)
	}

	return entities.User{}, nil
}

func (a adminRepository) BanUser(ctx context.Context, id int64, description string) error {
	tx, err := a.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin a transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	query := `INSERT INTO banned_users(user_id, description)
				VALUES(?, ?)`

	if _, err := tx.ExecContext(ctx, query, id, description); err != nil {
		return fmt.Errorf("failed to insert into table: %w", err)
	}

	return nil
}

func (a adminRepository) GetBannedUsers(ctx context.Context) ([]entities.BannedUser, error) {
	return nil, nil
}

func (a adminRepository) GetBannedUser(ctx context.Context) (entities.BannedUser, error) {
	return entities.BannedUser{}, nil
}

func (a adminRepository) UnbanUser(ctx context.Context, id int64) error {
	return nil
}

func (a adminRepository) DeleteUser(ctx context.Context, id int64, message string) error {
	return nil
}

func (a adminRepository) SetAdmin(ctx context.Context, id int64) error {
	return nil
}

func (a adminRepository) RemoveAdmin(ctx context.Context, id int64) error {
	return nil
}
