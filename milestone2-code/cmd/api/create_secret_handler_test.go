package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/travboz/secrets-sharing/milestone2/internal/store"
	"github.com/travboz/secrets-sharing/milestone2/pkg/testing/assert"
)

func TestCreateSecretHandler(t *testing.T) {
	t.Run("Successfully create a secret", func(t *testing.T) {

		// Create a mock store
		readFn := func(key string) (string, error) {
			if key == "random-key" {
				return "super-secret", nil
			}
			return "", nil
		}

		writeFn := func(data store.SecretData) error { return nil }

		mockStore := &MockFileStore{
			ReadFunc:  readFn,
			WriteFunc: writeFn,
		}

		mux := http.NewServeMux()
		setupRoutes(mux, mockStore)

		ts := httptest.NewServer(mux)
		defer ts.Close()

		inputSecret := "blue"
		expectedHash := HashSecret(inputSecret)
		payload := CreateSecretRequest{Plaintext: inputSecret}

		jsonBytes, err := json.Marshal(payload)
		if err != nil {
			t.Fatal("Error marshalling CreateSecretRequest to JSON", err)
		}

		resp, err := ts.Client().Post(
			ts.URL+"/",
			ContentTypeJSON,
			bytes.NewReader(jsonBytes),
		)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		// Check status code
		assert.Equal(t, resp.StatusCode, http.StatusCreated)

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("Error reading body", err)
		}

		body = bytes.TrimSpace(body)

		got := string(body)
		want := fmt.Sprintf(`{"id":"%s"}`, expectedHash)

		assert.Equal(t, got, want)
	})

	t.Run("Invalid request body returns 400", func(t *testing.T) {
		// Setup: Create a mock store
		readFn := func(key string) (string, error) { return "", nil }
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

		t.Run("empty body", func(t *testing.T) {
			emptySecret := ""
			payload := CreateSecretRequest{Plaintext: emptySecret}

			jsonBytes, err := json.Marshal(payload)
			if err != nil {
				t.Fatal("Error marshalling CreateSecretRequest to JSON", err)
			}

			resp, err := ts.Client().Post(
				ts.URL+"/",
				ContentTypeJSON,
				bytes.NewReader(jsonBytes),
			)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			// Check status code
			assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
		})

		t.Run("bad json data", func(t *testing.T) {
			/* Examples of bad json:
			// Missing quotes around key
			badJSON := `{secret:"blue"}`
			// Trailing comma
			badJSON := `{"secret":"blue",}`
			// Incomplete/truncated
			badJSON := `{"secret":"blue"`
			// Not JSON at all
			badJSON := `not json at all`
			// Invalid characters
			badJSON := `{"secret":undefined}`
			*/

			badJSON := `{"secret":"blue"s"}`

			resp, err := ts.Client().Post(
				ts.URL+"/",
				ContentTypeJSON,
				strings.NewReader(badJSON),
			)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			// Check status code
			assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
		})

	})
}
