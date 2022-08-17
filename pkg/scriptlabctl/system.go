package scriptlabctl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

func (c *Client) CheckStatus() (string, error) {

	url := c.url + "/v1/system/status"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	// add headers values
	c.setHeaders(req)

	//
	resp, err := c.http.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("unreachable")
	}

	if resp.StatusCode == 401 {
		return "", errors.New("unauthorized")
	}

	if resp.StatusCode != 200 {
		return "", errors.New("unreachable")
	}

	response := types.CheckStatusResponse{}

	json.NewDecoder(resp.Body).Decode(&response)

	return response.Status, nil
}
