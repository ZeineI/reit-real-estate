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

type WalletDTO struct {
	ID            string `json:"id"`
	WalletAddress string `json:"wallet_address"`
	UserID        string `json:"user_id"`
	CreatedAt     int64  `json:"created_at"`
}

type CreateWalletDTO struct {
	WalletAddress string `json:"wallet_address"`
	UserID        string `json:"user_id"`
}

type PropertyDTO struct {
	ID         string `json:"id"`
	OwnerID    string `json:"owner_id"`
	Name       string `json:"name"`
	TokenTotal int64  `json:"token_total"`
	CreatedAt  int64  `json:"created_at"`
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

type TokenDTO struct {
	ID         string `json:"id"`
	PropertyID string `json:"property_id"`
	Symbol     string `json:"symbol"`
	Price      int64  `json:"price"`
	CreatedAt  int64  `json:"created_at"`

	InvestedAmount int64
}

type CreateTokenDTO struct {
	PropertyID string `json:"property_id"`
	Symbol     string `json:"symbol"`
	Price      int64  `json:"price"`
}

type InvestDTO struct {
	InvestorID  string `json:"investor_id"`
	PropertyID  string `json:"property_id"`
	TokenAmount int64  `json:"token_amount"`
}

type CreateUserTokenDTO struct {
	InvestorID string `json:"investor_id"`
	TokenID    string `json:"token_id"`
	Quantity   int64  `json:"quantity"`
}
