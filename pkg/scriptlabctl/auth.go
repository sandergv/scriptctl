package scriptlabctl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

func (c *Client) Login(host string, username string, password string) (types.AuthDetails, error) {

	url := host + "/v1/auth/login"

	request := types.LoginRequest{
		Username: username,
		Password: password,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return types.AuthDetails{}, err
	}

	//
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return types.AuthDetails{}, err
	}

	c.setHeaders(req)
	resp, err := c.http.Do(req)
	if err != nil {
		return types.AuthDetails{}, err
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		br, err := io.ReadAll(resp.Body)
		if err != nil {
			return types.AuthDetails{}, err
		}
		return types.AuthDetails{}, errors.New(strings.ReplaceAll(string(br), "\n", ""))
	}

	response := types.LoginResponse{}
	json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return types.AuthDetails{}, err
	}

	return types.AuthDetails{
		WorkspaceID: response.WorkspaceID,
		Token:       response.Token,
		ExpiresAt:   response.ExpiresAt,
	}, nil
}
