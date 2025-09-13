package service

import (
	"context"
	"reit-real-estate/internal/dto"
)

type tokenRepository interface {
	CreateToken(ctx context.Context, dto *dto.CreateTokenDTO) (string, error)
}
