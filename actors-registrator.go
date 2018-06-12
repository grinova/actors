package actors

type actorRegistrator interface {
	register(id ActorID, handler messageHandler) bool
}
