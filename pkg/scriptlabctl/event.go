package scriptlabctl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

func (c *Client) SendEvent(channel string, message string) error {

	url := c.url + "/v1/events/" + channel + "/send"
	fmt.Println(url)
	request := types.SendEventMessageRequest{
		Message: message,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	// add headers values
	c.setHeaders(req)

	//
	res, err := c.http.Do(req)
	if err != nil {
		return err
	}

	fmt.Println(res.StatusCode)
	return nil
}
