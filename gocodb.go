package gocodb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func NewClient(dbUrl, apiToken string) *Client {
	return &Client{
		apiToken:   apiToken,
		dbUrl:      dbUrl,
		httpClient: &http.Client{},
	}
}

type Client struct {
	apiToken string
	dbUrl    string

	httpClient *http.Client
}

func (c *Client) SetHttpClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

func (c *Client) get(path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.dbUrl+path, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Add("xc-auth", c.apiToken)
	return c.httpClient.Do(req)
}

func (c *Client) post(path string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.dbUrl+path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Add("xc-auth", c.apiToken)
	req.Header.Add("Content-Type", "application/json")

	return c.httpClient.Do(req)
}

func (c *Client) patch(path string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodPatch, c.dbUrl+path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Add("xc-auth", c.apiToken)
	req.Header.Add("Content-Type", "application/json")

	return c.httpClient.Do(req)
}

func (c *Client) delete(path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, c.dbUrl+path, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Add("xc-auth", c.apiToken)
	return c.httpClient.Do(req)
}

func (c *Client) deleteWithBody(path string, body interface{}) (*http.Response, error) {
	// Marshal the body into JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	// Create a new DELETE request with the JSON body
	req, err := http.NewRequest(http.MethodDelete, c.dbUrl+path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set the required headers
	req.Header.Add("xc-auth", c.apiToken)
	req.Header.Add("Content-Type", "application/json")

	// Execute the request
	return c.httpClient.Do(req)
}
