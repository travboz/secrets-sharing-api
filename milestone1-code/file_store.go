package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"
)

var (
	ErrFileExists = errors.New("File already exists")
)

// A concurrency-safe file based store.
type fileStore struct {
	Mu           sync.RWMutex
	DataFileName string
	Store        map[string]string `json:"store"`
}

type myDataType struct {
	Key   string
	Value string
}

func (s *fileStore) Write(data myDataType) error {
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	err := s.readFromFile(s.DataFileName)
	if err != nil {
		return err
	}

	s.Store[data.Key] = data.Value
	return s.writeToFile()
}

// writeToFile is a helper function which essentially writes the Store object to disk.
// Truncates file if it exists, so we are wriitng the complete Store object every time.
func (s *fileStore) writeToFile() error {
	var f *os.File

	jsonData, err := json.Marshal(s.Store)
	if err != nil {
		return fmt.Errorf("Error marshalling data: %w", err)
	}

	f, err = os.Create(s.DataFileName)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(jsonData)
	if err != nil {
		return fmt.Errorf("Error writing store data to file: %w", err)
	}

	return nil
}

// Read fetches the item with the id of key. path is the filepath of the Store.
func (s *fileStore) Read(key string) (string, error) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	err := s.readFromFile(s.DataFileName)
	if err != nil {
		return "", fmt.Errorf("Error reading from file: %w", err)
	}

	data := s.Store[key]
	delete(s.Store, key)

	err = s.writeToFile()
	if err != nil {
		return "", fmt.Errorf("Error writing file after read: %w", err)
	}

	return data, nil
}

// readFromFile is a helper which reads the complete file and overwrites
// the current Store object in memory.
func (s *fileStore) readFromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Error opening file for read: %w", err)
	}

	jsonData, err := io.ReadAll(f)
	if err != nil {
		slog.Error("Error reading data from file", "error", err)
		os.Exit(1)
	}

	if len(jsonData) != 0 {
		return json.Unmarshal(jsonData, &s.Store)
	}

	return nil
}

// Init creates the file if it doesn't exist, and initialises the FileStoreConfig
// to be used in the rest of the app.
func NewFileStore(dataFileName string) (*fileStore, error) {
	if !fileExists(dataFileName) {
		_, err := createDataFile(dataFileName)
		if err != nil {
			return nil, fmt.Errorf("Error creating file: %w", err)
		}
	}

	fsc := &fileStore{
		DataFileName: dataFileName,
		Mu:           sync.RWMutex{},
		Store:        make(map[string]string),
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
