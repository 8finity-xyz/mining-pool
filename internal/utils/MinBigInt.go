package utils

import "math/big"

func MinBigInt(a *big.Int, b *big.Int) *big.Int {
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}
