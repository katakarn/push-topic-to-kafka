package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (c *CircuitBreakerProxy) GetPreRegistrationHistory(baseUrl string) func(requestBody GetPreRegistrationHistoryRequest, token string) ([]GetPreRegistrationHistoryResponse, error) {
	return func(requestBody GetPreRegistrationHistoryRequest, token string) ([]GetPreRegistrationHistoryResponse, error) {
		resp := []GetPreRegistrationHistoryResponse{}
		url := baseUrl + "/preregistration/history"
		requestBodyByte, err := json.Marshal(requestBody)
		if err != nil {
			return resp, err
		}
		body, statusCode, err := Request("POST", url, token, requestBodyByte, "application/json", c.Gb)
		if err != nil {
			return resp, err
		}

		if statusCode != http.StatusOK {
			c.Logger.Error(fmt.Sprintf("error http client status code %d get: %s", statusCode, url))
			return resp, errors.New("can not get pre registration history")
		}

		err = json.Unmarshal(body, &resp)
		if err != nil {
			return resp, err
		}
		return resp, nil
	}
}

func (c *CircuitBreakerProxy) GetRegistrationHistory(baseUrl string) func(requestBody GetRegistrationHistoryRequest, token string) ([]GetRegistrationHistoryResponse, error) {
	return func(requestBody GetRegistrationHistoryRequest, token string) ([]GetRegistrationHistoryResponse, error) {
		resp := []GetRegistrationHistoryResponse{}
		url := baseUrl + "/registration/history"
		requestBodyByte, err := json.Marshal(requestBody)
		if err != nil {
			return resp, err
		}
		body, statusCode, err := Request("POST", url, token, requestBodyByte, "application/json", c.Gb)
		if err != nil {
			return resp, err
		}

		if statusCode != http.StatusOK {
			c.Logger.Error(fmt.Sprintf("error http client status code %d get: %s", statusCode, url))
			return resp, errors.New("can not get registration history")
		}

		err = json.Unmarshal(body, &resp)
		if err != nil {
			return resp, err
		}
		return resp, nil
	}
}
