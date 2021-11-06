package messenger

import "context"

type Messenger interface {
	SendMessage(ctx context.Context, title string, message string) error
}
