package nsgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

//IpsecPops defines a struct used for list of Netskope IPSec PoPs returned from the tenant.
type IpsecPops []struct {
	Closestpop bool         `json:"closestpop"`
	Gateway    string       `json:"gateway"`
	ID         int          `json:"id"`
	Location   string       `json:"location"`
	Name       string       `json:"name"`
	Options    IpsecOptions `json:"-"`
	Probeip    string       `json:"probeip"`
	Region     string       `json:"region"`
}

//IpsecOptions Defines a struct to use return IPSec tunnel options.
type IpsecOptions struct {
	Phase1 IpsecPhase1 `json:"phase1"`
	Phase2 IpsecPhase2 `json:"phase2"`
}

//IpsecPhase1 Defines a struct used to return Phase 1 tunnel options.
type IpsecPhase1 struct {
	Dhgroup        string `json:"dhgroup"`
	Dpd            bool   `json:"-"`
	Encryptionalgo string `json:"encryptionalgo"`
	Ikeversion     string `json:"ikeversion"`
	Integrityalgo  string `json:"integrityalgo"`
	Salifetime     string `json:"salifetime"`
}

//IpsecPhase2 Defines a struct used to return Phase 2 tunnel options.
type IpsecPhase2 struct {
	Dhgroup        string `json:"dhgroup"`
	Encryptionalgo string `json:"encryptionalgo"`
	Integrityalgo  string `json:"integrityalgo"`
	Pfs            bool   `json:"-"`
	Salifetime     string `json:"salifetime"`
}

//IpsecTunnels defines a struct to return a list of IPSec tunnels.
type IpsecTunnels []struct {
	ID      int    `json:"id"`
	Site    string `json:"site"`
	Enabled bool   `json:"enabled"`
	Pops    []struct {
		Name    string `json:"name"`
		Gateway string `json:"gateway"`
		Probeip string `json:"probeip"`
		Primary bool   `json:"primary"`
	} `json:"pops"`
	Status struct {
		Status     string `json:"status"`
		Since      string `json:"since"`
		Throughput string `json:"throughput"`
	} `json:"status"`
	Template      string `json:"template"`
	Sourcetype    string `json:"sourcetype"`
	Notes         string `json:"notes"`
	Encryption    string `json:"encryption"`
	Srcidentity   string `json:"srcidentity"`
	Srcipidentity string `json:"srcipidentity"`
}

//NewIpsecTunnel defines a struct for creating an IPSec tunnel in Netskope.
type NewIpsecTunnel struct {
	Encryption    string        `json:"encryption,omitempty"`
	Site          string        `json:"site,omitempty"`
	Srcidentity   string        `json:"srcidentity,omitempty"`
	Srcipidentity string        `json:"srcipidentity,omitempty"`
	Psk           string        `json:"psk,omitempty"`
	Notes         string        `json:"notes,omitempty"`
	Sourcetype    string        `json:"sourcetype,omitempty"` //['User', 'Server', 'IoT', 'Guest wifi', 'Mixed']
	Pops          []interface{} `json:"pops,omitempty"`
	Bandwidth     int           `json:"bandwidth,omitempty"` //[50, 100, 150, 250]
}

func (c *Client) GetIpsecPops() (*IpsecPops, error) {
	//Setup the HTTP Request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/steering/ipsec/pops", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	res := ipsecResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == 200 {
		jsonData, err := json.Marshal(res.Result)
		if err != nil {
			return nil, err
		}
		dataStruct := IpsecPops{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil

		//return res.Result, nil
		//} else if res.Status == "error" {
		//	return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + strconv.Itoa(res.Status))
	}
}

