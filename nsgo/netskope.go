package nsgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

//The client struct defines a new HttpClient with the required connection details.
//BaseURL is a string that represents the Netskope tenant URL. (i.e. "https://example-tenant.goskope.com")
//apiToken is a string that represents the Netskope API v2 Token.
type Client struct {
	BaseURL    string
	apiToken   string
	HttpClient *http.Client
}

//The errorResponse struct defines an error response sent by the API.
//
type errorResponse struct {
	Message string `json:"message"`
}

//The successResponse struct defines an successful response sent by the API.
//
type successResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Total   int         `json:"total,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Message string      `json:"message,omitempty"`
}

//The ipsecResponse struct defines an successful response sent by the IPSec API Endpoint.
//
type ipsecResponse struct {
	Status  int         `json:"status"`
	Total   int         `json:"total,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Message string      `json:"message,omitempty"`
}

//The NewClient function accepts the BaseURL and apiToken and returns a client struct.
func NewClient(BaseURL, apiToken string) *Client {
	return &Client{
		BaseURL:  BaseURL,
		apiToken: apiToken,
		HttpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

//The sendRequest function is used to package an API request and send it to the defined Netskope tenant(BaseURL).
//It is called using the client struct, takes an http.Request as input and returns an interface.
func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Netskope-Api-Token", c.apiToken)

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	fullResponse := v

	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}
	return nil
}
