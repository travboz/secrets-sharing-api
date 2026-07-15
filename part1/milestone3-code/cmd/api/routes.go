package main

import (
	"net/http"

	"github.com/travboz/secrets-sharing/milestone3/internal/store"
)

func setupRoutes(m *http.ServeMux, s store.Store) {
	m.HandleFunc("GET /healthcheck", healthCheckHandler)
	m.HandleFunc("/", secretHandler(s))
}
