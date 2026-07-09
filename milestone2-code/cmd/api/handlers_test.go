package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/travboz/secrets-sharing/milestone2/internal/store"
)

// Selecting to mock but use a MockFileStore which mimics behaviour.
/*
This MockFileStore has *exactly the same* functions that are used by
the handlers in the test. We control the input and output.
Where the handler calls `s.Write(blah)`, it's instead going to call
our MockFileStore.Write(blah).
How cool is that?!
*/
type MockFileStore struct {
	ReadFunc  func(key string) (string, error)
	WriteFunc func(data store.SecretData) error
}

func (m *MockFileStore) Read(key string) (string, error) {
	return m.ReadFunc(key)
}

func (m *MockFileStore) Write(data store.SecretData) error {
	return m.WriteFunc(data)
}

/*
Example of the ReadFunc mock function:
	mock := &MockFileStore{
		ReadFunc: func(key string) (string, error) {
			# This function runs when handler calls mock.Read()
			return "my-secret-value", nil
		},
	}
*/

func TestEndpointRouting(t *testing.T) {
	t.Run("GET request to secret endpoint succeeds", func(t *testing.T) {
		readFn := func(key string) (string, error) { return "super-secret", nil }
		writeFn := func(data store.SecretData) error { return nil }
		mockStore := &MockFileStore{
			ReadFunc:  readFn,
			WriteFunc: writeFn,
		}

		req := httptest.NewRequest(http.MethodGet, "/someid", nil)
		rr := httptest.NewRecorder()

		secretHandler(mockStore)(rr, req)

		resp := rr.Result()

		want := http.StatusOK
		got := resp.StatusCode

		if got != want {
			t.Errorf("got '%d', but want '%d'", got, want)
		}
	})
	t.Run("POST request to secret endpoint succeeds", func(t *testing.T) {
		readFn := func(key string) (string, error) { return "super-secret", nil }
		writeFn := func(data store.SecretData) error { return nil }
		mockStore := &MockFileStore{
			ReadFunc:  readFn,
			WriteFunc: writeFn,
		}

		input := `{"secret":"blue"}`
		payload := strings.NewReader(input)

		req := httptest.NewRequest(http.MethodPost, "/", payload)
		rr := httptest.NewRecorder()

		secretHandler(mockStore)(rr, req)

		resp := rr.Result()

		want := http.StatusOK
		got := resp.StatusCode

		if got != want {
			t.Errorf("got '%d', but want '%d'", got, want)
		}

	})
	t.Run("PUT request to secret endpoint fails because method not allowed", func(t *testing.T) {
		readFn := func(key string) (string, error) { return "super-secret", nil }
		writeFn := func(data store.SecretData) error { return nil }
		mockStore := &MockFileStore{
			ReadFunc:  readFn,
			WriteFunc: writeFn,
		}

		input := `{"secret":"blue"}`
		payload := strings.NewReader(input)

		req := httptest.NewRequest(http.MethodPut, "/", payload)
		rr := httptest.NewRecorder()

		secretHandler(mockStore)(rr, req)

		resp := rr.Result()

		want := http.StatusMethodNotAllowed
		got := resp.StatusCode

		if got != want {
			t.Errorf("got '%d', but want '%d'", got, want)
		}
	})
}

func TestHealthCheck(t *testing.T) {
	t.Run("Request to healthcheck calls healthcheck", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/healthcheck", nil)
		rr := httptest.NewRecorder()

		healthCheckHandler(rr, req)

		resp := rr.Result()

		want := http.StatusOK
		got := resp.StatusCode

		if got != want {
			t.Errorf("got '%d', but want '%d'", got, want)
		}
	})

	t.Run("Healthcheck returns ok as response body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/healthcheck", nil)
		rr := httptest.NewRecorder()

		healthCheckHandler(rr, req)

		resp := rr.Result()
		body, _ := io.ReadAll(resp.Body)

		want := "ok"
		got := string(body)

		if got != want {
			t.Errorf("got %q, but want %q", got, want)
		}
	})
}

func TestEndpointErrors(t *testing.T) {
	t.Run("GET request with empty id to secret endpoint fails", func(t *testing.T) {
		// readFn := func(key string) (string, error) { return "", nil }
		// writeFn := func(data store.SecretData) error { return nil }
		mockStore := &MockFileStore{
			ReadFunc:  nil,
			WriteFunc: nil,
		}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()

		secretHandler(mockStore)(rr, req)

		resp := rr.Result()

		want := http.StatusBadRequest
		got := resp.StatusCode

		if got != want {
			t.Errorf("got '%d', but want '%d'", got, want)
		}
	})

	t.Run("POST request with empty body to secret endpoint fails", func(t *testing.T) {
		// readFn := func(key string) (string, error) { return "", nil }
		// writeFn := func(data store.SecretData) error { return nil }
		mockStore := &MockFileStore{
			ReadFunc:  nil,
			WriteFunc: nil,
		}

		input := ``
		payload := strings.NewReader(input)

		req := httptest.NewRequest(http.MethodPost, "/", payload)
		rr := httptest.NewRecorder()

		secretHandler(mockStore)(rr, req)

		resp := rr.Result()

		want := http.StatusBadRequest
		got := resp.StatusCode

		if got != want {
			t.Errorf("got '%d', but want '%d'", got, want)
		}
	})
}

// TODO: up to `Handler functionality is working as expected` testing section
