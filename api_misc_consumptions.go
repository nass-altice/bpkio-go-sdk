package broadpeakio

type AggregateConsumptionOutput struct {
	ServiceIds         []uint  `json:"serviceIds,omitempty"`
	From               string  `json:"from,omitempty"`
	To                 string  `json:"to,omitempty"`
	Egress             float32 `json:"egress,omitempty"`
	VirtualChannel     float32 `json:"virtualChannel,omitempty"`
	ContentReplacement float32 `json:"contentReplacement,omitempty"`
	InsertedAds        float32 `json:"insertedAds,omitempty"`
}

type AggregateConsumptionInput struct {
	ServiceIds []uint `json:"serviceIds,omitempty"`
	From       string `json:"from,omitempty"`
	To         string `json:"to,omitempty"`
}

func (client BroadpeakClient) GetAggregateConsumption(options AggregateConsumptionInput) (AggregateConsumptionOutput, error) {
	url := baseUrl + "consumption"

	resp, err := httpGetRequest(client, url)

	if err != nil {
		return AggregateConsumptionOutput{}, err
	}

	var consumptionApiResponse AggregateConsumptionOutput

	err = deserialiseApiResponse(resp, &consumptionApiResponse)
	if err != nil {
		return AggregateConsumptionOutput{}, err
	}

	return consumptionApiResponse, nil
}
