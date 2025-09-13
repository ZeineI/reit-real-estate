package dto

type UserDTO struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	Role      string `json:"role"`
	CreatedAt int64  `json:"created_at"`
}

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

type RegisterPropertyDTO struct {
	OwnerID    string `json:"owner_id"`
	Name       string `json:"name"`
	TokenTotal int64  `json:"token_total"`
	Symbol     string `json:"symbol"`
	Price      int64  `json:"price"`
}

type CreatePropertyDTO struct {
	OwnerID    string `json:"owner_id"`
	Name       string `json:"name"`
	TokenTotal int64  `json:"token_total"`
}

type CreateTokenDTO struct {
	PropertyID string `json:"property_id"`
	Symbol     string `json:"symbol"`
	Price      int64  `json:"price"`
}
