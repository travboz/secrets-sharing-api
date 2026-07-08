package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("Error loading the .env: %w", err)
	}

	return nil
}

func main() {
	if err := loadEnv(); err != nil {
		slog.Error("Error loading .env variables", "error", err)
		os.Exit(1)
	}

	DATA_FILE_PATH := os.Getenv("DATA_FILE_PATH")
	if DATA_FILE_PATH == "" {
		slog.Error("DATA_FILE_PATH must not be empty")
		os.Exit(1)
	}

	fmt.Println("File path value:", DATA_FILE_PATH)

	fStore, err := NewFileStore(DATA_FILE_PATH)
	if err != nil {
		slog.Error("Error creating file store", "error", err)
		os.Exit(1)
	}

	err = fStore.Write(myDataType{Key: "name", Value: "10000"})
	if err != nil {
		slog.Error("Error writing to file store", "error", err)
		os.Exit(1)
	}
}
