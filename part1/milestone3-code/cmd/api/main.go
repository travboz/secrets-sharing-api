package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/travboz/secrets-sharing/milestone3/internal/encryption/cryptoconfig"
	"github.com/travboz/secrets-sharing/milestone3/internal/store/filestore"
)

func main() {
	if err := loadEnv(); err != nil {
		slog.Error("Error loading .env variables", "error", err)
		os.Exit(1)
	}

	environmentVariables := []string{
		"DATA_FILE_PATH",
		"PASSWORD",
		"SALT",
	}

	envVars, err := fetchEnvVars(environmentVariables)
	if err != nil {
		slog.Error("An environment variable has not been set", "error", err)
		os.Exit(1)
	}

	slog.Debug("Env vars", "vars", envVars)

	cryptoCfg, err := cryptoconfig.New(envVars["PASSWORD"], envVars["SALT"])
	if err != nil {
		slog.Error("Error creating crypto config", "error", err)
		os.Exit(1)
	}

	fStore, err := filestore.New(envVars["DATA_FILE_PATH"], cryptoCfg)
	if err != nil {
		slog.Error("Error creating file store", "error", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	setupRoutes(mux, fStore)

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
