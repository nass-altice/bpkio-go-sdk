package broadpeakio

type SamplesOutput struct {
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Status      string `json:"status,omitempty"`
	Reason      string `json:"reason,omitempty"`
	Url         string `json:"url,omitempty"`
	AssetSample string `json:"assetSample,omitempty"`
	Content     string `json:"content,omitempty"`
}

func (client BroadpeakClient) CreateSamples() ([]SamplesOutput, error) {
	url := baseUrl + "samples"

	resp, err := httpPostRequest(client, url, "")

	if err != nil {
		return []SamplesOutput{}, err
	}

	var samplesApiResponse []SamplesOutput

	err = deserialiseApiResponse(resp, &samplesApiResponse)
	if err != nil {
		return []SamplesOutput{}, err
	}

	return samplesApiResponse, nil
}
