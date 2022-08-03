package scriptlabctl

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sandergv/scriptctl/pkg/scriptlabctl/types"
)

func (c *Client) RunCode(opts types.RunCodeOptions) (types.RunDetails, error) {

	request := types.RunCodeRequest{
		ExecEnv: opts.ExecEnv,
		Type:    opts.Type,
		Envs:    opts.Env,
		Args:    opts.Args,
		Content: opts.Code,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return types.RunDetails{}, err
	}

	req, err := http.NewRequest(http.MethodPost, c.url+"/run", bytes.NewBuffer(body))
	if err != nil {
		return types.RunDetails{},
			err
	}

	// add headers values
	c.setHeaders(req)

	//
	res, err := c.http.Do(req)

	// parsing response
	response := types.RunResponse{}
	json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return types.RunDetails{}, err
	}

	if response.Status != "success" {
		return types.RunDetails{}, errors.New(response.Error)
	}
	return response.Details, nil
}

func (c *Client) RunExec(opts types.RunExecOptions) (types.RunDetails, error) {

	request := types.RunExecRequest{
		Body: opts.Body,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return types.RunDetails{}, err
	}

	req, err := http.NewRequest(http.MethodPost, c.url+"/run/"+opts.ExecID, bytes.NewBuffer(body))
	if err != nil {
		return types.RunDetails{},
			err
	}

	// add headers values
	c.setHeaders(req)

	//
	res, err := c.http.Do(req)

	// parsing response
	response := types.RunResponse{}
	json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return types.RunDetails{}, err
	}

	if response.Status != "success" {
		return types.RunDetails{}, errors.New(response.Error)
	}
	return response.Details, nil
}
