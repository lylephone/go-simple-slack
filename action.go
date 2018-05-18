package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func (s *Client) Post(ctx context.Context, p *Payload) (*http.Response, error) {
	j, _ := json.MarshalIndent(p, "", "\t")
	log.Printf("%+v", string(j))

	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create slack message")
	}

	req, err := http.NewRequest(http.MethodPost, s.WebHookURL.String(), bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	req.Header.Set("Content-type", "application/json")
	req = req.WithContext(ctx)
	return s.httpClient.Do(req)

}
