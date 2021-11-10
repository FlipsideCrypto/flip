package api

import "fmt"

// Config allows a consuming app to set up API Key
type Config struct {
	BaseURL string
	JWT     string
}

// Client allows access to the Flip RPC Interface
type Client struct {
	ApiURL string
	JWT    string
}

// NewClient returns a new Databridge Client
func NewClient(config Config) (Client, error) {
	c := Client{}
	c.JWT = config.JWT
	c.ApiURL = fmt.Sprintf("%s", config.BaseURL)
	return c, nil
}
