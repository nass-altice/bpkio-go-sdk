package broadpeakio

import (
	"encoding/json"
	"errors"
	"fmt"
)

type LiveInput struct {
	Name        string `json:"name,omitempty"` //req
	Url         string `json:"url,omitempty"`  //req
	BackupIp    string `json:"backupIp,omitempty"`
	Description string `json:"description,omitempty"`
	MultiPeriod bool   `json:"multiPeriod,omitempty"`
	Origin      Origin `json:"origin"`
}

type LiveOutput struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	BackupIp    string `json:"backupIp"`
	Format      string `json:"format"`
	Type        string `json:"type"`
	Origin      Origin `json:"origin"`
	Description string `json:"description"`
	MultiPeriod bool   `json:"multiPeriod"`
	Id          uint   `json:"id"`
}

func (client BroadpeakClient) CreateLive(options LiveInput) (LiveOutput, error) {
	url := baseUrl + "sources/live"
	requiredFields := []string{"Name", "Url"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPostRequest(client, url, string(jsonBody))

		if err != nil {
			return LiveOutput{}, err
		}
		var liveApiResponse LiveOutput
		err = deserialiseApiResponse(resp, &liveApiResponse)
		if err != nil {
			return LiveOutput{}, err
		}

		return liveApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return LiveOutput{}, errMsg
}

func (client BroadpeakClient) GetLive(id uint) (LiveOutput, error) {
	url := fmt.Sprintf(baseUrl+"sources/live/%d", id)
	resp, err := httpGetRequest(client, url)

	if err != nil {
		return LiveOutput{}, err
	}
	var liveApiResponse LiveOutput
	err = deserialiseApiResponse(resp, &liveApiResponse)
	if err != nil {
		return LiveOutput{}, err
	}

	return liveApiResponse, nil
}

func (client BroadpeakClient) UpdateLive(id uint, options LiveInput) (LiveOutput, error) {
	url := fmt.Sprintf(baseUrl+"sources/live/%d", id)
	requiredFields := []string{"Name", "Url"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPutRequest(client, url, string(jsonBody))

		if err != nil {
			return LiveOutput{}, err
		}

		var liveApiResponse LiveOutput
		err = deserialiseApiResponse(resp, &liveApiResponse)
		if err != nil {
			return LiveOutput{}, err
		}

		return liveApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return LiveOutput{}, errMsg
}

func (client BroadpeakClient) DeleteLive(id uint) (string, error) {
	url := fmt.Sprintf(baseUrl+"sources/live/%d", id)
	resp, err := httpDeleteRequest(client, url)

	if err != nil {
		return "", err
	}
	return resp, nil
}
