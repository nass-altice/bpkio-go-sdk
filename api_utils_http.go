package broadpeakio

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
)

func httpGetRequest(client BroadpeakClient, url string) (string, error) {

	bearer := "Bearer " + client.apiKey

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	httpClient := &http.Client{}

	httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return "", err
	} else {
		defer resp.Body.Close()
		data, _ := io.ReadAll(resp.Body)
		if resp.StatusCode >= 300 || resp.StatusCode < 200 {
			errMsg := errors.New(resp.Status + string(data))
			return string(data), errMsg
		}
		return string(data), nil
	}
}

func httpPostRequest(client BroadpeakClient, url string, body string) (string, error) {

	bearer := "Bearer " + client.apiKey

	requestBody := []byte(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}

	httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return "", err
	} else {
		defer resp.Body.Close()
		data, _ := io.ReadAll(resp.Body)
		if resp.StatusCode >= 300 || resp.StatusCode < 200 {
			errMsg := errors.New(resp.Status + string(data))
			return string(data), errMsg
		}
		return string(data), err
	}
}

func httpDeleteRequest(client BroadpeakClient, url string) (string, error) {
	bearer := "Bearer " + client.apiKey

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	httpClient := &http.Client{}

	httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return "", err
	} else {
		defer resp.Body.Close()
		data, _ := io.ReadAll(resp.Body)
		if resp.StatusCode >= 300 || resp.StatusCode < 200 {
			errMsg := errors.New(resp.Status + string(data))
			return string(data), errMsg
		}
		return string(data), nil
	}
}

func httpPutRequest(client BroadpeakClient, url string, body string) (string, error) {

	bearer := "Bearer " + client.apiKey

	requestBody := []byte(body)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}

	httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return "", err
	} else {
		defer resp.Body.Close()
		data, _ := io.ReadAll(resp.Body)
		if resp.StatusCode >= 300 || resp.StatusCode < 200 {
			errMsg := errors.New(resp.Status + string(data))
			return string(data), errMsg
		}
		return string(data), nil
	}
}
