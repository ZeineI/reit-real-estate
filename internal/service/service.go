package service

import (
	"context"
	"reit-real-estate/pkg/adapter/internal/pkg/reit"
)

type reitService interface {
	CreatePropertyTx(ctx context.Context, p reit.CreatePropertyParams) (reit.TxBase64, error)
}

type service struct {
	userRepository      userRepository
	walletRepository    walletRepository
	propertyRepository  propertyRepository
	tokenRepository     tokenRepository
	userTokenRepository userTokenRepository
	reitService         reitService
	reitConfig          reitConfig
}

type reitConfig struct {
	rpcURL       string
	tokenAddress string
	reitMint     string
}

func NewService(
	userRepository userRepository,
	walletRepository walletRepository,
	propertyRepository propertyRepository,
	tokenRepository tokenRepository,
	userTokenRepository userTokenRepository,
	reitService reitService,
) *service {
	return &service{
		userRepository:      userRepository,
		walletRepository:    walletRepository,
		propertyRepository:  propertyRepository,
		tokenRepository:     tokenRepository,
		userTokenRepository: userTokenRepository,
		reitService:         reitService,
	}
}

func (s *service) WithReit(rpcUrl, tokenAddress, reitMint string) {
	s.reitConfig.rpcURL = rpcUrl
	s.reitConfig.tokenAddress = tokenAddress
	s.reitConfig.reitMint = reitMint
}
