package properties

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

func (r *Repository) CreatProperty(ctx context.Context, dto *dto.CreatePropertyDTO) (string, error) {
	query := `INSERT INTO properties (id, owner_id, name, token_total) VALUES ($1, $2, $3, $4)`

	newUUID := uuid.New()
	_, err := r.db.ExecContext(ctx, query, newUUID, dto.OwnerID, dto.Name, dto.TokenTotal)
	if err != nil {
		return "", fmt.Errorf("repository.properties.CreatProperty error: %w", err)
	}

	return newUUID.String(), nil
}
