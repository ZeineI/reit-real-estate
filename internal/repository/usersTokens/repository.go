package usersTokens

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

func (r *Repository) CreatUserToken(ctx context.Context, dto *dto.CreateUserTokenDTO) (string, error) {
	query := `INSERT INTO user_tokens (id, user_id, token_id, quantity) VALUES ($1, $2, $3, $4)`

	newUUID := uuid.New()
	_, err := r.db.ExecContext(ctx, query, newUUID, dto.InvestorID, dto.TokenID, dto.Quantity)
	if err != nil {
		return "", fmt.Errorf("repository.usersTokens.CreatUserToken error: %w", err)
	}

	return newUUID.String(), nil
}
