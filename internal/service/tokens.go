package service

import (
	"context"
	"reit-real-estate/internal/dto"
)

type tokenRepository interface {
	CreateToken(ctx context.Context, dto *dto.CreateTokenDTO) (string, error)
	GetTokenByPropertyID(ctx context.Context, id string) (*dto.TokenDTO, error)
}

type userTokenRepository interface {
	CreatUserToken(ctx context.Context, dto *dto.CreateUserTokenDTO) (string, error)
}
