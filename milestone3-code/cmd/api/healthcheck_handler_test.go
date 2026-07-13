package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	t.Run("Request to healthcheck calls healthcheck", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
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
		req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
		rr := httptest.NewRecorder()

		healthCheckHandler(rr, req)

		resp := rr.Result()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("Error reading body in healthcheck:", err)
		}

		want := "ok"
		got := string(body)

		if got != want {
			t.Errorf("got %q, but want %q", got, want)
		}
	})
}
