package main

import (
	"slices"

	"github.com/travboz/secrets-sharing/milestone3/internal/store"
)

// Selecting to mock but use a MockFileStore which mimics behaviour.
/*
This MockFileStore has *exactly the same* functions that are used by
the handlers in the test. We control the input and output.
Where the handler calls `s.Write(blah)`, it's instead going to call
our MockFileStore.Write(blah).
How cool is that?!
*/
type MockFileStore struct {
	ReadFunc  func(key string) (string, error)
	WriteFunc func(data store.SecretData) error
}

func (m *MockFileStore) Read(key string) (string, error) {
	return m.ReadFunc(key)
}

func (m *MockFileStore) Write(data store.SecretData) error {
	return m.WriteFunc(data)
}

/*
Example of the ReadFunc mock function:
	mock := &MockFileStore{
		ReadFunc: func(key string) (string, error) {
			# This function runs when handler calls mock.Read()
			return "my-secret-value", nil
		},
	}
*/

/*
Process:
1. Handler gets called once with valid id, returns secret.
2. Handlers gets called again with same id from (1), returns store.ErrKeyNotFound because we're
	expecting the store to delete an item after its viewed once.

So, need to:
- mock the store again
- use a spy to track how many calls were made to the store using that id
- if id has already been called, return store.ErrKeyNotFound
*/

type SpyFileStore struct {
	ReadFunc  func(key string) (string, error)
	WriteFunc func(data store.SecretData) error
	KeysSeen  []string
}

func (m *SpyFileStore) Read(key string) (string, error) {
	// New key/id, haven't seen this before
	if !slices.Contains(m.KeysSeen, key) {
		m.AppendKey(key)       // Add key to seen
		return m.ReadFunc(key) // Return secret as expected
	}

	// We've already seen this key/id, i.e. the GET request has already been made for this key.
	// Secrets should be seen once, and so we return an error.
	return "", store.ErrKeyNotFound
}

func (m *SpyFileStore) Write(data store.SecretData) error {
	return m.WriteFunc(data)
}

func (m *SpyFileStore) AppendKey(id string) {
	m.KeysSeen = append(m.KeysSeen, id)
}
