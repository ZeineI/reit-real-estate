package dto

type CreatePropertyAdapter struct {
	WalletAddress     string // кто создаёт объект (signer) адрес
	TokenAddress      string // адрес токена на который поступают арендные платежи  test-netовский токен адрес
	Price             int64  //TODO float64
	TotalSupplyTokens string //  макс долей
	Symbol            string //  Sabina_token только название
}
