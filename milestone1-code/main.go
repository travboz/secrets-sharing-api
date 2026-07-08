package main

import (
	"fmt"
	"log/slog"
	"net/http"
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

	_, err := NewFileStore(DATA_FILE_PATH)
	if err != nil {
		slog.Error("Error creating file store", "error", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	setupRoutes(mux)

	port := 8080

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	slog.Info("Server listening on", "port", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		slog.Error("Error starting server", "error", err)
		os.Exit(1)
	}

}
