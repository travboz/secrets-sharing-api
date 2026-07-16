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
