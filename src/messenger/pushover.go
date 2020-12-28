package messenger

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

const (
	pushoverUrl         = "https://api.pushover.net/1/messages.json"
	pushoverContentType = "application/x-www-form-urlencoded"
)

type Pushover struct {
	AppToken  string
	UserToken string
	httpPost  func(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

func NewPushover(appToken string, userToken string) *Pushover {
	return &Pushover{
		AppToken:  appToken,
		UserToken: userToken,
		httpPost:  http.Post,
	}
}

func (p Pushover) SendMessage(title string, message string) error {
	requestBody := []byte(
		"token=" + p.AppToken +
			"&user=" + p.UserToken +
			"&device=device" +
			"&title=" + title +
			"&message=" + message)

	resp, err := p.httpPost(pushoverUrl, pushoverContentType, bytes.NewBuffer(requestBody))
	if err != nil {
		return errors.Wrapf(err, "error while sending a request to %s", pushoverUrl)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error().Interface("headers", resp.Header).Int("response code", resp.StatusCode).Str("response msg", resp.Status).Send()
	}

	return nil
}
