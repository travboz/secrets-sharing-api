package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
)

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "healthy and ready to serve requests")
}

func secretGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "healthy and ready to serve *secret* GET requests")
}

type SecretPostRequest struct {
	Secret string `json:"secret"`
}

type SecretPostResponse struct {
	ID string `json:"id"`
}

func secretPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload SecretPostRequest
	if err := readJSON(w, r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rn := strconv.Itoa(rand.IntN(100))
	if err := writeJSON(w, http.StatusOK, map[string]any{"data": SecretPostResponse{ID: rn}}, nil); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func setupRoutes(m *http.ServeMux) {
	m.HandleFunc("GET /healthcheck", healthcheckHandler)
	m.HandleFunc("GET /", secretGetHandler)
	m.HandleFunc("POST /", secretPostHandler)
}
