package broadpeakio

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CreateContentReplacementSlotInput struct {
	Name        string        `json:"name,omitempty"`
	StartTime   string        `json:"startTime,omitempty"` //req
	EndTime     string        `json:"endTime,omitempty"`
	Duration    uint          `json:"duration,omitempty"`
	Replacement *Identifiable `json:"replacement,omitempty"`
}

type UpdateContentReplacementSlotInput struct {
	Name        string        `json:"name,omitempty"`
	StartTime   string        `json:"startTime,omitempty"` //req
	EndTime     string        `json:"endTime,omitempty"`
	Duration    uint          `json:"duration,omitempty"`
	Replacement *Identifiable `json:"replacement,omitempty"`
	Category    *Identifiable `json:"category,omitempty"`
}

type ContentReplacementSlotOutput struct {
	Name        string         `json:"name"`
	StartTime   string         `json:"startTime"`
	EndTime     string         `json:"endTime"`
	Duration    uint           `json:"duration"`
	Replacement Source         `json:"replacement"`
	Category    CategoryOutput `json:"category"`
	Id          uint           `json:"id"`
}

func (client BroadpeakClient) CreateContentReplacementSlot(serviceId uint, options CreateContentReplacementSlotInput) (ContentReplacementSlotOutput, error) {
	url := fmt.Sprintf(baseUrl+"services/content-replacement/%d/slots", serviceId)

	requiredFields := []string{"StartTime"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPostRequest(client, url, string(jsonBody))

		if err != nil {
			return ContentReplacementSlotOutput{}, err
		}

		var contentReplacementSlotApiResponse ContentReplacementSlotOutput

		err = deserialiseApiResponse(resp, &contentReplacementSlotApiResponse)
		if err != nil {
			return ContentReplacementSlotOutput{}, err
		}

		return contentReplacementSlotApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return ContentReplacementSlotOutput{}, errMsg
}

func (client BroadpeakClient) GetAllContentReplacementSlots(serviceId uint, options GetAllSlotsInput) ([]ContentReplacementSlotOutput, error) {
	apiUrl := fmt.Sprintf(baseUrl+"services/content-replacement/%d/slots", serviceId)

	reqUrl, err := getReqUrlWithQuery(apiUrl, structToMap(options))

	if err != nil {
		return []ContentReplacementSlotOutput{}, err
	}

	resp, err := httpGetRequest(client, reqUrl)

	if err != nil {
		return []ContentReplacementSlotOutput{}, err
	}

	var contentReplacementSlotApiResponse []ContentReplacementSlotOutput

	err = deserialiseApiResponse(resp, &contentReplacementSlotApiResponse)
	if err != nil {
		return []ContentReplacementSlotOutput{}, err
	}

	return contentReplacementSlotApiResponse, nil
}

func (client BroadpeakClient) GetContentReplacementSlot(serviceId uint, id uint) (ContentReplacementSlotOutput, error) {
	apiUrl := fmt.Sprintf(baseUrl+"services/content-replacement/%d/slots/%d", serviceId, id)

	resp, err := httpGetRequest(client, apiUrl)

	if err != nil {
		return ContentReplacementSlotOutput{}, err
	}

	var contentReplacementSlotApiResponse ContentReplacementSlotOutput

	err = deserialiseApiResponse(resp, &contentReplacementSlotApiResponse)
	if err != nil {
		return ContentReplacementSlotOutput{}, err
	}

	return contentReplacementSlotApiResponse, nil
}

func (client BroadpeakClient) UpdateContentReplacementSlot(serviceId uint, id uint, options UpdateContentReplacementSlotInput) (ContentReplacementSlotOutput, error) {
	apiUrl := fmt.Sprintf(baseUrl+"services/content-replacement/%d/slots/%d", serviceId, id)

	requiredFields := []string{"StartTime"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPutRequest(client, apiUrl, string(jsonBody))

		if err != nil {
			return ContentReplacementSlotOutput{}, err
		}

		var contentReplacementSlotApiResponse ContentReplacementSlotOutput

		err = deserialiseApiResponse(resp, &contentReplacementSlotApiResponse)
		if err != nil {
			return ContentReplacementSlotOutput{}, err
		}

		return contentReplacementSlotApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return ContentReplacementSlotOutput{}, errMsg
}

func (client BroadpeakClient) DeleteContentReplacementSlot(serviceId uint, id uint) (string, error) {
	apiUrl := fmt.Sprintf(baseUrl+"services/content-replacement/%d/slots/%d", serviceId, id)

	resp, err := httpDeleteRequest(client, apiUrl)

	if err != nil {
		return "", err
	}

	return resp, nil
}
