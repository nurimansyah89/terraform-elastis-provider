package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *ElastisClient) Fetch(url string) ([]byte, error) {
	token := c.Token

	var client = &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set header
	request.Header.Set("apikey", token)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Error: %d", response.StatusCode))
	}

	result, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		type errResponseType struct {
			Errors map[string]interface{} `json:"errors"`
		}
		errResponse := &errResponseType{}
		err := json.Unmarshal(result, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("Error %d: %s", response.StatusCode, errResponse.Errors["Error"]))
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ElastisClient) Post(url string, payload []byte) ([]byte, error) {
	token := c.Token

	var client = &http.Client{}
	request, err := http.NewRequest("POST", url, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	// Set header
	request.Header.Set("apikey", token)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		type errResponseType struct {
			Errors map[string]interface{} `json:"errors"`
		}
		errResponse := &errResponseType{}
		err := json.Unmarshal(result, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("(%d): %s", response.StatusCode, errResponse.Errors["Error"]))
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ElastisClient) Patch(url string, payload []byte) ([]byte, error) {
	token := c.Token

	var client = &http.Client{}
	request, err := http.NewRequest("PATCH", url, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	// Set header
	request.Header.Set("apikey", token)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		type errResponseType struct {
			Errors map[string]interface{} `json:"errors"`
		}
		errResponse := &errResponseType{}
		err := json.Unmarshal(result, &errResponse)
		if err != nil {
			return nil, err
		}

		// For modify VM
		if errResponse.Errors["Error"] != "Not changing vcpu and memory because values are same." {
			return nil, errors.New(fmt.Sprintf("(%d): %s", response.StatusCode, errResponse.Errors["Error"]))
		}
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ElastisClient) Delete(url string, payload []byte) ([]byte, error) {
	token := c.Token

	var client = &http.Client{}
	request, err := http.NewRequest("DELETE", url, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	// Set header
	request.Header.Set("apikey", token)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		type errResponseType struct {
			Errors map[string]interface{} `json:"errors"`
		}
		errResponse := &errResponseType{}
		err := json.Unmarshal(result, &errResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("(%d): %s", response.StatusCode, errResponse.Errors["Error"]))
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}
