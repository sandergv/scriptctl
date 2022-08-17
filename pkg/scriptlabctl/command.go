package scriptlabctl

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

func (c *Client) CreateCommand(opts types.CreateCommandRequest) (string, error) {

	url := c.url + "/v1/command"

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
	if err != nil {
		return "", err
	}

	response := types.CreateCommandResponse{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Status == "error" {
		return "", errors.New(response.Error)
	}
	return response.ID, nil
}

func (c *Client) GetCommandList() ([]types.Command, error) {

	url := c.url + "/v1/command"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []types.Command{}, err
	}

	// add headers values
	c.setHeaders(req)

	//
	res, err := c.http.Do(req)

	response := types.GetCommandListResponse{}
	json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return []types.Command{}, err
	}

	if response.Status == "error" {
		return []types.Command{}, errors.New(response.Error)
	}
	return response.Data, nil

}

func (c *Client) RunCommand(command string, args []string) (types.RunDetails, error) {

	url := c.url + "/v1/command/run"

	body, err := json.Marshal(types.RunCommandRequest{
		Name: command,
		Args: args,
	})
	if err != nil {
		return types.RunDetails{}, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return types.RunDetails{}, err
	}

	// add headers values
	c.setHeaders(req)

	//
	res, err := c.http.Do(req)
	if err != nil {
		return types.RunDetails{}, err
	}

	response := types.RunCommandResponse{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return types.RunDetails{}, err
	}

	if response.Status == "error" {
		return types.RunDetails{}, errors.New(response.Error)
	}
	return response.Details, nil
}
