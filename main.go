package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func randBase62() string {
	var buf [12]byte
	_, _ = rand.Read(buf[:])
	var i big.Int
	i.SetBytes(buf[:])
	return i.Text(62)
}

func nowBase62() string {
	var i big.Int
	i.SetInt64(time.Now().UnixMilli())
	return i.Text(62)
}

func genId() string {
	return fmt.Sprintf("%s_%s", nowBase62(), randBase62())
}

func main() {
	fmt.Println(genId())
}
