package cryptoconfig

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func New(password, salt string) (*CryptoConfig, error) {
	// Derive key from password and salt
	if password == "" || salt == "" {
		return nil, ErrKeyMaterialMissing
	}

	key, err := deriveKey([]byte(password), []byte(salt))
	if err != nil {
		return nil, fmt.Errorf("(InitCrypto) Error deriving key: %w", err)
	}

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("(InitCrypto) Error creating new cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, fmt.Errorf("(InitCrypto) Error creating new gcm: %w", err)
	}

	// 'number once used' -> non-repeated data that's only used once with a certain key.
	// Don't repeat the same combination of key + nonce.
	nonce := make([]byte, gcm.NonceSize())     // So, we make a byte slice to hold our nonce (which equals the size of the GCM's nonce created above).
	if _, err = rand.Read(nonce); err != nil { // We then fill it with random values using rand.Read().
		return nil, fmt.Errorf("(InitCrypto) Error reading random bytes into nonce: %w", err)
	}

	return &CryptoConfig{
		gcm:   gcm,
		nonce: nonce,
	}, nil
}

// deriveKey stretches our password to make it suitable as a cryptographic key.
func deriveKey(password, salt []byte) ([]byte, error) {
	key, err := scrypt.Key(password, salt, CRYPTO_COST, 8, 1, OUTPUT_KEY_LENGTH_IN_BYTES)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// Encrypt enciphers our data using a key created by passing in our password (key arg).
func (c *CryptoConfig) Encrypt(plaintext string) (ciphertext []byte) {
	ciphertext = c.gcm.Seal(c.nonce, c.nonce, []byte(plaintext), nil)
	return ciphertext
}

// Decrypt deenciphers the data using the password we've chosen.
func (c *CryptoConfig) Decrypt(ciphertext []byte) (plaintext []byte, err error) {
	nonce := ciphertext[:c.gcm.NonceSize()]
	ciphertext = ciphertext[c.gcm.NonceSize():]

	plaintext, err = c.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("Error decrypting data: %w", err)
	}

	return plaintext, err
}
