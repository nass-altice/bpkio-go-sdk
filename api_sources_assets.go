package broadpeakio

import (
	"encoding/json"
	"errors"
	"fmt"
)

type AssetInput struct {
	Name        string `json:"name,omitempty"` //req
	Url         string `json:"url,omitempty"`  //req
	Description string `json:"description,omitempty"`
	BackupIp    string `json:"backupIp,omitempty"`
}

type AssetOutput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	BackupIp    string `json:"backupIp"`
	Format      string `json:"format"`
	Id          uint   `json:"id"`
}

func (client BroadpeakClient) CreateAsset(options AssetInput) (AssetOutput, error) {
	url := baseUrl + "sources/asset"
	requiredFields := []string{"Name", "Url"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPostRequest(client, url, string(jsonBody))

		if err != nil {
			return AssetOutput{}, err
		}

		var assetApiResponse AssetOutput

		err = deserialiseApiResponse(resp, &assetApiResponse)

		if err != nil {
			return AssetOutput{}, err
		}

		return assetApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return AssetOutput{}, errMsg
}

func (client BroadpeakClient) GetAsset(id uint) (AssetOutput, error) {
	url := fmt.Sprintf(baseUrl+"sources/asset/%d", id)
	resp, err := httpGetRequest(client, url)

	if err != nil {
		return AssetOutput{}, err
	}
	var assetApiResponse AssetOutput

	err = deserialiseApiResponse(resp, &assetApiResponse)

	if err != nil {
		return AssetOutput{}, err
	}

	return assetApiResponse, nil
}

func (client BroadpeakClient) UpdateAsset(id uint, options AssetInput) (AssetOutput, error) {
	url := fmt.Sprintf(baseUrl+"sources/asset/%d", id)
	requiredFields := []string{"Name", "Url"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPutRequest(client, url, string(jsonBody))

		if err != nil {
			return AssetOutput{}, err
		}

		var assetApiResponse AssetOutput

		err = deserialiseApiResponse(resp, &assetApiResponse)

		if err != nil {
			return AssetOutput{}, err
		}

		return assetApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return AssetOutput{}, errMsg
}

func (client BroadpeakClient) DeleteAsset(id uint) (string, error) {
	url := fmt.Sprintf(baseUrl+"sources/asset/%d", id)
	resp, err := httpDeleteRequest(client, url)

	if err != nil {
		return "", err
	}
	return resp, nil
}
