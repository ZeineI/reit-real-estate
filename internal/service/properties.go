package service

import (
	"context"
	"fmt"
	"reit-real-estate/internal/dto"
)

type propertyRepository interface {
	CreatProperty(ctx context.Context, dto *dto.CreatePropertyDTO) (string, error)
}

func (s *service) RegisterProperty(ctx context.Context, request *dto.RegisterPropertyDTO) error {
	if err := request.Validate(); err != nil {
		return fmt.Errorf("service.RegisterProperty.Validate error: %w", err)
	}

	userDTO, err := s.userRepository.GetUserByID(ctx, request.OwnerID)
	if err != nil {
		return fmt.Errorf("service.RegisterProperty.GetUserByID error: %w", err)
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

	return nil
}
