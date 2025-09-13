package service

import (
	"context"
	"fmt"
	"reit-real-estate/internal/dto"
)

type userRepository interface {
	CreatUser(ctx context.Context, dto *dto.CreateUserDTO) (string, error)
	GetUserByID(ctx context.Context, id string) (*dto.UserDTO, error)
	GetUserByLogin(ctx context.Context, login string) (*dto.UserDTO, error)
}

func (s *service) RegisterUser(ctx context.Context, request *dto.RegisterUserDTO) error {
	if err := request.Validate(); err != nil {
		return fmt.Errorf("service.RegisterUser.Validate error: %w", err)
	}

	//TODO validate user login is unique

	userID, err := s.userRepository.CreatUser(ctx, &dto.CreateUserDTO{Login: request.Login, Role: request.Role})
	if err != nil {
		return fmt.Errorf("service.RegisterUser.CreatUser error: %w", err)
	}

	_, err = s.walletRepository.CreatWallet(ctx, &dto.CreateWalletDTO{
		WalletAddress: request.WalletAddress,
		UserID:        userID,
	})
	if err != nil {
		return fmt.Errorf("service.RegisterUser.CreatWallet error: %w", err)
	}

	return nil
}
