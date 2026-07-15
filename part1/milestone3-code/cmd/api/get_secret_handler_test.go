package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/travboz/secrets-sharing/milestone3/internal/store"
	"github.com/travboz/secrets-sharing/milestone3/pkg/testing/assert"
)

func TestGetSecretHandler(t *testing.T) {
	t.Run("Retrieve secret via ID", func(t *testing.T) {
		inputSecret := "blue"
		id := HashSecret(inputSecret)

		// Setup: Create a mock store
		readFn := func(key string) (string, error) {
			if key == id {
				return inputSecret, nil
			}
			return "", nil
		}

		writeFn := func(data store.SecretData) error { return nil }

		mockStore := &MockFileStore{
			ReadFunc:  readFn,
			WriteFunc: writeFn,
		}

		// Setup: Create test server
		mux := http.NewServeMux()
		setupRoutes(mux, mockStore)

		ts := httptest.NewServer(mux)
		defer ts.Close()

		// Start actual testing work
		resp, err := ts.Client().Get(ts.URL + fmt.Sprintf("/%s", id))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		// Check status code
		assert.Equal(t, resp.StatusCode, http.StatusOK)

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("Error reading body", err)
		}

		// Unmarshal response body
		p := GetSecretResponse{}
		err = json.Unmarshal(body, &p)
		if err != nil {
			t.Fatal("Error unmarshalling body into GetSecretResponse", err)
		}

		got := p.Data
		want := inputSecret

		assert.Equal(t, got, want)
	})

	t.Run("Secrets can be seen one time only", func(t *testing.T) {
		inputId := "7a819afa983d454b3a368c1422ba853c"
		expectedSecret := "My super secret1234151"

		// Setup: Create a Spy mock store
		readFn := func(key string) (string, error) {
			if key == inputId {
				return expectedSecret, nil
			}
			return "some-default-secret", nil
		}

		writeFn := func(data store.SecretData) error { return nil }

		mockStore := &SpyFileStore{
			ReadFunc:  readFn,
			WriteFunc: writeFn,
			KeysSeen:  make([]string, 0),
		}

		// Setup: Create test server
		mux := http.NewServeMux()
		setupRoutes(mux, mockStore)

		ts := httptest.NewServer(mux)
		defer ts.Close()

		t.Run("Retrieve ID on first GET request to endpoint", func(t *testing.T) {
			// Start actual testing work
			resp, err := ts.Client().Get(ts.URL + fmt.Sprintf("/%s", inputId))
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			// Check status code
			assert.Equal(t, resp.StatusCode, http.StatusOK)

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal("Error reading body", err)
			}

			// Check response body matches expected secret string
			p := GetSecretResponse{}
			err = json.Unmarshal(body, &p)
			if err != nil {
				t.Fatal("Error unmarshalling body into GetSecretResponse", err)
			}

			got := p.Data
			want := expectedSecret

			assert.Equal(t, got, want)
		})

		t.Run("Second GET request to endpoint with same id fails", func(t *testing.T) {
			// Start actual testing work
			resp, err := ts.Client().Get(ts.URL + fmt.Sprintf("/%s", inputId))
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			// Check status code
			assert.Equal(t, resp.StatusCode, http.StatusNotFound)
		})

	})
}
