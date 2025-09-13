package dto

type RegisterUserDTO struct {
	Login         string `json:"login"`
	Role          string `json:"role"`
	WalletAddress string `json:"wallet_address"`
}

type CreateUserDTO struct {
	Login string `json:"login"`
	Role  string `json:"role"`
}

type CreateWalletDTO struct {
	WalletAddress string `json:"wallet_address"`
	UserID        string `json:"user_id"`
}

type CreatePropertyDTO struct {
	OwnerID    string `json:"owner_id"`
	Name       string `json:"name"`
	TokenTotal int64  `json:"token_total"`
	Symbol     string `json:"symbol"`
	Price      int64  `json:"price"`
}
