package anchor

import "encoding/json"

type IDL struct {
	Name string `json:"name"`
	// ... остальное по Anchor IDL
}

func LoadIDL(raw []byte) (IDL, error) {
	var idl IDL
	err := json.Unmarshal(raw, &idl)
	return idl, err
}
