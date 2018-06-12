package actors

type messageHandler interface {
	handle(message Message)
}
