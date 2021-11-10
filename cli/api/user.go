package api

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Dataset is a set of rules that define attributes over a distinct set of members
type User struct {
	ID               string `mapstructure:"id" json:"id"`
	UserName         string `mapstructure:"user_name" json:"user_name"`
	OrganizationName string `mapstructure:"organization_name" json:"organization_name"`
	OrganizationID   string `mapstructure:"organization_id" json:"organization_id"`
}

type AuthResponse struct {
	User User   `mapstructure:"user" json:"user"`
	JWT  string `mapstructure:"jwt" json:"jwt"`
}

type AuthPayload struct {
	UserName string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
}

// GetDatasets returns the partitions belonging to a member
func (c Client) Authenticate(username string, password string) (*AuthResponse, error) {
	body, err := json.Marshal(AuthPayload{UserName: username, Password: password})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshaling json for condition %v", username))
	}
	res, err := c.Post("/users/authenticate", body)

	var response AuthResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c Client) GetMe() (*User, error) {
	res, err := c.Get("/users/me")

	var response User
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
