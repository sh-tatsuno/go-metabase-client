package metabase

import (
	"io"
	"net/http"
	"net/url"
	"path"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

const (
	metabaseTokenHeader = "X-Metabase-Session"
)

type Client struct {
	URL url.URL

	Token *string

	HTTPClient *http.Client
}

func NewClient(baseURL string, user string, password string) (*Client, error) {

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	client := &Client{
		URL:        *u,
		HTTPClient: cleanhttp.DefaultClient(),
	}

	_, err = client.CreateSession(user, password)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) newRequest(method, requestPath string, query url.Values, body io.Reader) (*http.Request, error) {

	url := c.URL
	url.Path = path.Join(url.Path, requestPath)
	url.RawQuery = query.Encode()

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return req, err
	}
	if c.Token != nil {
		req.Header.Add(metabaseTokenHeader, *(c.Token))
	}

	req.Header.Add("Content-Type", "application/json")
	return req, err
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	return resp, err
}
