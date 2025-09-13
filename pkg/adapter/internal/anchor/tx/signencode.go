package tx

import (
	"encoding/base64"

	gsol "github.com/gagliardetto/solana-go"
)

func EncodeMessageBase64(msg gsol.Message) (string, error) {
	bin, err := msg.MarshalBinary()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bin), nil
}

func EncodeTransactionBase64(tx *gsol.Transaction) (string, error) {
	bin, err := tx.MarshalBinary()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bin), nil
}
