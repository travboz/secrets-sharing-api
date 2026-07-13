package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("Error loading the .env: %w", err)
	}

	return nil
}

// EnvVarMap is the lookup for the environment variables.
type EnvVarMap map[string]string

var (
	ErrEnvVarNotSet = errors.New("Environment variable must be set")
)

// fetchEnvVars is a helper to make accessing environment variables more central.
func fetchEnvVars(varNames []string) (EnvVarMap, error) {
	evm := make(EnvVarMap)

	for _, name := range varNames {
		value := os.Getenv(name)

		if value == "" {
			return nil, fmt.Errorf("%s: %w", name, ErrEnvVarNotSet)
		}

		evm[name] = value
	}

	return evm, nil
}

// writeJSON helper for sending responses to the destination
// http.ResponseWriter, the HTTP status code to send, the data to encode to JSON, and a
// header map containing any additional HTTP headers we want to include in the response.
func writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// Encode the data to JSON, returning the error if there was one.
	js, err := json.MarshalIndent(data, " ", "\t") // use: json.Marshal(data) for more performant responses.
	if err != nil {
		return err
	}

	// Append a newline to make it easier to view in terminal applications.
	js = append(js, '\n')

	// Copy over headers we want to include in the response.
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Label as correct content type and then write json and status.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

// readJSON reads the request body into dst, and triages errors to return descriptive reasons why it errored.
func readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	// Use http.MaxBytesReader() to limit the size of the request body to 1MB.
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// Decode the request body, and DO NOT ALLOW any fields we haven't defined/mentioned in the struct
	// we're going to decode into.
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	// Decode the request body into the target destination.
	err := dec.Decode(dst)

	if err != nil {
		// If there is an error during decoding, start the triage...
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		// Add a new maxBytesError variable.
		var maxBytesError *http.MaxBytesError

		switch {
		// Check whether a wrapped error `has` the type we're looking for
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		// Check for an unexpected syntax error.
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		// Ensure json unmarshal and struct types match.
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" { // error relates to a particular field
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			// catch all if no field is determined
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		// Empty json body check.
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		// Checking for a field that can't actually be mapped to the target.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		// Request body is bigger than we've allowed it to be.
		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

		// Our mistake, we passed something other than a valid struct to decode into.
		case errors.As(err, &invalidUnmarshalError):
			panic(err) // so, we panic because we mucked up, not the users (most-likely)

		default:
			return err
		}
	}

	// Check for multiple json values - we're expecting only one.
	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

// HashSecret creates a sha256 hash of a string.
func HashSecret(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
