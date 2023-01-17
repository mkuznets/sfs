package auth

import (
	"crypto/rand"
	"math/big"
)

func RandomAccountNumber() string {
	var upper big.Int
	upper.SetString("90000000000000000000", 10)

	var adder big.Int
	adder.SetString("10000000000000000000", 10)

	n, err := rand.Int(rand.Reader, &upper)
	if err != nil {
		return ""
	}
	n.Add(n, &adder)

	return n.String()
}
