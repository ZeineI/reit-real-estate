package dto

import "strings"

func (dto *RegisterUserDTO) Validate() error {
	dto.normalize()
	if dto.Role != RoleAdmin && dto.Role != RoleInvestor {
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

func (dto *RegisterPropertyDTO) Validate() error {
	dto.normalize()
	if dto.OwnerID == "" {
		return ErrRequiredOwnerID
	}
	if dto.TokenTotal <= 0 {
		return ErrInvalidTokenTotal
	}
	if dto.Price <= 0 {
		return ErrInvalidPrice
	}

	return nil
}

func (dto *RegisterPropertyDTO) normalize() {
	dto.OwnerID = strings.TrimSpace(dto.OwnerID)
	dto.Name = strings.TrimSpace(dto.Name)
	dto.Symbol = strings.TrimSpace(dto.Symbol)
}

func (dto *InvestDTO) Validate() error {
	dto.normalize()
	if dto.InvestorID == "" {
		return ErrRequiredInvestorID
	}
	if dto.PropertyID == "" {
		return ErrRequiredPropertyID
	}
	if dto.TokenAmount <= 0 {
		return ErrInvalidTokenAmount
	}

	return nil
}

func (dto *InvestDTO) normalize() {
	dto.InvestorID = strings.TrimSpace(dto.InvestorID)
	dto.PropertyID = strings.TrimSpace(dto.PropertyID)
}
