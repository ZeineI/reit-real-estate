package properties

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

func (r *Repository) CreatProperty(ctx context.Context, dto *dto.CreatePropertyDTO) (string, error) {
	query := `INSERT INTO properties (id, owner_id, name, token_total) VALUES ($1, $2, $3, $4)`

	newUUID := uuid.New()
	_, err := r.db.ExecContext(ctx, query, newUUID, dto.OwnerID, dto.Name, dto.TokenTotal)
	if err != nil {
		return "", fmt.Errorf("repository.properties.CreatProperty error: %w", err)
	}

	return newUUID.String(), nil
}

func (r *Repository) GetPropertyByID(ctx context.Context, id string) (*dto.PropertyDTO, error) {
	query := `SELECT id, owner_id, name, token_total, created_at FROM properties WHERE id=$1`

	row := r.db.QueryRowContext(ctx, query, id)

	var model property
	if err := row.Scan(
		&model.id,
		&model.ownerID,
		&model.name,
		&model.tokenTotal,
		&model.createdAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("repository.properties.GetPropertyByID error: %w", err)
	}

	propertyDTO := &dto.PropertyDTO{
		ID:         model.id.String(),
		OwnerID:    model.ownerID.String(),
		Name:       model.name,
		TokenTotal: model.tokenTotal,
		CreatedAt:  model.createdAt.Unix(),
	}

	return propertyDTO, nil
}
