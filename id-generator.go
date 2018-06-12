package actors

// IDGenerator - генератор идентификаторов акторов
type IDGenerator interface {
	NewID() ActorID
}
