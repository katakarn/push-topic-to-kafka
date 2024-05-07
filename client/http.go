package client

import (
	"bytes"
	"io"
	"net/http"

	"github.com/sony/gobreaker"
)

func Request(method string, url string, token string, requestBody []byte, contentType string, gb *gobreaker.CircuitBreaker) ([]byte, int, error) {
	var statusCode int
	body, err := gb.Execute(func() (interface{}, error) {
		var bodyReader *bytes.Reader
		client := http.Client{}
		if requestBody != nil {
			bodyReader = bytes.NewReader(requestBody)
		}

		req, err := httpNewRequest(method, url, bodyReader)
		if err != nil {
			return nil, err
		}

		if contentType != "" {
			req.Header.Set("content-type", contentType)
		}

		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		statusCode = resp.StatusCode

		return body, nil
	})

	if err != nil {
		return nil, statusCode, err
	}
	return body.([]byte), statusCode, nil
}

func httpNewRequest(method string, url string, bodyReader *bytes.Reader) (*http.Request, error) {
	if bodyReader != nil {
		return http.NewRequest(method, url, bodyReader)
	}
	return http.NewRequest(method, url, nil)
}
