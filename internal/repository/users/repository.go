package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"reit-real-estate/internal/dto"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreatUser(ctx context.Context, dto *dto.CreateUserDTO) (string, error) {
	query := `INSERT INTO users (id, login, role) VALUES ($1, $2, $3)`

	newUUID := uuid.New()
	_, err := r.db.ExecContext(ctx, query, newUUID, dto.Login, dto.Role)
	if err != nil {
		return "", fmt.Errorf("repository.users.CreatUser error: %w", err)
	}

	return newUUID.String(), nil
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*dto.UserDTO, error) {
	query := `SELECT id, login, role, created_at FROM users WHERE id=$1`

	row := r.db.QueryRowContext(ctx, query, id)

	var model user
	if err := row.Scan(&model.id, &model.login, &model.role, &model.createdAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("repository.users.GetUserByID error: %w", err)
	}

	userDTO := &dto.UserDTO{
		ID:        model.id.String(),
		Login:     model.login,
		Role:      model.role,
		CreatedAt: model.createdAt.Unix(),
	}

	return userDTO, nil
}
