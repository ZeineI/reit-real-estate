package solana

type Pubkey string
type Signature string

type RawMessage struct {
	Bytes []byte
}

type rpcRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params,omitempty"`
}

type rpcResponse[T any] struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  T      `json:"result"`
	Error   *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

type getLatestBlockhashResult struct {
	Value struct {
		Blockhash            string `json:"blockhash"`
		LastValidBlockHeight uint64 `json:"lastValidBlockHeight"`
	} `json:"value"`
}
