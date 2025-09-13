package dto

import "errors"

var (
	ErrInvalidRole           = errors.New("role is invalid")
	ErrInvalidTokenTotal     = errors.New("token_total is invalid")
	ErrInvalidTokenAmount    = errors.New("token_amount is invalid")
	ErrInvalidPrice          = errors.New("price is invalid")
	ErrRequiredWalletAddress = errors.New("wallet_address is required")
	ErrRequiredOwnerID       = errors.New("owner_id is required")
	ErrRequiredInvestorID    = errors.New("investor_id is required")
	ErrRequiredPropertyID    = errors.New("property_id is required")
)

const (
	RoleAdmin    string = "admin"
	RoleInvestor string = "investor"
)