//GetIpsecPopId function is used to GET an individual Pop by ID.
func (c *Client) GetIpsecPopId(options RequestOptions) (*IpsecPops, error) {
	//Setup the HTTP Request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/steering/ipsec/pops/%s", c.BaseURL, options.Id), nil)

	if err != nil {
		return nil, err
	}

	res := ipsecResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == 200 {
		jsonData, err := json.Marshal(res.Result)
		if err != nil {
			return nil, err
		}
		dataStruct := IpsecPops{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil

		//return res.Result, nil
	} else {
		return nil, errors.New("Unkown Status: " + strconv.Itoa(res.Status))
	}
}

//GetIpsecTunnels defines a function to get a list of IPSec Tunnels from a Netskope tennant.
func (c *Client) GetIpsecTunnels() (*IpsecTunnels, error) {
	//Setup the HTTP Request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/steering/ipsec/tunnels", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	res := ipsecResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == 200 {
		jsonData, err := json.Marshal(res.Result)
		if err != nil {
			return nil, err
		}
		dataStruct := IpsecTunnels{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil

		//return res.Result, nil
		//} else if res.Status == "error" {
		//	return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + strconv.Itoa(res.Status))
	}
}

//GetIpsecTunnelId function is used to GET an individual Tunnel by ID.
func (c *Client) GetIpsecTunnelId(options RequestOptions) (*IpsecTunnels, error) {
	//Setup the HTTP Request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/steering/ipsec/tunnels/%s", c.BaseURL, options.Id), nil)

	if err != nil {
		return nil, err
	}

	res := ipsecResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == 200 {
		jsonData, err := json.Marshal(res.Result)
		if err != nil {
			return nil, err
		}
		dataStruct := IpsecTunnels{}
		json.Unmarshal(jsonData, &dataStruct)
		return &dataStruct, nil

		//return res.Result, nil
		//} else if res.Status == "error" {
		//	return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + strconv.Itoa(res.Status))
	}
}

//CreateIpsecTunnel defines a function to create a new IPSec Tunnel in a Netskope tennant.
func (c *Client) CreateIpsecTunnel(ipsectunnel NewIpsecTunnel) (interface{}, error) {
	json_body, err := json.Marshal(ipsectunnel)
	if err != nil {
		return nil, errors.New("bad json options")
	}

	//Setup the HTTP Request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2/steering/ipsec/tunnels", c.BaseURL), bytes.NewBuffer(json_body))
	if err != nil {
		return nil, err
	}

	res := ipsecResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == 201 {
		/*
			jsonData, err := json.Marshal(res.Result)
			if err != nil {
				return nil, err
			}
			dataStruct := IpsecTunnels{}
			json.Unmarshal(jsonData, &dataStruct)
			return &dataStruct, nil
		*/
		return res.Result, nil
		//} else if res.Status == "error" {
		//	return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + strconv.Itoa(res.Status))
	}
}

//UpdateIpsecTunnel defines a function to create a new IPSec Tunnel in a Netskope tennant.
func (c *Client) UpdateIpsecTunnel(options RequestOptions, ipsectunnel NewIpsecTunnel) (interface{}, error) {
	json_body, err := json.Marshal(ipsectunnel)
	if err != nil {
		return nil, errors.New("bad json options")
	}

	//Setup the HTTP Request
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/v2/steering/ipsec/tunnels/%s", c.BaseURL, options.Id), bytes.NewBuffer(json_body))
	if err != nil {
		return nil, err
	}

	res := ipsecResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == 200 {
		/*
			jsonData, err := json.Marshal(res.Result)
			if err != nil {
				return nil, err
			}
			dataStruct := IpsecTunnels{}
			json.Unmarshal(jsonData, &dataStruct)
			return &dataStruct, nil
		*/
		return res.Result, nil
		//} else if res.Status == "error" {
		//	return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + strconv.Itoa(res.Status))
		//return nil, errors.New("Unkown Status: " + res)
	}
}

//DeleteIpsecTunnel defines a function to create a new IPSec Tunnel in a Netskope tennant.
func (c *Client) DeleteIpsecTunnel(options RequestOptions) (interface{}, error) {
	//Setup the HTTP Request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v2/steering/ipsec/tunnels/%s", c.BaseURL, options.Id), nil)
	if err != nil {
		return nil, err
	}

	res := ipsecResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	if res.Status == 200 {
		/*
			jsonData, err := json.Marshal(res.Result)
			if err != nil {
				return nil, err
			}
			dataStruct := IpsecTunnels{}
			json.Unmarshal(jsonData, &dataStruct)
			return &dataStruct, nil
		*/
		return res.Result, nil
		//} else if res.Status == "error" {
		//	return nil, errors.New(res.Message)
	} else {
		return nil, errors.New("Unkown Status: " + strconv.Itoa(res.Status))
	}
}
