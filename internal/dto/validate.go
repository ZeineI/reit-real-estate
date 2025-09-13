package dto

import "strings"

func (dto *RegisterUserDTO) Validate() error {
	dto.normalize()
	if dto.Role != roleInvestor && dto.Role != roleAdmin {
		return ErrInvalidRole
	}

	if dto.WalletAddress == "" {
		return ErrRequiredWalletAddress
	}

	return nil
}

func (dto *RegisterUserDTO) normalize() {
	dto.Login = strings.TrimSpace(dto.Login)
	dto.Role = strings.TrimSpace(dto.Role)
	dto.WalletAddress = strings.TrimSpace(dto.WalletAddress)
}
