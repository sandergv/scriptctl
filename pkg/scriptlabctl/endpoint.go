package scriptlabctl

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

func (c *Client) CreateEndpoint(opts types.CreateEndpointOptions) (string, error) {

	url := c.url + "/v1/endpoint"

	body, err := json.Marshal(opts)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	// add headers values
	c.setHeaders(req)

	//
	res, err := c.http.Do(req)

	response := types.CreateEndpointResponse{}
	json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Status != "success" {
		return "", errors.New(response.Error)
	}
	return response.ID, nil
}

func (c *Client) GetEndpointList(namespace string) ([]types.Endpoint, error) {

	url := c.url + "/v1/endpoint"

	if namespace != "" {
		url += "?namespace=" + namespace
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []types.Endpoint{}, err
	}

	// add headers values
	c.setHeaders(req)

	//
	res, err := c.http.Do(req)

	response := types.GetEndpointListResponse{}
	json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return []types.Endpoint{}, err
	}

	if response.Status != "success" {
		// return "", errors.New(response.Error)
	}
	return response.Data, nil

}
