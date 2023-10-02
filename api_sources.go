package broadpeakio

import "fmt"

type SourceOutput struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Type   string `json:"type"`
	Format string `json:"format"`
}

type SourceStatusOutput struct {
	MessageText   string `json:"messageText"`
	SeverityLEvel string `json:"severityLevel"`
}

func (client BroadpeakClient) GetAllSources(offset uint, limit uint) ([]SourceOutput, error) {
	url := baseUrl + "sources"
	url = addOffsetUrl(url, offset, limit)
	resp, err := httpGetRequest(client, url)

	if err != nil {
		return []SourceOutput{}, err
	}
	var sourceApiResponse []SourceOutput

	err = deserialiseApiResponse(resp, &sourceApiResponse)
	if err != nil {
		return []SourceOutput{}, err
	}

	return sourceApiResponse, nil
}

func (client BroadpeakClient) CheckSourceStatus(sourceType string, sourceUrl string) (string, error) {
	url := fmt.Sprintf(baseUrl+"sources/%s/check?url=%s", sourceType, sourceUrl)
	resp, err := httpGetRequest(client, url)

	if err != nil {
		return "", err
	}
	return resp, nil
}
