package broadpeakio

import (
	"encoding/json"
	"fmt"
)

type CreateApiKeyInput struct {
	Name           string `json:"name,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
}

type ApiKeyOutput struct {
	Name           string `json:"name"`
	ExpirationDate string `json:"expirationDate"`
	CreationDate   string `json:"creationDate"`
	Token          string `json:"token"`
}

func (client BroadpeakClient) CreateApiKey(name string, expirationDate string) (ApiKeyOutput, error) {
	url := baseUrl + "tokens"

	body := CreateApiKeyInput{name, expirationDate}

	jsonBody, _ := json.Marshal(body)

	resp, err := httpPostRequest(client, url, string(jsonBody))

	if err != nil {
		return ApiKeyOutput{}, err
	}

	var createApiResponse ApiKeyOutput

	err = deserialiseApiResponse(resp, &createApiResponse)
	if err != nil {
		return ApiKeyOutput{}, err
	}

	return createApiResponse, nil
}

func (client BroadpeakClient) GetAllApiKeys(offset uint, limit uint) ([]ApiKeyOutput, error) {
	url := baseUrl + "tokens"

	url = addOffsetUrl(url, offset, limit)

	resp, err := httpGetRequest(client, url)

	if err != nil {
		return []ApiKeyOutput{}, err
	}

	var getAllApiKeysResponse []ApiKeyOutput

	err = deserialiseApiResponse(resp, &getAllApiKeysResponse)
	if err != nil {
		return []ApiKeyOutput{}, err
	}

	return getAllApiKeysResponse, nil
}

func (client BroadpeakClient) DeleteApiKey(name string) (string, error) {
	url := fmt.Sprintf(baseUrl+"tokens/%s", name)

	resp, err := httpDeleteRequest(client, url)

	if err != nil {
		return "", err
	}

	return resp, nil
}
