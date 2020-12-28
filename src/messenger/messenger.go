package messenger

type Messenger interface {
	SendMessage(title string, message string) error
}
