package broadpeakio

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CreateAdInsertionInput struct {
	Name                 string                `json:"name,omitempty"` //required
	EnvTags              []string              `json:"environmentTags,omitempty"`
	LiveAdPreRoll        *LiveAdPreRoll        `json:"liveAdPreRoll,omitempty"`
	LiveAdReplacement    *LiveAdReplacement    `json:"liveAdReplacement,omitempty"`
	VodAdInsertion       *VodAdInsertion       `json:"vodAdInsertion,omitempty"`
	TranscodingProfile   *Identifiable         `json:"transcodingProfile,omitempty"`
	EnableAdTranscoding  bool                  `json:"enableAdTranscoding,omitempty"`
	ServerSideAdTracking *ServerSideAdTracking `json:"serverSideAdTracking,omitempty"`
	Source               *Identifiable         `json:"source,omitempty"` //required
}

type UpdateAdInsertionInput struct {
	Name                 string                `json:"name,omitempty"` //required
	EnvTags              []string              `json:"environmentTags,omitempty"`
	LiveAdPreRoll        *LiveAdPreRoll        `json:"liveAdPreRoll,omitempty"`
	LiveAdReplacement    *LiveAdReplacement    `json:"liveAdReplacement,omitempty"`
	VodAdInsertion       *VodAdInsertion       `json:"vodAdInsertion,omitempty"`
	TranscodingProfile   *Identifiable         `json:"transcodingProfile,omitempty"`
	EnableAdTranscoding  bool                  `json:"enableAdTranscoding,omitempty"`
	ServerSideAdTracking *ServerSideAdTracking `json:"serverSideAdTracking,omitempty"`
}

type AuthorizationHeader struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
type AdvancedOptions struct {
	AuthorizationHeader AuthorizationHeader `json:"authorizationHeader,omitempty"`
}

type AdInsertionOutput struct {
	Name                 string                  `json:"name,omitempty"`
	Tags                 []string                `json:"tags,omitempty"`
	Type                 string                  `json:"type"`
	State                string                  `json:"state"`
	LiveAdPreRoll        LiveAdPreRollOutput     `json:"liveAdPreRoll"`
	LiveAdReplacement    LiveAdReplacementOutput `json:"liveAdReplacement"`
	VodAdInsertion       VodAdInsertionOutput    `json:"vodAdInsertion"`
	TranscodingProfile   TranscodingProfile      `json:"transcodingProfile"`
	EnableAdTranscoding  bool                    `json:"enableAdTranscoding"`
	ServerSideAdTracking ServerSideAdTracking    `json:"serverSideAdTracking"`
	AdvancedOptions      AdvancedOptions         `json:"advancedOptions"`
	Source               Source                  `json:"source"`
	Id                   uint                    `json:"id"`
	CreationDate         string                  `json:"creationDate"`
	UpdateDate           string                  `json:"updateDate"`
	Url                  string                  `json:"url"`
}

func (client BroadpeakClient) CreateAdInsertion(options CreateAdInsertionInput) (AdInsertionOutput, error) {
	url := baseUrl + "services/ad-insertion"

	requiredFields := []string{"Name", "Source"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPostRequest(client, url, string(jsonBody))

		if err != nil {
			return AdInsertionOutput{}, err
		}

		var adInsertionApiResponse AdInsertionOutput

		err = deserialiseApiResponse(resp, &adInsertionApiResponse)
		if err != nil {
			return AdInsertionOutput{}, err
		}

		return adInsertionApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return AdInsertionOutput{}, errMsg
}

func (client BroadpeakClient) GetAdInsertion(id uint) (AdInsertionOutput, error) {
	url := fmt.Sprintf(baseUrl+"services/ad-insertion/%d", id)

	resp, err := httpGetRequest(client, url)

	if err != nil {
		return AdInsertionOutput{}, err
	}

	var adInsertionApiResponse AdInsertionOutput

	err = deserialiseApiResponse(resp, &adInsertionApiResponse)
	if err != nil {
		return AdInsertionOutput{}, err
	}

	return adInsertionApiResponse, nil
}

func (client BroadpeakClient) UpdateAdInsertion(id uint, options UpdateAdInsertionInput) (AdInsertionOutput, error) {
	url := fmt.Sprintf(baseUrl+"services/ad-insertion/%d", id)

	requiredFields := []string{"Name"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPutRequest(client, url, string(jsonBody))

		if err != nil {
			return AdInsertionOutput{}, err
		}

		var adInsertionApiResponse AdInsertionOutput

		err = deserialiseApiResponse(resp, &adInsertionApiResponse)
		if err != nil {
			return AdInsertionOutput{}, err
		}

		return adInsertionApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return AdInsertionOutput{}, errMsg
}

func (client BroadpeakClient) DeleteAdInsertion(id uint) (string, error) {
	url := fmt.Sprintf(baseUrl+"services/ad-insertion/%d", id)

	resp, err := httpDeleteRequest(client, url)

	if err != nil {
		return "", err
	}

	return resp, nil
}
