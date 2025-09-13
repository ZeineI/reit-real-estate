package users

import (
	"context"
	"database/sql"
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
