package ycrypto

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/argon2"
)

var (
	fixedSalt = "64ca378646c4e2887a5e594deb7b56a85fc0337d303f91f7de89c0ae56b75f8992e3e1757e1514cff3c3cdf3ac0fc49b3d147631b30b34a003d4032f7455db75"
)

func HashPassword(password string, salt []byte) (string, error) {
	if salt == nil {
		s, err := hex.DecodeString(fixedSalt)
		if err != nil {
			return "", err
		}
		salt = s
	}

	key := argon2.IDKey([]byte(password), salt, 2, 32*1024, 2, 32)
	result := fmt.Sprintf("argon2id:%s:%s", base64.StdEncoding.EncodeToString(salt), base64.StdEncoding.EncodeToString(key))

	return result, nil
}
