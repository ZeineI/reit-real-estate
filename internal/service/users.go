package service

import (
	"context"
	"fmt"
	"reit-real-estate/internal/dto"
)

type userRepository interface {
	CreatUser(context.Context, *dto.CreateUserDTO) error
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

	if err := us.userRepository.CreatUser(ctx, &dto.CreateUserDTO{Login: request.Login, Role: request.Role}); err != nil {
		return fmt.Errorf("service.users.RegisterUser.CreatUser error: %w", err)
	}

	if err := us.walletRepository.CreatWallet(ctx, &dto.CreateWalletDTO{
		WalletAddress: request.WalletAddress,
		UserID:        "",
	}); err != nil {
		return fmt.Errorf("service.users.RegisterUser.CreatWallet error: %w", err)
	}

	return nil
}
