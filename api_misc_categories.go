package broadpeakio

import (
	"encoding/json"
	"fmt"
)

type Category struct {
	Name          string        `json:"name,omitempty"`
	Subcategories []Subcategory `json:"subcategories,omitempty"`
}

type Subcategory struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type CategoryOutput struct {
	Name        string        `json:"name"`
	Subcategory []Subcategory `json:"subcategories"`
	Id          uint          `json:"id"`
}

func (client BroadpeakClient) CreateCategory(name string, subcategories []Subcategory) (CategoryOutput, error) {
	url := baseUrl + "categories"
	newCategory := Category{name, subcategories}

	newCategoryJson, _ := json.Marshal(newCategory)

	resp, err := httpPostRequest(client, url, string(newCategoryJson))

	if err != nil {
		return CategoryOutput{}, err
	}

	var categoryApiResponse CategoryOutput

	err = deserialiseApiResponse(resp, &categoryApiResponse)

	if err != nil {
		return CategoryOutput{}, err
	}

	return categoryApiResponse, nil
}

func (client BroadpeakClient) GetAllCategories(offset uint, limit uint) ([]CategoryOutput, error) {
	url := baseUrl + "categories"
	url = addOffsetUrl(url, offset, limit)
	resp, err := httpGetRequest(client, url)

	if err != nil {
		return []CategoryOutput{}, err
	}
	var categoryApiResponse []CategoryOutput

	err = deserialiseApiResponse(resp, &categoryApiResponse)

	if err != nil {
		return []CategoryOutput{}, err
	}

	return categoryApiResponse, nil
}

func (client BroadpeakClient) GetCategory(id uint) (CategoryOutput, error) {
	url := fmt.Sprintf(baseUrl+"categories/%d", id)
	resp, err := httpGetRequest(client, url)

	if err != nil {
		return CategoryOutput{}, err
	}
	var categoryApiResponse CategoryOutput

	err = deserialiseApiResponse(resp, &categoryApiResponse)

	if err != nil {
		return CategoryOutput{}, err
	}

	return categoryApiResponse, nil
}

func (client BroadpeakClient) DeleteCategory(id uint) (string, error) {
	url := fmt.Sprintf(baseUrl+"categories/%d", id)
	resp, err := httpDeleteRequest(client, url)

	if err != nil {
		return "", err
	}
	return resp, nil
}

func (client BroadpeakClient) UpdateCategory(id uint, name string, subcategories []Subcategory) (CategoryOutput, error) {
	url := fmt.Sprintf(baseUrl+"categories/%d", id)
	newCategory := Category{name, subcategories}

	newCategoryJson, _ := json.Marshal(newCategory)

	resp, err := httpPutRequest(client, url, string(newCategoryJson))

	if err != nil {
		return CategoryOutput{}, err
	}

	var categoryApiResponse CategoryOutput

	err = deserialiseApiResponse(resp, &categoryApiResponse)

	if err != nil {
		return CategoryOutput{}, err
	}

	return categoryApiResponse, nil
}
