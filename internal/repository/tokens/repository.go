package tokens

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

func (r *Repository) CreateToken(ctx context.Context, dto *dto.CreateTokenDTO) (string, error) {
	query := `INSERT INTO tokens (id, property_id, symbol, price) VALUES ($1, $2, $3, $4)`

	newUUID := uuid.New()
	_, err := r.db.ExecContext(ctx, query, newUUID, dto.PropertyID, dto.Symbol, dto.Price)
	if err != nil {
		return "", fmt.Errorf("repository.tokens.CreateToken error: %w", err)
	}

	return newUUID.String(), nil
}
