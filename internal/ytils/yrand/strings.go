package yrand

import (
	"crypto/rand"
	"math/big"
	"sync"
)

var randMutex = sync.Mutex{}

func Base62(nBytes int) string {
	b := make([]byte, nBytes)
	randMutex.Lock()
	n, err := rand.Read(b[:])
	randMutex.Unlock()

	if err != nil || n < nBytes {
		panic("failed to read random bytes")
	}
	var i big.Int
	i.SetBytes(b[:])
	return i.Text(62)
}
