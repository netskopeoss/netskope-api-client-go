package nsgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

//The client struct defines a new HttpClient with the required connection details.
//BaseURL is a string that represents the Netskope tenant URL. (i.e. "https://example-tenant.goskope.com")
//apiToken is a string that represents the Netskope API v2 Token.
type Client struct {
	BaseURL    string
	apiToken   string
	HttpClient *http.Client
}

//RequestOptions defines a struct to pass options to functions.
type RequestOptions struct {
	Name string `json:"name,omitempty"`
	Id   string `json:"id,omitempty"`
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

type PopFilters struct {
	Name    string `url:"name,omitempty" validate:"excluded_with=Region Country Lat Long Ip Fields"`
	Region  string `url:"region,omitempty" validate:"excluded_with=Name Country Lat Long Ip Fields"`
	Country string `url:"country,omitempty" validate:"excluded_with=Name Region Lat Long Ip Fields"`
	Lat     string `url:"lat,omitempty" validate:"required_with=Long,excluded_with=Name Region Country Ip Fields"`
	Long    string `url:"long,omitempty" validate:"required_with=Lat,excluded_with=Name Region Country Ip Fields"`
	Ip      string `url:"ip,omitempty" validate:"excluded_with=Name Region Country Lat Long Fields"`
	Fields  string `url:"fields,omitempty" validate:"excluded_with=Name Region Country Lat Long Ip"`
	Offset  int    `url:"offset,omitempty"`
	Limit   int    `url:"limit,omitempty"`
}

var defaultRetry = &RetryConfig{
	RetryMax:     100,
	RetryWaitMin: 5,
	RetryWaitMax: 20,
}

type RetryConfig struct {
	RetryMax     int
	RetryWaitMin int
	RetryWaitMax int
	Logger       interface{}
}

type Config struct {
	BaseURL     string
	ApiToken    string
	RetryConfig *RetryConfig
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

//The NewRetryClient function accepts the BaseURL and apiToken and returns a retryableclient.
//Use this in place of NewClient to enable automatic retry logic for rate limiting etc.
func NewRetryClient(config Config) *Client {

	retryClient := retryablehttp.NewClient()

	if config.RetryConfig != nil {
		retryClient.RetryMax = config.RetryConfig.RetryMax
		retryClient.RetryWaitMin = time.Second * time.Duration(config.RetryConfig.RetryWaitMin)
		retryClient.RetryWaitMax = time.Second * time.Duration(config.RetryConfig.RetryWaitMax)
		retryClient.Logger = config.RetryConfig.Logger

	} else {
		retryClient.RetryMax = defaultRetry.RetryMax
		retryClient.RetryWaitMin = time.Second * time.Duration(defaultRetry.RetryWaitMin)
		retryClient.RetryWaitMax = time.Second * time.Duration(defaultRetry.RetryWaitMax)
		retryClient.Logger = defaultRetry.Logger
	}

	return &Client{
		BaseURL:    config.BaseURL,
		apiToken:   config.ApiToken,
		HttpClient: retryClient.StandardClient(),
	}
}

//The sendRequest function is used to package an API request and send it to the defined Netskope tenant(BaseURL).
//It is called using the client struct, takes an http.Request as input and returns an interface.
func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("user-Agent", "nsgo-api-client/0.3.0")
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
