package main

var (
	ContentTypeJSON = "application/json"
)

type CreateSecretRequest struct {
	Plaintext string `json:"plain_text"`
}

type CreateSecretResponse struct {
	ID string `json:"id"`
}

type GetSecretResponse struct {
	Data string `json:"data"`
}
