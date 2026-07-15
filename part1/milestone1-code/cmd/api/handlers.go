package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/travboz/secrets-sharing/milestone1/internal/filestore"
)

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "healthy and ready to serve requests")
}

func secretHandler(s *filestore.FileStore) http.HandlerFunc {
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

type GetSecretResponse struct {
	Secret string `json:"secret"`
}

func getSecretHandler(s *filestore.FileStore) http.HandlerFunc {
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
			case errors.Is(err, filestore.ErrKeyNotFound):
				http.Error(w, err.Error(), http.StatusNotFound)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if err := writeJSON(w, http.StatusOK, GetSecretResponse{Secret: secret}, nil); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

type CreateSecretRequest struct {
	Secret string `json:"secret"`
}

type CreateSecretResponse struct {
	ID string `json:"id"`
}

func createSecretHandler(s *filestore.FileStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload CreateSecretRequest
		if err := readJSON(w, r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Generate hash of secret and store it in file store
		id := HashSecret(payload.Secret)
		if err := s.Write(filestore.SecretData{Id: id, Secret: payload.Secret}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := writeJSON(w, http.StatusOK, CreateSecretResponse{ID: id}, nil); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
