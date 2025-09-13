package dto

import "errors"

var (
	ErrInvalidRole           = errors.New("role is invalid")
	ErrRequiredWalletAddress = errors.New("wallet_address is required")
)

const (
	roleAdmin    string = "admin"
	roleInvestor string = "investor"
)
