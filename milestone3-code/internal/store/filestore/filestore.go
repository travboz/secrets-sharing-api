package filestore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/travboz/secrets-sharing/milestone3/internal/encryption"
	"github.com/travboz/secrets-sharing/milestone3/internal/encryption/cryptoconfig"
	"github.com/travboz/secrets-sharing/milestone3/internal/store"
)

var (
	ErrFileExists  = errors.New("File already exists")
	ErrKeyNotFound = errors.New("Key does not exist")
)

// A concurrency-safe file based store.
type FileStore struct {
	mu           sync.Mutex
	DataFileName string
	Store        map[string][]byte `json:"store"`
	cryptoCfg    encryption.Encrypter
}

func (s *FileStore) Write(data store.SecretData) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.readFromFile(s.DataFileName)
	if err != nil {
		return err
	}

	// s.Store[data.Id] = data.Secret
	// To more closely align with the full solution, encrypt only the secret's value, not the entire file.
	s.Store[data.Id] = s.cryptoCfg.Encrypt(data.Secret)

	return s.writeToFile()
}

// writeToFile is a helper function which essentially writes the Store object to disk.
// Truncates file if it exists, so we are wriitng the complete Store object every time.
func (s *FileStore) writeToFile() error {
	var f *os.File

	jsonData, err := json.Marshal(s.Store)
	if err != nil {
		return fmt.Errorf("error marshalling data: %w", err)
	}

	// No longer encrypting entire file - just secret value.
	// ciphertext := s.cryptoCfg.Encrypt(string(jsonData))

	f, err = os.Create(s.DataFileName)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error writing store data to file: %w", err)
	}

	return nil
}

// Read fetches the item with the id of key. path is the filepath of the Store.
func (s *FileStore) Read(id string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.readFromFile(s.DataFileName)
	if err != nil {
		return "", fmt.Errorf("error reading from file: %w", err)
	}

	ciphertext, ok := s.Store[id]
	if !ok {
		return "", ErrKeyNotFound
	}

	data, err := s.cryptoCfg.Decrypt(ciphertext)
	if err != nil {
		return "", fmt.Errorf("error decrypting data: %w", err)

	}

	delete(s.Store, id)

	err = s.writeToFile()
	if err != nil {
		return "", fmt.Errorf("error writing file after read: %w", err)
	}

	return string(data), nil
}

// readFromFile is a helper which reads the complete file and overwrites
// the current Store object in memory.
func (s *FileStore) readFromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening file for read: %w", err)
	}
	defer f.Close()

	ciphertext, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("error reading data from file: %w", err)
	}

	if len(ciphertext) == 0 {
		return nil
	}

	return json.Unmarshal(ciphertext, &s.Store)
}

// Init creates the file if it doesn't exist, and initialises the FileStoreConfig
// to be used in the rest of the app.
func New(dataFileName string, cryptoCfg *cryptoconfig.CryptoConfig) (*FileStore, error) {
	if !fileExists(dataFileName) {
		_, err := createDataFile(dataFileName)
		if err != nil {
			return nil, fmt.Errorf("error creating file: %w", err)
		}
	}

	fsc := &FileStore{
		mu:           sync.Mutex{},
		DataFileName: dataFileName,
		Store:        make(map[string][]byte),
		cryptoCfg:    cryptoCfg,
	}

	return fsc, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true // File exists
	}

	if errors.Is(err, os.ErrNotExist) {
		return false // File doesn't exists because ErrNotExists explicity says this
	}

	return false
}

func createDataFile(filename string) (*os.File, error) {
	if !fileExists(filename) {
		f, err := os.Create(filename)
		if err != nil {
			return nil, err
		}

		return f, nil
	}

	return nil, ErrFileExists
}
