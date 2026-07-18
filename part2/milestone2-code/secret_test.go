package main

// ----- This file is kept for my learning experience because I wrapped my head around the http.Handler interface.

// // Custom testServer which embeds a *httptest.Server instance.
// type testServer struct {
// 	*httptest.Server
// }

// // Create a newTestServer helper which initalizes and returns a new instance
// // of our custom testServer type.
// func newTestServer(t *testing.T, h http.Handler) *testServer {
// 	ts := httptest.NewServer(h)
// 	return &testServer{ts}
// }

// type SecretsMock struct{}

// func (s *SecretsMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	SecretsMockHandler(w, r)
// }

// func SecretsMockHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {

// 	case http.MethodGet:
// 		if r.URL.Path != "/" {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(GetSecretResponse{
// 			Data: "tomatoes",
// 		})

// 	case http.MethodPost:
// 		if r.URL.Path != "/" {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusCreated)
// 		json.NewEncoder(w).Encode(CreateSecretResponse{
// 			Id: "hash-of-super-secret",
// 		})

// 	default:
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 	}
// }

// func TestCreateSecret(t *testing.T) {
// 	t.Run("returns a valid CreateSecretResponse when the server responds with a well-formed JSON body", func(t *testing.T) {
// 		testSecret := "secret"
// 		wantId := "hash-of-super-secret"

// 		handler := &SecretsMock{}

// 		ts := newTestServer(t, handler)
// 		defer ts.Close()

// 		result, err := createSecret(ts.URL, testSecret)
// 		if err != nil {
// 			t.Fatalf("unexpected error: %q", err)
// 		}

// 		if result.Id != wantId {
// 			t.Errorf("got ID %q, want %q", result.Id, wantId)
// 		}
// 	})

// 	t.Run("fails when a request is made using the wrong apiUrl", func(t *testing.T) {
// 		testSecret := "secret"

// 		handler := &SecretsMock{}

// 		ts := newTestServer(t, handler)
// 		defer ts.Close()

// 		wrongUrl := ts.URL + "/wrong/endpoint"
// 		_, err := createSecret(wrongUrl, testSecret)

// 		if err == nil {
// 			t.Fatal("wanted error, but didn't get one")
// 		}

// 	})

// }

// func TestGetSecrets(t *testing.T) {}
