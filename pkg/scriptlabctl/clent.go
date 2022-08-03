package scriptlabctl

import "net/http"

type Client struct {
	url     string
	http    *http.Client
	headers map[string]string
}

func NewClient(token string) *Client {

	return &Client{
		url:  "http://localhost:6892",
		http: &http.Client{},
		headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + token,
		},
	}
}

func (c *Client) setHeaders(req *http.Request) {
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
}
