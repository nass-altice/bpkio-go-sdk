package broadpeakio

type BroadpeakClient struct {
	apiKey string
}

type Identifiable struct {
	Id uint `json:"id,omitempty"`
}

type ApiResponseOutput struct {
	StatusCode uint   `json:"statusCode"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

func MakeClient(apiKey string) BroadpeakClient {
	client := BroadpeakClient{apiKey}
	return client
}
