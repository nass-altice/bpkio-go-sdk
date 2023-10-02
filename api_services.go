package broadpeakio

type ServiceOutput struct {
	Name            string   `json:"name"`
	EnvironmentTags []string `json:"environmentTags"`
	Id              uint     `json:"id"`
	Type            string   `json:"type"`
	Url             string   `json:"url"`
	UpdateDate      string   `json:"updateDate"`
	CreationDate    string   `json:"creationDate"`
}

func (client BroadpeakClient) GetAllServices(offset uint, limit uint) ([]ServiceOutput, error) {
	url := baseUrl + "services"
	url = addOffsetUrl(url, offset, limit)
	resp, err := httpGetRequest(client, url)

	if err != nil {
		return []ServiceOutput{}, err
	}

	var serviceApiResponse []ServiceOutput

	err = deserialiseApiResponse(resp, &serviceApiResponse)
	if err != nil {
		return []ServiceOutput{}, err
	}

	return serviceApiResponse, nil
}
