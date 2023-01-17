package auth

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

type AccountNumberService interface {
	Random() string
	Validate(string) bool
}

type accountNumberService struct{}

func NewAccountNumberService() AccountNumberService {
	return &accountNumberService{}
}

func (a *accountNumberService) Random() string {
	n, err := rand.Int(rand.Reader, big.NewInt(8999_9999_9999_999))
	if err != nil {
		return ""
	}
	an := n.Uint64() + 1000_0000_0000_000

	checkDigit := checksum(an)
	if checkDigit != 0 {
		checkDigit = 10 - checkDigit
	}
	an = an*10 + checkDigit

	return strconv.FormatUint(an, 10)
}

func (a *accountNumberService) Validate(number string) bool {
	n, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		return false
	}
	return isValid(n)
}

func isValid(number uint64) bool {
	return (number%10+checksum(number/10))%10 == 0
}

func checksum(number uint64) uint64 {
	var l uint64
	for i := uint64(0); number > 0; i++ {
		cur := number % 10
		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}
		l += cur
		number = number / 10
	}
	return l % 10
}
