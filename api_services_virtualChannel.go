package broadpeakio

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CreateVirtualChannelInput struct {
	Name                 string                `json:"name,omitempty"` //req
	EnvTags              []string              `json:"environmentTags,omitempty"`
	AdBreakInsertion     *AdBreakInsertion     `json:"adBreakInsertion,omitempty"`
	TranscodingProfile   *Identifiable         `json:"transcodingProfile,omitempty"`
	ServerSideAdTracking *ServerSideAdTracking `json:"serverSideAdTracking,omitempty"`
	EnableAdTranscoding  bool                  `json:"enableAdTranscoding,omitempty"`
	BaseLive             *Identifiable         `json:"baseLive,omitempty"` //req
}

type UpdateVirtualChannelInput struct {
	Name                 string                `json:"name,omitempty"` //req
	EnvTags              []string              `json:"environmentTags,omitempty"`
	AdBreakInsertion     *AdBreakInsertion     `json:"adBreakInsertion,omitempty"`
	TranscodingProfile   *Identifiable         `json:"transcodingProfile,omitempty"`
	ServerSideAdTracking *ServerSideAdTracking `json:"serverSideAdTracking,omitempty"`
	EnableAdTranscoding  bool                  `json:"enableAdTranscoding,omitempty"`
}

type VirtualChannelOutput struct {
	Name                 string                   `json:"name"`
	EnvironmentTags      []string                 `json:"environmentTags"`
	AdBreakInsertion     AdBreakInsertionOutput   `json:"adBreakInsertion"`
	TranscodingProfile   TranscodingProfileOutput `json:"transcodingProfile"`
	ServerSideAdTracking ServerSideAdTracking     `json:"serverSideAdTracking"`
	EnableAdTranscoding  bool                     `json:"enableAdTranscoding"`
	BaseLive             Source                   `json:"baseLive"`
	Id                   uint                     `json:"id"`
	Url                  string                   `json:"url"`
	UpdateDate           string                   `json:"updateDate"`
	CreationDate         string                   `json:"creationDate"`
}

func (client BroadpeakClient) CreateVirtualChannel(options CreateVirtualChannelInput) (VirtualChannelOutput, error) {
	url := baseUrl + "services/virtual-channel"

	requiredFields := []string{"Name", "BaseLive"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPostRequest(client, url, string(jsonBody))

		if err != nil {
			return VirtualChannelOutput{}, err
		}

		var virtualChannelApiResponse VirtualChannelOutput

		err = deserialiseApiResponse(resp, &virtualChannelApiResponse)
		if err != nil {
			return VirtualChannelOutput{}, err
		}

		return virtualChannelApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return VirtualChannelOutput{}, errMsg

}

func (client BroadpeakClient) GetVirtualChannel(id uint) (VirtualChannelOutput, error) {
	url := fmt.Sprintf(baseUrl+"services/virtual-channel/%d", id)

	resp, err := httpGetRequest(client, url)

	if err != nil {
		return VirtualChannelOutput{}, err
	}

	var virtualChannelApiResponse VirtualChannelOutput

	err = deserialiseApiResponse(resp, &virtualChannelApiResponse)
	if err != nil {
		return VirtualChannelOutput{}, err
	}

	return virtualChannelApiResponse, nil
}

func (client BroadpeakClient) UpdateVirtualChannel(id uint, options UpdateVirtualChannelInput) (VirtualChannelOutput, error) {
	url := fmt.Sprintf(baseUrl+"services/virtual-channel/%d", id)

	requiredFields := []string{"Name"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPutRequest(client, url, string(jsonBody))

		if err != nil {
			return VirtualChannelOutput{}, err
		}

		var virtualChannelApiResponse VirtualChannelOutput

		err = deserialiseApiResponse(resp, &virtualChannelApiResponse)
		if err != nil {
			return VirtualChannelOutput{}, err
		}

		return virtualChannelApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return VirtualChannelOutput{}, errMsg
}

func (client BroadpeakClient) DeleteVirtualChannel(id uint) (string, error) {
	url := fmt.Sprintf(baseUrl+"services/virtual-channel/%d", id)

	resp, err := httpDeleteRequest(client, url)

	if err != nil {
		return "", err
	}

	return resp, nil
}
