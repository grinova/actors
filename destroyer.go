package actors

type destroyer interface {
	destroy(id ActorID)
}
