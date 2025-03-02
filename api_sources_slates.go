package broadpeakio

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SlateInput struct {
	Name        string `json:"name,omitempty"` //req
	Url         string `json:"url,omitempty"`  //req
	Description string `json:"description,omitempty"`
}

type SlateOutput struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Format      string `json:"format"`
}

func (client BroadpeakClient) CreateSlate(options SlateInput) (SlateOutput, error) {
	url := baseUrl + "sources/slate"
	requiredFields := []string{"Name", "Url"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPostRequest(client, url, string(jsonBody))

		if err != nil {
			return SlateOutput{}, err
		}

		var slateApiResponse SlateOutput

		err = deserialiseApiResponse(resp, &slateApiResponse)
		if err != nil {
			return SlateOutput{}, err
		}

		return slateApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return SlateOutput{}, errMsg
}

func (client BroadpeakClient) GetSlate(id uint) (SlateOutput, error) {
	url := fmt.Sprintf(baseUrl+"sources/slate/%d", id)
	resp, err := httpGetRequest(client, url)

	if err != nil {
		return SlateOutput{}, err
	}
	var slateApiResponse SlateOutput

	err = deserialiseApiResponse(resp, &slateApiResponse)
	if err != nil {
		return SlateOutput{}, err
	}

	return slateApiResponse, nil
}

func (client BroadpeakClient) UpdateSlate(id uint, options SlateInput) (SlateOutput, error) {
	url := fmt.Sprintf(baseUrl+"sources/slate/%d", id)
	requiredFields := []string{"Name", "Url"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPutRequest(client, url, string(jsonBody))

		if err != nil {
			return SlateOutput{}, err
		}

		var slateApiResponse SlateOutput

		err = deserialiseApiResponse(resp, &slateApiResponse)
		if err != nil {
			return SlateOutput{}, err
		}

		return slateApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return SlateOutput{}, errMsg
}

func (client BroadpeakClient) DeleteSlate(id uint) (string, error) {
	url := fmt.Sprintf(baseUrl+"sources/slate/%d", id)
	resp, err := httpDeleteRequest(client, url)

	if err != nil {
		return "", err
	}
	return resp, nil
}
