package reit

// TxBase64 — сериализованная транзакция (без подписи).
type TxBase64 struct {
	Tx string `json:"tx_base64"`
}

// CreatePropertyParams — атомное создание property:
// 1) создать REIT-mint (9 dec), authority = PDA
// 2) заминтить ВЕСЬ фиксированный supply админу
// 3) set_authority(MintTokens) = None
// 4) инициализировать state + казну под TokenAddress mint (USDC-подобный)
type CreatePropertyParams struct {
	WalletAddress     string // admin signer pubkey (fee payer)
	TokenAddress      string // payout mint (USDC на devnet/local)
	Price             int64  // оффчейн (не нужен ончейн-инструкции)
	TotalSupplyTokens string // "1000" → внутри *1e9 = nano
	ReitMint          string // адрес смарт-кошелька-минта
	Symbol            string // оффчейн (тикер)
}

type PropertyState struct {
	ReitMint        string `json:"reit_mint"`
	TotalSupplyNano int64  `json:"total_supply_nano"`
	IndexHi         uint64 `json:"accrual_index_hi"`
	IndexLo         uint64 `json:"accrual_index_lo"`
	TreasuryATA     string `json:"treasury_ata"` // ← адрес «копилки» из ончейна
}
