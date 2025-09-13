package wallets

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

func (r *Repository) CreatWallet(ctx context.Context, dto *dto.CreateWalletDTO) (string, error) {
	query := `INSERT INTO wallets (id, user_id, address) VALUES ($1, $2, $3)`

	newUUID := uuid.New()
	_, err := r.db.ExecContext(ctx, query, uuid.New(), dto.UserID, dto.WalletAddress)
	if err != nil {
		return "", fmt.Errorf("repository.wallets.CreatWallet error: %w", err)
	}

	return newUUID.String(), nil
}

func (r *Repository) GetWalletByUserID(ctx context.Context, id string) (*dto.WalletDTO, error) {
	query := `SELECT id, user_id, address, created_at FROM wallets WHERE user_id=$1`

	row := r.db.QueryRowContext(ctx, query, id)

	var model wallet
	if err := row.Scan(&model.id, &model.userID, &model.address, &model.createdAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("repository.users.GetUserByID error: %w", err)
	}

	walletDTO := &dto.WalletDTO{
		ID:            model.id.String(),
		UserID:        model.userID.String(),
		WalletAddress: model.address,
		CreatedAt:     model.createdAt.Unix(),
	}

	return walletDTO, nil
}
