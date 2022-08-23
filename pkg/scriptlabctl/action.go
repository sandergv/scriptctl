package scriptlabctl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

func (c *Client) CreateAction(opts types.CreateActionRequest) (string, error) {

	url := c.url + "/v1/action"

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

	response := types.CreateActionResponse{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", err
	}
	fmt.Println(response)
	if response.Status == "error" {
		return "", errors.New(response.Error)
	}
	return response.ID, nil
}

func (c *Client) GetActionList() ([]types.Action, error) {

	url := c.url + "/v1/action"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []types.Action{}, err
	}

	// add headers values
	c.setHeaders(req)

	//
	res, err := c.http.Do(req)

	response := types.GetActionListResponse{}
	json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return []types.Action{}, err
	}

	if response.Status == "error" {
		return []types.Action{}, errors.New(response.Error)
	}
	return response.Data, nil

}
