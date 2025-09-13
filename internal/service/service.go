package service

type service struct {
	userRepository      userRepository
	walletRepository    walletRepository
	propertyRepository  propertyRepository
	tokenRepository     tokenRepository
	userTokenRepository userTokenRepository
}

func NewService(
	userRepository userRepository,
	walletRepository walletRepository,
	propertyRepository propertyRepository,
	tokenRepository tokenRepository,
	userTokenRepository userTokenRepository,
) *service {
	return &service{
		userRepository:      userRepository,
		walletRepository:    walletRepository,
		propertyRepository:  propertyRepository,
		tokenRepository:     tokenRepository,
		userTokenRepository: userTokenRepository,
	}
}
