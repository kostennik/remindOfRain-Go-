package messenger

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
	client "remind-of-rain/src/httpClient"
	"strings"
	"time"
)

const (
	pushoverUrl = "https://api.pushover.net/1/messages.json"
)

type Pushover struct {
	AppToken  string
	UserToken string
	httpDo    client.HttpClient
}

func NewPushover(appToken string, userToken string) *Pushover {
	return &Pushover{
		AppToken:  appToken,
		UserToken: userToken,
		httpDo:    client.NewHttpClient(20 * time.Second),
	}
}

func (p Pushover) SendMessage(ctx context.Context, title string, message string) error {
	log.Debug().Msg("starting SendMessage()")

	if message == "" {
		return errEmptyMessage
	}

	requestBody := []byte(
		"token=" + p.AppToken +
			"&user=" + p.UserToken +
			"&device=device" +
			"&title=" + title +
			"&message=" + message)

	resp, err := p.httpDo.Do(ctx, pushoverUrl, http.MethodPost, bytes.NewBuffer(requestBody))
	if err != nil {
		return errors.Wrapf(err, "error while sending a request to %s", pushoverUrl)
	}

	err = parseErrorResponse(resp)
	if err != nil {
		return err
	}

	log.Debug().Msg("end SendMessage()")
	return nil
}

func parseErrorResponse(resp []byte) error {
	var respBody = new(responseBody)
	if err := json.Unmarshal(resp, respBody); err != nil {
		return errors.Wrapf(err, "SendMessage(): error while encoding struct from json")
	}

	if respBody != nil && respBody.Status != 1 && respBody.Errors != nil {
		var errMsg strings.Builder
		for _, msg := range respBody.Errors {
			errMsg.WriteString(msg)
		}
		return errors.New(errMsg.String())
	}
	return nil
}

type responseBody struct {
	Status int      `json:"status"`
	Errors []string `json:"errors"`
}
