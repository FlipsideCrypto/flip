package api

import (
	"encoding/json"
)

// Dataset is a set of rules that define attributes over a distinct set of members
type DataExchangeCredsResp struct {
	Url         string `json:"url"`
	Account     string `json:"account"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Database    string `json:"database"`
	Warehouse   string `json:"warehouse"`
	Region      string `json:"region"`
	Role        string `json:"role"`
	DockerImage string `json:"docker_image"`
}

func (c Client) GetDataExchangeCreds() (*DataExchangeCredsResp, error) {
	res, err := c.Get("/data-exchange/creds")
	if err != nil {
		return nil, err
	}

	var response DataExchangeCredsResp
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
