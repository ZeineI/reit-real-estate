package service

import (
	"context"
	"fmt"
	"reit-real-estate/internal/dto"
)

type userRepository interface {
	CreatUser(ctx context.Context, dto *dto.CreateUserDTO) (string, error)
}

type userService struct {
	userRepository   userRepository
	walletRepository walletRepository
}

func NewUserService(userRepository userRepository, walletRepository walletRepository) *userService {
	return &userService{
		userRepository:   userRepository,
		walletRepository: walletRepository,
	}
}

func (us *userService) RegisterUser(ctx context.Context, request *dto.RegisterUserDTO) error {
	if err := request.Validate(); err != nil {
		return fmt.Errorf("service.users.RegisterUser.Validate error: %w", err)
	}

	//TODO validate user login is unique

	userID, err := us.userRepository.CreatUser(ctx, &dto.CreateUserDTO{Login: request.Login, Role: request.Role})
	if err != nil {
		return fmt.Errorf("service.users.RegisterUser.CreatUser error: %w", err)
	}

	_, err = us.walletRepository.CreatWallet(ctx, &dto.CreateWalletDTO{
		WalletAddress: request.WalletAddress,
		UserID:        userID,
	})
	if err != nil {
		return fmt.Errorf("service.users.RegisterUser.CreatWallet error: %w", err)
	}

	return nil
}
