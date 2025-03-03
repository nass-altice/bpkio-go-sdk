package broadpeakio

import (
	"encoding/json"
	"errors"
	"fmt"
)

type AdServerInput struct {
	Name            string       `json:"name,omitempty"` //req
	Url             string       `json:"url,omitempty"`  //req
	Description     string       `json:"description,omitempty"`
	Queries         string       `json:"queries,omitempty"`
	QueryParameters []QueryParam `json:"queryParameters,omitempty"`
	Template        string       `json:"template,omitempty"` //req
}

type AdServerOutput struct {
	Name            string       `json:"name"`
	Description     string       `json:"description"`
	Type            string       `json:"type"`
	Url             string       `json:"url"`
	Queries         string       `json:"queries"`
	QueryParameters []QueryParam `json:"queryParameters"`
	Template        string       `json:"template"`
	Id              uint         `json:"id"`
}

func (client BroadpeakClient) CreateAdServer(options AdServerInput) (AdServerOutput, error) {
	url := baseUrl + "sources/ad-server"

	requiredFields := []string{"Name", "Url"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPostRequest(client, url, string(jsonBody))

		if err != nil {
			return AdServerOutput{}, err
		}

		var adServerApiResponse AdServerOutput

		err = deserialiseApiResponse(resp, &adServerApiResponse)
		if err != nil {
			return AdServerOutput{}, err
		}

		return adServerApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return AdServerOutput{}, errMsg
}

func (client BroadpeakClient) GetAdServer(id uint) (AdServerOutput, error) {
	url := fmt.Sprintf(baseUrl+"sources/ad-server/%d", id)
	resp, err := httpGetRequest(client, url)

	if err != nil {
		return AdServerOutput{}, err
	}
	var adServerApiResponse AdServerOutput

	err = deserialiseApiResponse(resp, &adServerApiResponse)
	if err != nil {
		return AdServerOutput{}, err
	}

	return adServerApiResponse, nil
}

func (client BroadpeakClient) UpdateAdServer(id uint, options AdServerInput) (AdServerOutput, error) {
	url := fmt.Sprintf(baseUrl+"sources/ad-server/%d", id)
	requiredFields := []string{"Name", "Url"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPutRequest(client, url, string(jsonBody))

		if err != nil {
			return AdServerOutput{}, err
		}

		var adServerApiResponse AdServerOutput

		err = deserialiseApiResponse(resp, &adServerApiResponse)
		if err != nil {
			return AdServerOutput{}, err
		}

		return adServerApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return AdServerOutput{}, errMsg
}

func (client BroadpeakClient) DeleteAdServer(id uint) (string, error) {
	url := fmt.Sprintf(baseUrl+"sources/ad-server/%d", id)
	resp, err := httpDeleteRequest(client, url)

	if err != nil {
		return "", err
	}
	return resp, nil
}
