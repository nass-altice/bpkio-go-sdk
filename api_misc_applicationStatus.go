package broadpeakio

func (client BroadpeakClient) CheckApplicationStatus() (string, error) {
	url := baseUrl + "status"

	resp, err := httpGetRequest(client, url)

	if err != nil {
		return "", err
	}

	return resp, nil
}
