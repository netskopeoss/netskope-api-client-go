package nsgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

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

type IpsecOptions struct {
	Phase1 IpsecPhase1 `json:"phase1"`
	Phase2 IpsecPhase2 `json:"phase2"`
}

type IpsecPhase1 struct {
	Dhgroup        string `json:"dhgroup"`
	Dpd            bool   `json:"-"`
	Encryptionalgo string `json:"encryptionalgo"`
	Ikeversion     string `json:"ikeversion"`
	Integrityalgo  string `json:"integrityalgo"`
	Salifetime     string `json:"salifetime"`
}

type IpsecPhase2 struct {
	Dhgroup        string `json:"dhgroup"`
	Encryptionalgo string `json:"encryptionalgo"`
	Integrityalgo  string `json:"integrityalgo"`
	Pfs            bool   `json:"-"`
	Salifetime     string `json:"salifetime"`
}

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

type NewIpsecTunnel struct {
	Encryption    string   `json:"encryption"`
	Site          string   `json:"site"`
	Srcidentity   string   `json:"srcidentity"`
	Srcipidentity string   `json:"srcipidentity,omitempty"`
	Psk           string   `json:"psk"`
	Notes         string   `json:"notes,omitempty"`
	Sourcetype    string   `json:"sourcetype,omitempty"` //['User', 'Server', 'IoT', 'Guest wifi', 'Mixed']
	Pops          []string `json:"pops"`
	Bandwidth     int      `json:"bandwidth"` //[50, 100, 150, 250]
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
