package scriptlabctl

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

func (c *Client) CreateScript(opts types.CreateScriptOptions) (string, error) {

	url := c.url + "/v1/script"

	if len(opts.Name) > 24 {
		return "", errors.New("script name is too long")
	}

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

	response := types.CreateScriptResponse{}
	json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Status != "success" {
		return "", errors.New(response.Error)
	}
	return response.ID, nil
}

func (c *Client) UpdateScript(opts types.UpdateScriptFileRequest) (time.Time, error) {
	url := c.url + "/v1/script/" + opts.ID

	body, err := json.Marshal(opts)
	if err != nil {
		return time.Time{}, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return time.Time{}, err
	}

	// add headers values
	c.setHeaders(req)

	//
	res, err := c.http.Do(req)

	response := types.UpdateScriptFileResponse{}
	json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return time.Time{}, err
	}

	if response.Status != "success" {
		return time.Time{}, errors.New(response.Error)
	}
	return time.Now(), nil
}

func (c *Client) GetScriptList() ([]types.Script, error) {

	url := c.url + "/v1/script"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []types.Script{}, err
	}

	// add headers values
	c.setHeaders(req)

	//
	res, err := c.http.Do(req)

	response := types.GetScriptListResponse{}
	json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return []types.Script{}, err
	}

	if response.Status != "success" {
		return []types.Script{}, errors.New(response.Error)
	}
	return response.Data, nil
}
