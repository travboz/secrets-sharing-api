package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateSecretResponse struct {
	Id string `json:"id"`
}

func (r CreateSecretResponse) String() string {
	return fmt.Sprintf("%s", r.Id)
}

func createSecret(apiURL string, plainText string) (CreateSecretResponse, error) {
	payload := map[string]string{
		"plain_text": plainText,
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return CreateSecretResponse{}, err
	}

	bytesReader := bytes.NewReader(jsonBytes)
	req, err := http.NewRequest(http.MethodPost, apiURL, bytesReader)
	if err != nil {
		return CreateSecretResponse{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return CreateSecretResponse{}, err
	}
	defer resp.Body.Close()

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

func (r GetSecretResponse) String() string {
	return fmt.Sprintf("%s", r.Data)
}

func getSecret(apiURL string, secretID string) (GetSecretResponse, error) {
	completeUrl := fmt.Sprintf("%s/%s", apiURL, secretID)
	req, err := http.NewRequest(http.MethodGet, completeUrl, nil)
	if err != nil {
		return GetSecretResponse{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return GetSecretResponse{}, err
	}
	defer resp.Body.Close()

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
