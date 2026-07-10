package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/travboz/secrets-sharing/milestone2/internal/store"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func secretHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			createSecretHandler(s)(w, r)
		case "GET":
			getSecretHandler(s)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

func getSecretHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path
		id = strings.TrimPrefix(id, "/")
		if len(id) == 0 {
			http.Error(w, "id value must be provided", http.StatusBadRequest)
			return
		}

		secret, err := s.Read(id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrKeyNotFound):
				http.Error(w, err.Error(), http.StatusNotFound)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if err := writeJSON(w, http.StatusOK, GetSecretResponse{Data: secret}, nil); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func createSecretHandler(s store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode request body
		var payload CreateSecretRequest
		if err := readJSON(w, r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Check for length of secret
		if len(payload.Plaintext) == 0 {
			http.Error(w, "secret cannot be empty", http.StatusBadRequest)
			return
		}

		// Generate hash of secret and insert it into the store
		id := HashSecret(payload.Plaintext)
		if err := s.Write(store.SecretData{Id: id, Secret: payload.Plaintext}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Successful response sends the key ID of the newly created secret
		if err := writeJSON(w, http.StatusCreated, CreateSecretResponse{ID: id}, nil); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
