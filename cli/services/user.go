package services

import (
	"flip/api"
	"fmt"

	"github.com/spf13/viper"
)

type User struct {
	Jwt    string
	User   *api.User
	DataX  *api.DataExchangeCredsResp
	APIUrl string
}

func NewUser() (*User, error) {
	jwt := viper.Get("jwt")
	if jwt == nil || jwt == "" {
		return nil, fmt.Errorf("⛔️ No user is currently logged-in. Run `flip auth login` to login to your account")
	}
	return &User{Jwt: jwt.(string), APIUrl: "https://fsc-flip-api.herokuapp.com"}, nil
}

func (u *User) GetMe() (*api.User, error) {
	config := api.Config{BaseURL: u.APIUrl, JWT: u.Jwt}
	client, _ := api.NewClient(config)
	me, err := client.GetMe()
	if err != nil {
		return nil, err
	}
	u.User = me
	return me, nil
}

func (u *User) GetDatax() (*api.DataExchangeCredsResp, error) {
	config := api.Config{BaseURL: u.APIUrl, JWT: u.Jwt}
	client, _ := api.NewClient(config)
	creds, err2 := client.GetDataExchangeCreds()
	if err2 != nil {
		return nil, fmt.Errorf("⛔️ An error occurred fetching your credentials. Please try again")
	}
	u.DataX = creds
	return creds, nil
}
