package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type CreateSecretResponse struct {
	Id string `json:"id"`
}

// createSecret makes a request to the Secret Sharing API server's POST endpoint and
// returns the stored secret's ID if successful.
func createSecret(apiURL string, plainText string) (CreateSecretResponse, error) {
	payload := fmt.Sprintf(`{"plain_text":"%s"}`, plainText)

	resp, err := http.Post(apiURL, "application/json", strings.NewReader(payload))
	if err != nil {
		return CreateSecretResponse{}, err
	}

	if resp.StatusCode != http.StatusCreated {
		return CreateSecretResponse{}, err
	}

	var result CreateSecretResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return CreateSecretResponse{}, err
	}

	return result, nil
}

type GetSecretResponse struct {
	Data string `json:"data"`
}

// getSecret makes a request to the Secret Sharing API server's GET /id endpoint and
// returns the plaintext secret if successful.
func getSecret(apiURL string, secretID string) (GetSecretResponse, error) {
	completeUrl := fmt.Sprintf("%s/%s", apiURL, secretID)
	resp, err := http.Get(completeUrl)
	if err != nil {
		return GetSecretResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return GetSecretResponse{}, err
	}

	var result GetSecretResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return GetSecretResponse{}, err
	}

	return result, nil
}
