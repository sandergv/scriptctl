package scriptlabctl

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sandergv/scriptctl/pkg/scriptlabctl/types"
)

func (c *Client) CreateNamespace(opts types.CreateNamespaceOptions) (string, error) {

	url := c.url + "/v1/namespace"

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

	response := types.CreateNamespaceResponse{}
	json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Status != "success" {
		return "", errors.New(response.Error)
	}
	return response.ID, nil
}
