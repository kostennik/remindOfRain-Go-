package httpClient

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient interface {
	Do(url, method string, body io.Reader) ([]byte, error)
}

type httpClient struct {
	timeout time.Duration
}

func NewHttpClient(timeout time.Duration) *httpClient {
	return &httpClient{timeout: timeout}
}

func (h httpClient) Do(url, method string, body io.Reader) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, errors.Wrapf(err, "error while creating a request to %s", url)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "error while sending a request to %s", url)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error().Interface("headers", resp.Header).Int("response code", resp.StatusCode).Str("response msg", resp.Status).Send()
	}

	responseBodyRaw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error occurred while reading response body")
	}

	return responseBodyRaw, nil
}
