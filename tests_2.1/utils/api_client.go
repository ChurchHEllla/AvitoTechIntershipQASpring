package utils

import (
	"bytes"
	"net/http"
)

type APIClient struct {
	baseURL string
	client  *http.Client
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (c *APIClient) doRequest(method, endpoint string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, c.baseURL+endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return c.client.Do(req)
}

func (c *APIClient) Get(endpoint string) (*http.Response, error) {
	return c.doRequest("GET", endpoint, nil)
}

func (c *APIClient) Post(endpoint string, data []byte) (*http.Response, error) {
	return c.doRequest("POST", endpoint, data)
}

func (c *APIClient) Delete(endpoint string) (*http.Response, error) {
	return c.doRequest("DELETE", endpoint, nil)
}