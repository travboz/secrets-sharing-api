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
	return fmt.Sprintf("{\"id\": %q}", r.Id)
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
