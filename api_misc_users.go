package broadpeakio

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type UserInput struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}

type UserOutput struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	TenantId     uint   `json:"tenantId"`
	Id           uint   `json:"id"`
	UpdateDate   string `json:"updateDate"`
	CreationDate string `json:"creationDate"`
}

func (client BroadpeakClient) CreateUser(options UserInput) (UserOutput, error) {
	url := baseUrl + "users"

	requiredFields := []string{"FistName", "LastName", "Email"}

	isOk, missingField := checkRequiredFieldsNonEmpty(options, requiredFields)

	if isOk {
		jsonBody, _ := json.Marshal(options)

		resp, err := httpPostRequest(client, url, string(jsonBody))

		if err != nil {
			return UserOutput{}, err
		}

		var userApiResponse UserOutput

		err = deserialiseApiResponse(resp, &userApiResponse)
		if err != nil {
			return UserOutput{}, err
		}

		return userApiResponse, nil
	}

	errMsg := errors.New("Error: " + missingField + " is missing")
	return UserOutput{}, errMsg
}

func (client BroadpeakClient) GetAllUsers(offset uint, limit uint, withEmailStatus bool) ([]UserOutput, error) {
	url := baseUrl + "users"
	questionMark := '?'

	url = addOffsetUrl(url, offset, limit)

	if strings.ContainsRune(url, questionMark) {
		url = url + "&with-email-status=" + fmt.Sprint(withEmailStatus)
	} else {
		url = url + "?with-email-status=" + fmt.Sprint(withEmailStatus)
	}

	resp, err := httpGetRequest(client, url)

	if err != nil {
		return []UserOutput{}, err
	}

	var userApiResponse []UserOutput

	err = deserialiseApiResponse(resp, &userApiResponse)
	if err != nil {
		return []UserOutput{}, err
	}

	return userApiResponse, nil
}

func (client BroadpeakClient) GetUser(id uint) (UserOutput, error) {
	url := fmt.Sprintf(baseUrl+"users/%d", id)

	resp, err := httpGetRequest(client, url)

	if err != nil {
		return UserOutput{}, err
	}

	var userApiResponse UserOutput

	err = deserialiseApiResponse(resp, &userApiResponse)
	if err != nil {
		return UserOutput{}, err
	}

	return userApiResponse, nil
}

func (client BroadpeakClient) UpdateUser(id uint, options UserInput) (UserOutput, error) {
	url := fmt.Sprintf(baseUrl+"users/%d", id)

	jsonBody, _ := json.Marshal(options)

	resp, err := httpPutRequest(client, url, string(jsonBody))

	if err != nil {
		return UserOutput{}, err
	}

	var userApiResponse UserOutput

	err = deserialiseApiResponse(resp, &userApiResponse)
	if err != nil {
		return UserOutput{}, err
	}

	return userApiResponse, nil
}

func (client BroadpeakClient) DeleteUser(id uint) (string, error) {
	url := fmt.Sprintf(baseUrl+"users/%d", id)

	resp, err := httpDeleteRequest(client, url)

	if err != nil {
		return "", err
	}

	return resp, nil
}
