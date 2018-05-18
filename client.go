package slack

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/pkg/errors"
)

func NewClient(u string, opts ...Option) (*Client, error) {
	webHookUrl, err := url.Parse(u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse url string")

	}

	// http client setting
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			ServerName: webHookUrl.Hostname(),
		},
	}
	hc := &http.Client{
		Transport: tr,
	}
	c := &Client{
		WebHookURL: webHookUrl,
		httpClient: hc,
	}

	// default logger setting
	c.log = log.New(ioutil.Discard, "", 0)

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
	log        *log.Logger
}

type Option func(*Client) error

func WithHTTPClient(c *http.Client) Option {
	return func(sc *Client) error {
		sc.httpClient = c
		return nil
	}
}

func WithDebug() Option {
	return func(sc *Client) error {
		sc.log = log.New(os.Stdout, "go-slack-client: ", log.LstdFlags)
		return nil
	}
}
