package nsgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// PublisherList struct is used to define a list of Netskope publishers returned from the GET method.

type PublishersList struct {
	Publishers []struct {
		Assessment struct {
			EeeSupport string `json:"eee_support"`
			HddFree    string `json:"hdd_free"`
			HddTotal   string `json:"hdd_total"`
			IPAddress  string `json:"ip_address"`
			Latency    string `json:"latency"`
			Version    string `json:"version"`
		} `json:"assessment"`
		CommonName                         string   `json:"common_name"`
		Lbrokerconnect                     bool     `json:"lbrokerconnect"`
		PublisherID                        int      `json:"publisher_id"`
		PublisherName                      string   `json:"publisher_name"`
		PublisherUpgradeProfilesExternalID int      `json:"publisher_upgrade_profiles_external_id"`
		Registered                         bool     `json:"registered"`
		Status                             string   `json:"status"`
		StitcherID                         int      `json:"stitcher_id"`
		Tags                               []string `json:"tags"`
		UpgradeFailedReason                struct {
			Detail    string `json:"detail"`
			ErrorCode string `json:"error_code"`
			Timestamp string `json:"timestamp"`
			Version   string `json:"version"`
		} `json:"upgrade_failed_reason"`
		UpgradeRequest bool `json:"upgrade_request"`
		UpgradeStatus  struct {
			StatusFailureCode string `json:"status_failure_code"`
			Upstat            string `json:"upstat"`
		} `json:"upgrade_status,omitempty"`
	} `json:"publishers"`
}

// Publisher is a struct used to define and individual Netskope Publisher.

type Publisher struct {
	Assessment Assessment `json:"assessment"`
	CommonName string     `json:"common_name"`
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Registered bool       `json:"registered"`
	Status     string     `json:"status"`
	StitcherID int        `json:"stitcher_id"`
}

// Assessment is a struct used inside of the Publisher struct.
// BUG(terraform-provider-netskope): Need tp modify EeeSupport so it isn't returned in JSON.
type Assessment struct {
	EeeSupport bool   `json:"eee_support"`
	HddFree    string `json:"hdd_free"`
	HddTotal   string `json:"hdd_total"`
	IPAddress  string `json:"ip_address"`
	Version    string `json:"version"`
}

// PublisherOptions struct defines details used in GET by ID, Create, Update and Delete methods.
//
// - Name: a string that represents the publisher name
//
// - Id: a string that represents the publisher Id
//
//	newpublisher := nsgo.PublisherOptions{
//		Name: "MyNewPublisher",
//	}
//
//	updatepublisher := nsgo.PublisherOptions{
//		Id: "987",
//	}
type PublisherOptions struct {
	Name string `json:"name,omitempty"`
	Id   string `json:"id,omitempty"`
}

// PublisherToken struct is used to define the token response data.
type PublisherToken struct {
	Token string `json:"token"`
}

// GetPublishers function is used to build API request which is sent to sendRequest().
// It is called using the client struct, and returns an interface with a list of Publishers.
// The interface can be marshaled into the PublishersList struct.
//
// BUG(terraform-provider-netskope): Need tp modify the Assessment struct so that this request can return a PublishersList struct instead of an interface.
func (c *Client) GetPublishers() (interface{}, error) {
	//Setup the HTTP Request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/infrastructure/publishers", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	res := successResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == "success" {
		/*
			jsonData, err := json.Marshal(res.Data)
			if err != nil {
				return nil, err
			}
			dataStruct := PublishersList{}
			json.Unmarshal(jsonData, &dataStruct)
			return &dataStruct, nil
		*/
		return res.Data, nil
	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}
}

// GetPublishersWithFilters function is used to build API request which is sent to sendRequest().
// It is called using the client struct and the NpaFilters Struct. It returns an interface with a list of Filtered Publishers.
// The interface can be marshaled into the PublishersList struct.
//
// BUG(terraform-provider-netskope): Need tp modify the Assessment struct so that this request can return a PublishersList struct instead of an interface.
func (c *Client) GetPublishersWithFilter(filter string) (interface{}, error) {
	//Escape Filter
	filter = url.QueryEscape(filter)
	//Setup the HTTP Request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/infrastructure/publishers?query=%s", c.BaseURL, filter), nil)
	if err != nil {
		return nil, err
	}

	//Debug
	//reqDump, err := httputil.DumpRequestOut(req, true)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(reqDump))

	res := successResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == "success" {
		/*
			jsonData, err := json.Marshal(res.Data)
			if err != nil {
				return nil, err
			}
			dataStruct := PublishersList{}
			json.Unmarshal(jsonData, &dataStruct)
			return &dataStruct, nil
		*/
		return res.Data, nil
	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}
}

// GetPublisherId function is used to build API request which is sent to sendRequest().
// It is called using the client struct, takes and returns an interface.
func (c *Client) GetPublisherId(options PublisherOptions) (*Publisher, error) {
	//Setup the HTTP Request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/infrastructure/publishers/%s", c.BaseURL, options.Id), nil)
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
		dataStruct := Publisher{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil
	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}
}

func (c *Client) CreatePublisher(options PublisherOptions) (*Publisher, error) {
	//Define JSON Body
	json_body, err := json.Marshal(options)
	if err != nil {
		return nil, errors.New("bad json options")
	}

	//Setup the HTTP Request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2/infrastructure/publishers", c.BaseURL), bytes.NewBuffer(json_body))
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
		dataStruct := Publisher{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil
	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}

}

func (c *Client) GetToken(options PublisherOptions) (*PublisherToken, error) {
	//Define JSON Body

	//Setup the HTTP Request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2/infrastructure/publishers/%s/registration_token", c.BaseURL, options.Id), nil)
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
		dataStruct := PublisherToken{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil

	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}

}

func (c *Client) DeletePublisher(options PublisherOptions) (*successResponse, error) {
	//Setup the HTTP Request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v2/infrastructure/publishers/%s", c.BaseURL, options.Id), nil)
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

func (c *Client) UpdatePublisher(options PublisherOptions) (interface{}, error) {
	//Define JSON Body
	json_body, err := json.Marshal(options)
	if err != nil {
		return nil, errors.New("bad json options")
	}

	//Setup the HTTP Request
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/v2/infrastructure/publishers/%s", c.BaseURL, options.Id), bytes.NewBuffer(json_body))
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
		dataStruct := Publisher{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil
	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}

}

func (c *Client) ReplacePublisher(options PublisherOptions) (interface{}, error) {
	//Define JSON Body
	json_body, err := json.Marshal(options)
	if err != nil {
		return nil, errors.New("bad json options")
	}

	//Setup the HTTP Request
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v2/infrastructure/publishers/%s", c.BaseURL, options.Id), bytes.NewBuffer(json_body))
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
		dataStruct := Publisher{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil
	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}

}
