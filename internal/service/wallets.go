package service

import (
	"context"
	"reit-real-estate/internal/dto"
)

type walletRepository interface {
	CreatWallet(context.Context, *dto.CreateWalletDTO) (string, error)
}
