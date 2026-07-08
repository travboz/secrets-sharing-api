package main

import (
	"net/http"

	"github.com/travboz/secrets-sharing/milestone1/internal/filestore"
)

func setupRoutes(m *http.ServeMux, s *filestore.FileStore) {
	m.HandleFunc("GET /healthcheck", healthcheckHandler)
	m.HandleFunc("/", secretHandler(s))
}
