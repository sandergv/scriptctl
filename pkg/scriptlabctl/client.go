package scriptlabctl

import "net/http"

type ClientOptions struct {
	Url   string
	Token string
}

type Client struct {
	url     string
	token   string
	http    *http.Client
	headers map[string]string
}

func NewClient(opts ClientOptions) *Client {

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	if opts.Token != "" {
		headers["Authorization"] = "Bearer " + opts.Token
	}

	return &Client{
		url:     opts.Url,
		http:    &http.Client{},
		token:   opts.Token,
		headers: headers,
	}
}

func (c *Client) setHeaders(req *http.Request) {
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
}
