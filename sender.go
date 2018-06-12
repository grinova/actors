package actors

type sender interface {
	send(id ActorID, message Message)
}
