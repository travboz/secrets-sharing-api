package store

type Store interface {
	Write(data SecretData) error
	Read(key string) (string, error)
}
