package yrand

import (
	"crypto/rand"
	"math/big"
)

func Base62(nBytes int) string {
	b := make([]byte, nBytes)
	n, err := rand.Read(b[:])
	if err != nil || n < nBytes {
		panic("failed to read random bytes")
	}
	var i big.Int
	i.SetBytes(b[:])
	return i.Text(62)
}
