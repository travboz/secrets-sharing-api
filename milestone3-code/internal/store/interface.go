package store

type Store interface {
	// Writes the SecretData to the storage.
	Write(data SecretData) error
	// Fetches a Secret by its ID from the store.
	Read(id string) (string, error)
}
