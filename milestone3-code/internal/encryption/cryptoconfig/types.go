package cryptoconfig

import (
	"crypto/cipher"
	"errors"
)

const (
	SALT_LENGTH                = 32
	CRYPTO_COST                = 1048576
	OUTPUT_KEY_LENGTH_IN_BYTES = 32
)

var (
	ErrKeyMaterialMissing = errors.New("Cannot create key with empty password or salt")
)

type CryptoConfig struct {
	gcm   cipher.AEAD
	nonce []byte
}
