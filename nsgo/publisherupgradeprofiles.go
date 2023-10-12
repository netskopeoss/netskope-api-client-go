package nsgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Struct that defines data returned when getting a list of publisher upgrade profiles

type PublisherUpgradeProfiles struct {
	UpgradeProfiles []struct {
		CreatedAt              string `json:"created_at"`
		DockerTag              string `json:"docker_tag"`
		Enabled                bool   `json:"enabled"`
		ExternalID             int    `json:"external_id"`
		Frequency              string `json:"frequency"`
		ID                     int    `json:"id"`
		Name                   string `json:"name"`
		NextUpdateTime         int    `json:"next_update_time"`
		NumAssociatedPublisher int    `json:"num_associated_publisher"`
		ReleaseType            string `json:"release_type"`
		Timezone               string `json:"timezone"`
		UpdatedAt              string `json:"updated_at"`
		UpgradingStage         int    `json:"upgrading_stage"`
		WillStart              bool   `json:"will_start"`
	} `json:"upgrade_profiles"`
}

type PublisherUpgradeProfile struct {
	CreatedAt      string `json:"created_at"`
	DockerTag      string `json:"docker_tag"`
	Enabled        bool   `json:"enabled"`
	Frequency      string `json:"frequency"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	NextUpdateTime int    `json:"next_update_time"`
	ReleaseType    string `json:"release_type"`
	Timezone       string `json:"timezone"`
	UpdatedAt      string `json:"updated_at"`
	UpgradingStage int    `json:"upgrading_stage"`
	WillStart      bool   `json:"will_start"`
}

type PublisherUpgradeProfileOptions struct {
	ExternalID  string `json:"external_id,omitempty"`  // Used when deleting a publisher upgrade profile
	ID          string `json:"id,omitempty"`           // Used when deleting a publisher upgrade profile
	Name        string `json:"name,omitempty"`         // Used when creating a publisher upgrade profile
	Timezone    string `json:"timezone,omitempty"`     // Used when creating a publisher upgrade profile
	ReleaseType string `json:"release_type,omitempty"` // Used when creating a publisher upgrade profile
	DockerTag   string `json:"docker_tag,omitempty"`   // Used when creating a publisher upgrade profile
	Frequency   string `json:"frequency,omitempty"`    // Used when creating a publisher upgrade profile
	Enabled     bool   `json:"enabled,omitempty"`      // Used when creating a publisher upgrade profile
}

// GetPublisherUpgradeProfiles function is used to build API request which is sent to sendRequest().
// It is called using the client struct, and returns an interface with a list of Publisher Upgrade Profiles.
// The output can be marshalled into a PublisherUpgradeProfiles struct.

func (c *Client) GetPublisherUpgradeProfiles() (interface{}, error) {
	//Setup the HTTP Request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/infrastructure/publisherupgradeprofiles", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	res := successResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == "success" {
		return res.Data, nil
	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}
}

// GetPublisherUpgradeProfileId function is used to build API request which is sent to sendRequest().

func (c *Client) GetPublisherUpgradeProfileId(options PublisherUpgradeProfileOptions) (*successResponse, error) {
	//Setup the HTTP Request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/infrastructure/publisherupgradeprofiles/%s", c.BaseURL, options.ExternalID), nil)
	if err != nil {
		return nil, err
	}

	res := successResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == "success" {
		jsonData, err := json.Marshal(res.Data)
		if err != nil {
			return nil, err
		}
		dataStruct := successResponse{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil
	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}
}

// CreatePublisherUpgradeProfile function is used to build API request which is sent to sendRequest().
// It is called using the client struct, and returns.

func (c *Client) CreatePublisherUpgradeProfile(options PublisherUpgradeProfileOptions) (*PublisherUpgradeProfile, error) {
	//Define JSON Body
	json_body, err := json.Marshal(options)
	if err != nil {
		return nil, errors.New("bad json options")
	}

	//Setup the HTTP Request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2/infrastructure/publisherupgradeprofiles", c.BaseURL), bytes.NewBuffer(json_body))
	if err != nil {
		return nil, err
	}

	res := successResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == "success" {
		jsonData, err := json.Marshal(res.Data)
		if err != nil {
			return nil, err
		}
		dataStruct := PublisherUpgradeProfile{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil
	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}

}

// UpdatePublisherUpgradeProfile function is used to build API request which is sent to sendRequest().

func (c *Client) UpdatePublisherUpgradeProfile(options PublisherUpgradeProfileOptions) (*PublisherUpgradeProfile, error) {
	//Define JSON Body
	json_body, err := json.Marshal(options)
	if err != nil {
		return nil, errors.New("bad json options")
	}

	//Setup the HTTP Request
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v2/infrastructure/publisherupgradeprofiles/%s", c.BaseURL, options.ID), bytes.NewBuffer(json_body))
	if err != nil {
		return nil, err
	}

	res := successResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == "success" {
		jsonData, err := json.Marshal(res.Data)
		if err != nil {
			return nil, err
		}
		dataStruct := PublisherUpgradeProfile{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil
	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}

}

// DeletePublisherUpgradeProfile function is used to build API request which is sent to sendRequest().

func (c *Client) DeletePublisherUpgradeProfile(options PublisherUpgradeProfileOptions) (*successResponse, error) {
	//Setup the HTTP Request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v2/infrastructure/publisherupgradeprofiles/%s", c.BaseURL, options.ExternalID), nil)
	if err != nil {
		return nil, err
	}

	res := successResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == "success" {
		jsonData, err := json.Marshal(res.Data)
		if err != nil {
			return nil, err
		}
		dataStruct := successResponse{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil
	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}
}
