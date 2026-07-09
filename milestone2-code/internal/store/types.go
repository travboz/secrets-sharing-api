package store

import "errors"

var (
	ErrFileExists  = errors.New("File already exists")
	ErrKeyNotFound = errors.New("Key does not exist")
)

type SecretData struct {
	Id     string
	Secret string
}
