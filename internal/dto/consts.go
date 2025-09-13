package dto

import "errors"

var (
	ErrInvalidRole           = errors.New("role is invalid")
	ErrInvalidTokenTotal     = errors.New("token_total is invalid")
	ErrInvalidPrice          = errors.New("price is invalid")
	ErrRequiredWalletAddress = errors.New("wallet_address is required")
	ErrRequiredOwnerID       = errors.New("owner_id is required")
)

const (
	roleAdmin    string = "admin"
	roleInvestor string = "investor"
)
