package messenger

import (
	"context"
	"io"
	"remind-of-rain/src/httpClient"
	"testing"
)

func TestPushover_SendMessage(t *testing.T) {
	type fields struct {
		AppToken  string
		UserToken string
		httpDo    httpClient.HttpClient
	}
	type args struct {
		title   string
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "correct send a message",
			fields: fields{
				AppToken:  "correct-app-token",
				UserToken: "correct-user-token",
				httpDo: httpGetterMock{
					t: t,
					DoMock: func(ctx context.Context, url, method string, body io.Reader) ([]byte, error) {
						return []byte(`
							{
							    "status": 1,
							    "request": "398e18d0-f6ae-4e08-a3cd-44a88eb8b9a9"
							}
						`), nil
					},
				},
			},
			args: args{
				title:   "the Title",
				message: "The message",
			},
			wantErr: false,
		},
		{
			name: "try to send with incorrect application-token",
			fields: fields{
				AppToken:  "incorrect-application-token",
				UserToken: "correct-user-token",
				httpDo: httpGetterMock{
					t: t,
					DoMock: func(ctx context.Context, url, method string, body io.Reader) ([]byte, error) {
						return []byte(`
							{
							    "token": "invalid",
							    "errors": [
							        "application token is invalid"
							    ],
							    "status": 0,
							    "request": "4177ba90-77f6-4711-b37f-0a6e6c36ec10"
							}
						`), nil
					},
				},
			},
			args: args{
				title:   "the Title",
				message: "the Message",
			},
			wantErr: true,
		},
		{
			name: "try to send with incorrect user-token",
			fields: fields{
				AppToken:  "correct-app-token",
				UserToken: "incorrect-user-token",
				httpDo: httpGetterMock{
					t: t,
					DoMock: func(ctx context.Context, url, method string, body io.Reader) ([]byte, error) {
						return []byte(`
							{
							    "user": "invalid",
							    "errors": [
							        "user identifier is not a valid user, group, or subscribed user key"
							    ],
							    "status": 0,
							    "request": "00a99b89-91f4-4fe9-bde4-480bf397ef44"
							}
						`), nil
					},
				},
			},
			args: args{
				title:   "the Title",
				message: "the Message",
			},
			wantErr: true,
		},
		{
			name: "try to send with empty message",
			fields: fields{
				AppToken:  "correct-app-token",
				UserToken: "correct-user-token",
				httpDo: httpGetterMock{
					t: t,
					DoMock: func(ctx context.Context, url, method string, body io.Reader) ([]byte, error) {
						return []byte(`
							{
							    "message": "cannot be blank",
							    "errors": [
							        "message cannot be blank"
							    ],
							    "status": 0,
							    "request": "751f9fe6-84b3-42b5-b1cf-02b1c4d5546d"
							}
						`), nil
					},
				},
			},
			args: args{
				title:   "the Title",
				message: "",
			},
			wantErr: true,
		},
		{
			name: "try to send with empty message, bad tokens",
			fields: fields{
				AppToken:  "incorrect-app-token",
				UserToken: "incorrect-user-token",
				httpDo: httpGetterMock{
					t: t,
					DoMock: func(ctx context.Context, url, method string, body io.Reader) ([]byte, error) {
						return []byte(`
							{
							    "message": "cannot be blank",
								"token": "invalid",
								"user": "invalid",
							    "errors": [
							        "message cannot be blank",
							        "application token is invalid",
							        "user identifier is not a valid user, group, or subscribed user key"
							    ],
							    "status": 0,
							    "request": "751f9fe6-84b3-42b5-b1cf-02b1c4d5546d"
							}
						`), nil
					},
				},
			},
			args: args{
				title:   "the Title",
				message: "",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Pushover{
				AppToken:  tt.fields.AppToken,
				UserToken: tt.fields.UserToken,
				httpDo:    tt.fields.httpDo,
			}
			if err := p.SendMessage(ctx, tt.args.title, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type httpGetterMock struct {
	t      *testing.T
	DoMock func(ctx context.Context, url, method string, body io.Reader) ([]byte, error)
}

func (h httpGetterMock) Do(ctx context.Context, url, method string, body io.Reader) ([]byte, error) {
	if url == "" || method == "" {
		h.t.Error("Do(): url and method is required")
	}
	return h.DoMock(ctx, url, method, body)
}
