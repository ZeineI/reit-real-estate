package service

import (
	"context"
	"fmt"
	"log"
	"reit-real-estate/internal/dto"
	"reit-real-estate/pkg/adapter/internal/pkg/reit"
)

type propertyRepository interface {
	CreatProperty(ctx context.Context, dto *dto.CreatePropertyDTO) (string, error)
	GetPropertyByID(ctx context.Context, id string) (*dto.PropertyDTO, error)
}

func (s *service) RegisterProperty(ctx context.Context, request *dto.RegisterPropertyDTO) error {
	if err := request.Validate(); err != nil {
		return fmt.Errorf("service.RegisterProperty.Validate error: %w", err)
	}

	userDTO, err := s.userRepository.GetUserByID(ctx, request.OwnerID)
	if err != nil {
		return fmt.Errorf("service.RegisterProperty.GetUserByID error: %w", err)
	}
	if userDTO.Role != dto.RoleAdmin {
		return fmt.Errorf("service.RegisterProperty error: %w", dto.ErrInvalidRole)
	}

	walletDTO, err := s.walletRepository.GetWalletByUserID(ctx, request.OwnerID)
	if err != nil {
		return fmt.Errorf("service.RegisterProperty.GetWalletByUserID error: %w", err)
	}

	propertyID, err := s.propertyRepository.CreatProperty(ctx, &dto.CreatePropertyDTO{
		OwnerID:    userDTO.ID,
		Name:       request.Name,
		TokenTotal: request.TokenTotal,
	})
	if err != nil {
		return fmt.Errorf("service.RegisterProperty.CreatProperty error: %w", err)
	}

	_, err = s.tokenRepository.CreateToken(ctx, &dto.CreateTokenDTO{
		PropertyID: propertyID,
		Symbol:     request.Symbol,
		Price:      request.Price,
	})
	if err != nil {
		return fmt.Errorf("service.RegisterProperty.CreateToken error: %w", err)
	}

	transaction, err := s.reitService.CreatePropertyTx(ctx, reit.CreatePropertyParams{
		WalletAddress:     walletDTO.WalletAddress,
		TokenAddress:      s.reitConfig.tokenAddress,
		Price:             request.Price,
		TotalSupplyTokens: fmt.Sprintf("%d", request.TokenTotal),
		ReitMint:          s.reitConfig.reitMint,
		Symbol:            request.Symbol,
	})
	if err != nil {
		return fmt.Errorf("service.RegisterProperty.CreatePropertyTx error: %w", err)
	}
	log.Println(transaction)

	return nil
}

func (s *service) Invest(ctx context.Context, request *dto.InvestDTO) error {
	if err := request.Validate(); err != nil {
		return fmt.Errorf("service.RegisterProperty.Validate error: %w", err)
	}

	userDTO, err := s.userRepository.GetUserByID(ctx, request.InvestorID)
	if err != nil {
		return fmt.Errorf("service.RegisterProperty.GetUserByID error: %w", err)
	}
	if userDTO.Role != dto.RoleInvestor {
		return fmt.Errorf("service.RegisterProperty error: %w", dto.ErrInvalidRole)
	}

	propertyDTO, err := s.propertyRepository.GetPropertyByID(ctx, request.PropertyID)
	if err != nil {
		return fmt.Errorf("service.RegisterProperty.GetPropertyByID error: %w", err)
	}
	if propertyDTO.TokenTotal < request.TokenAmount {
		return fmt.Errorf("service.RegisterProperty error: %w", dto.ErrInvalidTokenAmount)
	}

	tokenDTO, err := s.tokenRepository.GetTokenByPropertyID(ctx, propertyDTO.ID)
	if err != nil {
		return fmt.Errorf("service.RegisterProperty.GetTokenByPropertyID error: %w", err)
	}

	if propertyDTO.TokenTotal-tokenDTO.InvestedAmount < request.TokenAmount {
		return fmt.Errorf("service.RegisterProperty error: %w", dto.ErrInvalidTokenAmount)
	}

	_, err = s.userTokenRepository.CreatUserToken(ctx, &dto.CreateUserTokenDTO{
		InvestorID: userDTO.ID,
		TokenID:    tokenDTO.ID,
		Quantity:   request.TokenAmount,
	})
	if err != nil {
		return fmt.Errorf("service.RegisterProperty.CreatUserToken error: %w", err)
	}

	return nil
}
