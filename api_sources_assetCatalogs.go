package broadpeakio

import (
	"encoding/json"
	"errors"
	"fmt"
)

type AssetCatalogOutput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	BackupIp    string `json:"backupIp"`
	AssetSample string `json:"assetSample"`
	Id          uint   `json:"id"`
}

type AssetCatalogInput struct {
	Name        string `json:"name,omitempty"` //req
	Url         string `json:"url,omitempty"`  //req
	Description string `json:"description,omitempty"`
	AssetSample string `json:"assetSample,omitempty"` //req
	BackupIp    string `json:"backupIp,omitempty"`
}

func (client BroadpeakClient) CreateAssetCatalog(options AssetCatalogInput) (AssetCatalogOutput, error) {
	url := baseUrl + "sources/asset-catalog"

	requiredFields := []string{"Name", "Url", "AssetSample"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPostRequest(client, url, string(jsonBody))

		if err != nil {
			return AssetCatalogOutput{}, err
		}

		var assetCatalogApiResponse AssetCatalogOutput

		err = deserialiseApiResponse(resp, &assetCatalogApiResponse)
		if err != nil {
			return AssetCatalogOutput{}, err
		}

		return assetCatalogApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return AssetCatalogOutput{}, errMsg
}

func (client BroadpeakClient) GetAssetCatalog(id uint) (AssetCatalogOutput, error) {
	url := fmt.Sprintf(baseUrl+"sources/asset-catalog/%d", id)
	resp, err := httpGetRequest(client, url)

	if err != nil {
		return AssetCatalogOutput{}, err
	}

	var assetCatalogApiResponse AssetCatalogOutput

	err = deserialiseApiResponse(resp, &assetCatalogApiResponse)
	if err != nil {
		return AssetCatalogOutput{}, err
	}

	return assetCatalogApiResponse, nil
}

func (client BroadpeakClient) UpdateAssetCatalog(id uint, options AssetCatalogInput) (AssetCatalogOutput, error) {
	url := fmt.Sprintf(baseUrl+"sources/asset-catalog/%d", id)

	requiredFields := []string{"Name", "Url", "AssetSample"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPutRequest(client, url, string(jsonBody))

		if err != nil {
			return AssetCatalogOutput{}, err
		}

		var assetCatalogApiResponse AssetCatalogOutput

		err = deserialiseApiResponse(resp, &assetCatalogApiResponse)
		if err != nil {
			return AssetCatalogOutput{}, err
		}

		return assetCatalogApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return AssetCatalogOutput{}, errMsg
}

func (client BroadpeakClient) DeleteAssetCatalog(id uint) (string, error) {
	url := fmt.Sprintf(baseUrl+"sources/asset-catalog/%d", id)
	resp, err := httpDeleteRequest(client, url)

	if err != nil {
		return "", err
	}
	return resp, nil
}
