package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/travboz/secrets-sharing/milestone3/internal/store"
)

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

		input := `{"plain_text":"blue"}`
		payload := strings.NewReader(input)

		req := httptest.NewRequest(http.MethodPost, "/", payload)
		rr := httptest.NewRecorder()

		secretHandler(mockStore)(rr, req)

		resp := rr.Result()

		want := http.StatusCreated
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

		input := `{"plain_text":"blue"}`
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

func TestEndpointErrors(t *testing.T) {
	t.Run("GET request with empty id to secret endpoint fails", func(t *testing.T) {
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
