package ycrypto

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHashPassword(t *testing.T) {
	v, err := HashPassword("password", nil)
	require.NoError(t, err)
	assert.Equal(t, `argon2id:ZMo3hkbE4oh6XllN63tWqF/AM30wP5H33onArla3X4mS4+F1fhUUz/PDzfOsD8SbPRR2MbMLNKAD1AMvdFXbdQ==:ZQQn+GINP4IpTOGTLUquDPawPYXoMJCasiJgUOlnlGI=`, v)
}
