package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// CallRPC returns a response from the RPC interface
func (c Client) Post(path string, json_body []byte) ([]byte, error) {
	req, _ := http.NewRequest("POST", c.ApiURL+path, bytes.NewBuffer(json_body))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("flip api responded with a non-200 for %s, err: %s", c.ApiURL, err))
	}

	responseBody, errBodyRead := ioutil.ReadAll(res.Body)
	if errBodyRead != nil {
		return nil, errBodyRead
	}

	return []byte(responseBody), nil
}

func (c Client) Get(path string) ([]byte, error) {
	req, _ := http.NewRequest("GET", c.ApiURL+path, nil)
	req.Header.Set("authorization", c.JWT)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("flip api responded with a non-200 for %s, err: %s", c.ApiURL, err))
	}

	responseBody, errBodyRead := ioutil.ReadAll(res.Body)
	if errBodyRead != nil {
		return nil, errBodyRead
	}

	return []byte(responseBody), nil
}
