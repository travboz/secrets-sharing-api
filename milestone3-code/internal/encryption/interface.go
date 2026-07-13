package encryption

type Encrypter interface {
	// Encrypt enciphers our data using a key created by passing in our password (key arg).
	Encrypt(plaintext string) (ciphertext []byte)
	// Decrypt deenciphers the data using the password we've chosen.
	Decrypt(ciphertext []byte) (plaintext []byte, err error)
}
