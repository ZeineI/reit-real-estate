package wallets

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

func (r *Repository) CreatWallet(ctx context.Context, dto *dto.CreateWalletDTO) error {
	query := `INSERT INTO wallets (id, user_id, address) VALUES ($1, $2, $3)`

	_, err := r.db.ExecContext(ctx, query, uuid.New(), dto.UserID, dto.WalletAddress)
	if err != nil {
		return fmt.Errorf("repository.wallets.CreatWallet error: %w", err)
	}

	return nil
}
