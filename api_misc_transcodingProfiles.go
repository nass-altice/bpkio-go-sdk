package broadpeakio

import "fmt"

type TranscodingProfileOutput struct {
	Name       string `json:"name"`
	Content    string `json:"content"`
	Id         uint   `json:"id"`
	InternalId string `json:"internalId"`
}

func (client BroadpeakClient) GetAllTranscodingProfiles(offset uint, limit uint) ([]TranscodingProfileOutput, error) {
	url := baseUrl + "transcoding-profiles"

	url = addOffsetUrl(url, offset, limit)

	resp, err := httpGetRequest(client, url)

	if err != nil {
		return []TranscodingProfileOutput{}, err
	}

	var transcodingProfileApiResponse []TranscodingProfileOutput

	err = deserialiseApiResponse(resp, &transcodingProfileApiResponse)
	if err != nil {
		return []TranscodingProfileOutput{}, err
	}

	return transcodingProfileApiResponse, nil
}

func (client BroadpeakClient) GetTranscodingProfile(id uint) (TranscodingProfileOutput, error) {
	url := fmt.Sprintf(baseUrl+"transcoding-profiles/%d", id)

	resp, err := httpGetRequest(client, url)

	if err != nil {
		return TranscodingProfileOutput{}, err
	}

	var transcodingProfileApiResponse TranscodingProfileOutput

	err = deserialiseApiResponse(resp, &transcodingProfileApiResponse)
	if err != nil {
		return TranscodingProfileOutput{}, err
	}

	return transcodingProfileApiResponse, nil
}
