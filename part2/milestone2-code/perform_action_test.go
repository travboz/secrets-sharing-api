package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testPlaintextSecret string = "blueberries"
	testHash            string = "bcc286bbbe4353e6a97ae169729ed4a5"
)

// Custom testServer which embeds a *httptest.Server instance.
type testServer struct {
	*httptest.Server
}

// Create a newTestServer helper which initalizes and returns a new instance
// of our custom testServer type.
func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)
	return &testServer{ts}
}

func SecretsMockHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(GetSecretResponse{
			Data: testPlaintextSecret,
		})

	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(CreateSecretResponse{
			Id: testHash,
		})

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func TestPerformActionCreate(t *testing.T) {
	// Test server set up for test runs
	mux := http.NewServeMux()
	mux.HandleFunc("/", SecretsMockHandler)

	ts := newTestServer(t, mux)
	defer ts.Close()

	t.Run("Valid 'create' successfully creates a new secret and returns an id", func(t *testing.T) {

		expectedSecretCreationResponse := testHash

		// Build the config we expect for valid create call
		c := ClientConfig{
			Action: ActionCreate,
			URL:    ts.URL,
			Data:   testPlaintextSecret,
		}

		output, err := performAction(c)
		if err != nil {
			t.Fatalf("wanted nil error, got: %q", err)
		}

		if output == "" {
			t.Fatalf("want %q, but got empty string", expectedSecretCreationResponse)
		}

		if output != expectedSecretCreationResponse {
			t.Errorf("want %q, but got %q", expectedSecretCreationResponse, output)
		}
	})

	t.Run("Valid 'view' successfully fetches a secret by id", func(t *testing.T) {
		expectedSecretViewResponse := testPlaintextSecret

		// Build the config we expect for valid create call
		c := ClientConfig{
			Action: ActionView,
			URL:    ts.URL,
			Id:     testHash,
		}

		output, err := performAction(c)
		if err != nil {
			t.Fatalf("wanted nil error, got: %q", err)
		}

		if output == "" {
			t.Fatalf("want %q, but got empty string", expectedSecretViewResponse)
		}

		if output != expectedSecretViewResponse {
			t.Errorf("want %q, but got %q", expectedSecretViewResponse, output)
		}
	})
}

func TestPerformActionInvalid(t *testing.T) {
	t.Run("Invalid action fails to send request", func(t *testing.T) {
		// Build the config we expect for valid create call
		c := ClientConfig{
			Action: "invalid",
			URL:    "valid url",
			Id:     "valid hash",
		}

		_, err := performAction(c)
		if err == nil {
			t.Fatalf("wanted error, but got nothing")
		}

		if err != ErrInvalidAction {
			t.Errorf("want error: %q, for error: %q", ErrInvalidAction, err)
		}
	})
}
