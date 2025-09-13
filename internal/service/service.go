package service

type service struct {
	userRepository     userRepository
	walletRepository   walletRepository
	propertyRepository propertyRepository
	tokenRepository    tokenRepository
}

func NewService(
	userRepository userRepository,
	walletRepository walletRepository,
	propertyRepository propertyRepository,
	tokenRepository tokenRepository,
) *service {
	return &service{
		userRepository:     userRepository,
		walletRepository:   walletRepository,
		propertyRepository: propertyRepository,
		tokenRepository:    tokenRepository,
	}
}
