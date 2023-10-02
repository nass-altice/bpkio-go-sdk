package broadpeakio

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CreateVirtualChannelSlotInput struct {
	Name        string        `json:"name,omitempty"`
	StartTime   string        `json:"startTime,omitempty"` //req
	EndTime     string        `json:"endTime,omitempty"`
	Duration    uint          `json:"duration,omitempty"`
	Replacement *Identifiable `json:"replacement,omitempty"`
	Category    *Identifiable `json:"category,omitempty"`
	Type        string        `json:"type,omitempty"`
}

type UpdateVirtualChannelSlotInput struct {
	Name        string        `json:"name,omitempty"`
	StartTime   string        `json:"startTime,omitempty"` //req
	EndTime     string        `json:"endTime,omitempty"`
	Duration    uint          `json:"duration,omitempty"`
	Replacement *Identifiable `json:"replacement,omitempty"`
	Category    *Identifiable `json:"category,omitempty"`
	Type        string        `json:"type,omitempty"`
}

type VirtualChannelSlotOutput struct {
	Name        string       `json:"name"`
	StartTime   string       `json:"startTime"`
	EndTime     string       `json:"endTime"`
	Duration    uint         `json:"duration"`
	Replacement Source       `json:"replacement"`
	Category    Identifiable `json:"category"`
	Type        string       `json:"type"`
	Id          uint         `json:"id"`
}

func (client BroadpeakClient) CreateVirtualChannelSlot(serviceId uint, options CreateVirtualChannelSlotInput) (VirtualChannelSlotOutput, error) {
	url := fmt.Sprintf(baseUrl+"services/virtual-channel/%d/slots", serviceId)

	requiredFields := []string{"StartTime"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPostRequest(client, url, string(jsonBody))

		if err != nil {
			return VirtualChannelSlotOutput{}, err
		}

		var virtualChannelSlotApiResponse VirtualChannelSlotOutput

		err = deserialiseApiResponse(resp, &virtualChannelSlotApiResponse)
		if err != nil {
			return VirtualChannelSlotOutput{}, err
		}

		return virtualChannelSlotApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return VirtualChannelSlotOutput{}, errMsg
}

func (client BroadpeakClient) GetAllVirtualChannelSlots(serviceId uint, options GetAllSlotsInput) ([]VirtualChannelSlotOutput, error) {
	apiUrl := fmt.Sprintf(baseUrl+"services/virtual-channel/%d/slots", serviceId)

	reqUrl, err := getReqUrlWithQuery(apiUrl, structToMap(options))

	if err != nil {
		return []VirtualChannelSlotOutput{}, err
	}

	resp, err := httpGetRequest(client, reqUrl)

	if err != nil {
		return []VirtualChannelSlotOutput{}, err
	}

	var virtualChannelSlotApiResponse []VirtualChannelSlotOutput

	err = deserialiseApiResponse(resp, &virtualChannelSlotApiResponse)
	if err != nil {
		return []VirtualChannelSlotOutput{}, err
	}

	return virtualChannelSlotApiResponse, nil
}

func (client BroadpeakClient) GetVirtualChannelSlot(serviceId uint, id uint) (VirtualChannelSlotOutput, error) {
	apiUrl := fmt.Sprintf(baseUrl+"services/virtual-channel/%d/slots/%d", serviceId, id)

	resp, err := httpGetRequest(client, apiUrl)

	if err != nil {
		return VirtualChannelSlotOutput{}, nil
	}

	var virtualChannelSlotApiResponse VirtualChannelSlotOutput

	err = deserialiseApiResponse(resp, &virtualChannelSlotApiResponse)
	if err != nil {
		return VirtualChannelSlotOutput{}, err
	}

	return virtualChannelSlotApiResponse, nil
}

func (client BroadpeakClient) UpdateVirtualChannelSlot(serviceId uint, id uint, options UpdateVirtualChannelSlotInput) (VirtualChannelSlotOutput, error) {
	apiUrl := fmt.Sprintf(baseUrl+"services/virtual-channel/%d/slots/%d", serviceId, id)

	requiredFields := []string{"StartTime"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPutRequest(client, apiUrl, string(jsonBody))

		if err != nil {
			return VirtualChannelSlotOutput{}, err
		}

		var virtualChannelSlotApiResponse VirtualChannelSlotOutput

		err = deserialiseApiResponse(resp, &virtualChannelSlotApiResponse)
		if err != nil {
			return VirtualChannelSlotOutput{}, err
		}

		return virtualChannelSlotApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return VirtualChannelSlotOutput{}, errMsg
}

func (client BroadpeakClient) DeleteVirtualChannelSlot(serviceId uint, id uint) (string, error) {
	apiUrl := fmt.Sprintf(baseUrl+"services/virtual-channel/%d/slots/%d", serviceId, id)

	resp, err := httpDeleteRequest(client, apiUrl)

	if err != nil {
		return "", err
	}

	return resp, nil
}
