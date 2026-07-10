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

var (
	ContentTypeJSON = "application/json"
)

func TestHandlerFunctionality(t *testing.T) {
	// Testing if the handler functionality is working as we expected:
	// 	1. We can create a new secret by sending a POST request to the secret endpoint
	//	2. We can retrieve the secret by sending a GET request to the secret endpoint
	//	3. We cannot create a secret with an invalid request body (an empty body), returns a 400
	//	4. We cannot create a secret with an invalid request body (bad JSON data), returns a 400
	// t.Run("Test for test http server to check if it is working", func(t *testing.T) {

	// 	// Create a mock store
	// 	readFn := func(key string) (string, error) { return "super-secret", nil }
	// 	writeFn := func(data store.SecretData) error { return nil }
	// 	mockStore := &MockFileStore{
	// 		ReadFunc:  readFn,
	// 		WriteFunc: writeFn,
	// 	}

	// 	mux := http.NewServeMux()
	// 	setupRoutes(mux, mockStore)

	// 	ts := httptest.NewServer(mux)
	// 	defer ts.Close()

	// 	// Example ping request to the test server
	// 	rs, err := ts.Client().Get(ts.URL + "/healthcheck")
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	assert.Equal(t, rs.StatusCode, http.StatusOK)
	// })

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
		payload := CreateSecretRequest{Secret: inputSecret}

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

		got := p.Secret
		want := inputSecret

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
			payload := CreateSecretRequest{Secret: emptySecret}

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
