package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func (s *Client) Post(ctx context.Context, p *Payload) (*http.Response, error) {
	jsonBytes, err := json.Marshal(p)

	if err != nil {
		message := "failed to create slack message"
		s.log.Println(message)
		s.log.Println(err.Error())
		return nil, errors.Wrap(err, message)
	}
	s.log.Printf("%+v", string(jsonBytes))

	req, err := http.NewRequest(http.MethodPost, s.WebHookURL.String(), bytes.NewReader(jsonBytes))
	if err != nil {
		message := "failed to create request"
		s.log.Println(message)
		s.log.Println(err.Error())
		return nil, errors.Wrap(err, message)
	}

	req.Header.Set("Content-type", "application/json")
	req = req.WithContext(ctx)
	return s.httpClient.Do(req)

}
