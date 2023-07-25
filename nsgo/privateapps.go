package nsgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type PrivateAppOptions struct {
	Name string `json:"name,omitempty"`
	Id   string `json:"id,omitempty"`
}

type PrivateAppsList struct {
	PrivateApps []struct {
		AppID              int    `json:"app_id"`
		AppName            string `json:"app_name"`
		ClientlessAccess   bool   `json:"clientless_access"`
		Host               string `json:"host"`
		PrivateAppProtocol string `json:"private_app_protocol"`
		Protocols          []struct {
			CreatedAt time.Time `json:"created_at"`
			ID        int       `json:"id"`
			Port      string    `json:"port"`
			ServiceID int       `json:"service_id"`
			Transport string    `json:"transport"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"protocols"`
		Reachability struct {
			ErrorCode   int    `json:"-"`
			ErrorString string `json:"-"`
			Reachable   bool   `json:"reachable"`
		} `json:"reachability"`
		//Reachability                interface{} `json:"-"`
		ServicePublisherAssignments []struct {
			Primary     string `json:"primary"`
			PublisherID int    `json:"publisher_id"`
			//Reachability string `json:"reachability"`
			Reachability struct {
				ErrorCode   int    `json:"-"`
				ErrorString string `json:"-"`
				Reachable   bool   `json:"reachable"`
			} `json:"reachability"`
			ServiceID int `json:"service_id"`
		} `json:"service_publisher_assignments"`
		TrustSelfSignedCerts bool `json:"trust_self_signed_certs"`
		UsePublisherDNS      bool `json:"use_publisher_dns"`
	} `json:"private_apps"`
}

type PrivateApp struct {
	AppName              string              `json:"app_name"`
	Id                   int                 `json:"id",omitempty`
	Host                 string              `json:"host"`
	Protocols            []Protocol          `json:"protocols"`
	Publishers           []PublisherIdentity `json:"publishers,omitempty"`
	Tags                 []PrivateAppTags    `json:"tags,omitempty"`
	UsePublisherDNS      bool                `json:"use_publisher_dns,omitempty"`
	ClientlessAccess     bool                `json:"clientless_access,omitempty"`
	TrustSelfSignedCerts bool                `json:"trust_self_signed_certs,omitempty"`
}

type Protocol struct {
	Type string `json:"type"`
	Port string `json:"port"`
}

type PublisherIdentity struct {
	PublisherID   string `json:"publisher_id"`
	PublisherName string `json:"publisher_name"`
}

type PrivateAppTags struct {
	TagName string `json:"tag_name"`
}

func (c *Client) GetPrivateApps() (interface{}, error) {
	//Setup the HTTP Request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/steering/apps/private", c.BaseURL), nil)
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

func (c *Client) GetPrivateAppsWithFilter(filter string) (interface{}, error) {
	//Escape Filter
	filter = url.QueryEscape(filter)

	//Setup the HTTP Request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/steering/apps/private?query=%s", c.BaseURL, filter), nil)
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
		return res.Data, nil
	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}
}

func (c *Client) GetPrivateAppId(options PrivateAppOptions) (interface{}, error) {
	//Setup the HTTP Request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/steering/apps/private/%s", c.BaseURL, options.Id), nil)
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

func (c *Client) CreatePrivateApp(privateapp PrivateApp) (*PrivateApp, error) {
	//Define JSON Body
	json_body, err := json.Marshal(privateapp)
	if err != nil {
		return nil, errors.New("bad json options")
	}

	//Setup the HTTP Request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2/steering/apps/private", c.BaseURL), bytes.NewBuffer(json_body))
	if err != nil {
		return nil, err
	}

	res := successResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == "success" {
		//return res.Data, nil
		jsonData, err := json.Marshal(res.Data)
		if err != nil {
			return nil, err
		}
		dataStruct := PrivateApp{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil

	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}

}

func (c *Client) DeletePrivateApp(options PrivateAppOptions) (*successResponse, error) {
	//Setup the HTTP Request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v2/steering/apps/private/%s", c.BaseURL, options.Id), nil)
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

func (c *Client) UpdatePrivateApp(options PrivateAppOptions, privateapp PrivateApp) (*PrivateApp, error) {
	//Define JSON Body
	json_body, err := json.Marshal(privateapp)
	if err != nil {
		return nil, errors.New("bad json options")
	}

	//Setup the HTTP Request
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/v2/steering/apps/private/%s", c.BaseURL, options.Id), bytes.NewBuffer(json_body))
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
		dataStruct := PrivateApp{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil
	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}

}

func (c *Client) ReplacePrivateApp(options PrivateAppOptions, privateapp PrivateApp) (*PrivateApp, error) {
	//Define JSON Body
	json_body, err := json.Marshal(privateapp)
	if err != nil {
		return nil, errors.New("bad json options")
	}

	//Setup the HTTP Request
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v2/steering/apps/private/%s", c.BaseURL, options.Id), bytes.NewBuffer(json_body))
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
		dataStruct := PrivateApp{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil
	} else if res.Status == "error" {
		return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + res.Status)
	}

}
