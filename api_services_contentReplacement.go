package broadpeakio

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CreateContentReplacementInput struct {
	Name        string        `json:"name,omitempty"`        //req
	Replacement *Identifiable `json:"replacement,omitempty"` //req
	Source      *Identifiable `json:"source,omitempty"`      //req
	EnvTags     []string      `json:"environmentTags,omitempty"`
}

type UpdateContentReplacementInput struct {
	Name        string        `json:"name,omitempty"`        //req
	Replacement *Identifiable `json:"replacement,omitempty"` //req
	EnvTags     []string      `json:"environmentTags,omitempty"`
}

type ContentReplacementOutput struct {
	Name            string   `json:"name"`
	EnvironmentTags []string `json:"environmentTags"`
	Replacement     Source   `json:"replacement"`
	Source          Source   `json:"source"`
	Id              uint     `json:"id"`
	Url             string   `json:"url"`
	UpdateDate      string   `json:"updateDate"`
	CreationDate    string   `json:"creationDate"`
}

func (client BroadpeakClient) CreateContentReplacement(options CreateContentReplacementInput) (ContentReplacementOutput, error) {
	url := baseUrl + "services/content-replacement"
	requiredFields := []string{"Name", "Replacement", "Source"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPostRequest(client, url, string(jsonBody))

		if err != nil {
			return ContentReplacementOutput{}, err
		}

		var contentReplacementApiResponse ContentReplacementOutput

		err = deserialiseApiResponse(resp, &contentReplacementApiResponse)

		if err != nil {
			return ContentReplacementOutput{}, err
		}

		return contentReplacementApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return ContentReplacementOutput{}, errMsg
}

func (client BroadpeakClient) GetContentReplacement(id uint) (ContentReplacementOutput, error) {
	url := fmt.Sprintf(baseUrl+"services/content-replacement/%d", id)

	resp, err := httpGetRequest(client, url)

	if err != nil {
		return ContentReplacementOutput{}, err
	}
	var contentReplacementApiResponse ContentReplacementOutput

	err = deserialiseApiResponse(resp, &contentReplacementApiResponse)

	if err != nil {
		return ContentReplacementOutput{}, err
	}

	return contentReplacementApiResponse, nil
}

func (client BroadpeakClient) UpdateContentReplacement(id uint, options UpdateContentReplacementInput) (ContentReplacementOutput, error) {
	url := fmt.Sprintf(baseUrl+"services/content-replacement/%d", id)

	requiredFields := []string{"Name", "Replacement"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPutRequest(client, url, string(jsonBody))

		if err != nil {
			return ContentReplacementOutput{}, err
		}

		var contentReplacementApiResponse ContentReplacementOutput

		err = deserialiseApiResponse(resp, &contentReplacementApiResponse)

		if err != nil {
			return ContentReplacementOutput{}, err
		}

		return contentReplacementApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return ContentReplacementOutput{}, errMsg
}

func (client BroadpeakClient) DeleteContentReplacement(id uint) (string, error) {
	url := fmt.Sprintf(baseUrl+"services/content-replacement/%d", id)

	resp, err := httpDeleteRequest(client, url)

	if err != nil {
		return "", err
	}
	return resp, nil
}
