package scriptlabctl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

func (c *Client) Login(username string, password string) (string, time.Time, error) {

	url := c.url + "/v1/auth/login"

	request := types.LoginRequest{
		Username: username,
		Password: password,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return "", time.Time{}, err
	}

	//
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", time.Time{}, err
	}

	c.setHeaders(req)
	resp, err := c.http.Do(req)
	if err != nil {
		return "", time.Time{}, err
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		br, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", time.Time{}, err
		}
		return "", time.Time{}, errors.New(strings.ReplaceAll(string(br), "\n", ""))
	}

	response := types.LoginResponse{}
	json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", time.Time{}, err
	}

	return response.Token, response.ExpiresAt, nil
}
