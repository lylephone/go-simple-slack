package slack

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func NewClient(u string, opts ...Option) (*Client, error) {
	webHookUrl, err := url.Parse(u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse url string")

	}
	c := &Client{
		WebHookURL: webHookUrl,
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, errors.Wrap(err, "invalid option")
		}
	}

	return c, nil
}

type Client struct {
	WebHookURL *url.URL
	httpClient *http.Client
}

type Option func(*Client) error

func WithHTTPClient(c *http.Client) Option {
	return func(sc *Client) error {
		sc.httpClient = c
		return nil
	}
}
