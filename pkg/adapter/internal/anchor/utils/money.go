package util

import (
	"errors"
	"math/big"
	"strings"
)

// REIT 9 dec → nano
func TokensToNano(s string) (int64, error) {
	n, err := toIntScaled(s, 0)
	if err != nil {
		return 0, err
	}
	return n * 1_000_000_000, nil
}

func USDCToMicro(s string) (int64, error) { return toIntScaled(s, 6) }
func MicroToUSDC(x int64) string          { return fromIntScaled(x, 6) }
func NanoToTokens(x int64) string         { return fromIntScaled(x, 9) }

// helpers
func toIntScaled(s string, dec int) (int64, error) {
	if s == "" {
		return 0, errors.New("empty")
	}
	parts := strings.SplitN(s, ".", 2)
	intp := parts[0]
	frac := ""
	if len(parts) == 2 {
		frac = parts[1]
	}
	if len(frac) > dec {
		frac = frac[:dec]
	}
	for len(frac) < dec {
		frac += "0"
	}
	bi, ok := new(big.Int).SetString(intp+frac, 10)
	if !ok || !bi.IsInt64() {
		return 0, errors.New("parse/overflow")
	}
	return bi.Int64(), nil
}
func fromIntScaled(x int64, dec int) string {
	sign := ""
	v := x
	if v < 0 {
		sign = "-"
		v = -v
	}
	pow := int64(1)
	for i := 0; i < dec; i++ {
		pow *= 10
	}
	hi := v / pow
	lo := v % pow
	frac := make([]byte, dec)
	for i := dec - 1; i >= 0; i-- {
		frac[i] = byte('0' + (lo % 10))
		lo /= 10
	}
	if dec == 0 {
		return sign + big.NewInt(hi).String()
	}
	return sign + big.NewInt(hi).String() + "." + string(frac)
}
