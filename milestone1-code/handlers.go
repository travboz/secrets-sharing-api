package main

import (
	"fmt"
	"net/http"
)

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "healthy and ready to serve requests")
}

func secretGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "healthy and ready to serve *secret* GET requests")
}

func secretPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "healthy and ready to serve *secret* POST requests")
}

func setupRoutes(m *http.ServeMux) {
	m.HandleFunc("GET /healthcheck", healthcheckHandler)
	m.HandleFunc("GET /", secretGetHandler)
	m.HandleFunc("POST /", secretPostHandler)
}
