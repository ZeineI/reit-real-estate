package tokens

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

func (r *Repository) CreateToken(ctx context.Context, dto *dto.CreateTokenDTO) (string, error) {
	query := `INSERT INTO tokens (id, property_id, symbol, price) VALUES ($1, $2, $3, $4)`

	newUUID := uuid.New()
	_, err := r.db.ExecContext(ctx, query, newUUID, dto.PropertyID, dto.Symbol, dto.Price)
	if err != nil {
		return "", fmt.Errorf("repository.tokens.CreateToken error: %w", err)
	}

	return newUUID.String(), nil
}

func (r *Repository) GetTokenByPropertyID(ctx context.Context, id string) (*dto.TokenDTO, error) {
	query := `
	SELECT 
    	id, 
    	property_id,
    	symbol,
    	price, 
    	created_at,
    	COALESCE((SELECT SUM(quantity) FROM user_tokens WHERE user_tokens.token_id = tokens.id), 0)
	FROM tokens WHERE property_id=$1`

	row := r.db.QueryRowContext(ctx, query, id)

	var (
		model          token
		investedAmount int64
	)
	if err := row.Scan(
		&model.id,
		&model.propertyID,
		&model.symbol,
		&model.price,
		&model.createdAt,
		&investedAmount,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("repository.tokens.GetTokenByPropertyID error: %w", err)
	}

	tokenDTO := &dto.TokenDTO{
		ID:             model.id.String(),
		PropertyID:     model.propertyID.String(),
		Symbol:         model.symbol,
		Price:          int64(model.price),
		CreatedAt:      model.createdAt.Unix(),
		InvestedAmount: investedAmount,
	}

	return tokenDTO, nil
}
