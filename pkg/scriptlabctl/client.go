package scriptlabctl

import "net/http"

type Client struct {
	url     string
	token   string
	http    *http.Client
	headers map[string]string
}

func NewClient(token string) *Client {

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}

	return &Client{
		url:     "http://localhost:6892",
		http:    &http.Client{},
		token:   token,
		headers: headers,
	}
}

func (c *Client) setHeaders(req *http.Request) {
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
}
