package actors

type actorOwner struct {
	parentID  ActorID
	id        ActorID
	actor     Actor
	spawner   spawner
	sender    sender
	destroyer destroyer
}

func (ao *actorOwner) init() {
	if ao.actor != nil {
		ao.actor.OnInit(ao.id, ao.sender.send, ao.spawn, ao.exit)
	}
}

func (ao *actorOwner) handle(message Message) {
	if ao.actor != nil {
		ao.actor.OnMessage(message, ao.sender.send, ao.spawn, ao.exit)
	}
}

func (ao *actorOwner) spawn(creator ActorCreator) (ActorID, bool) {
	return ao.spawner.spawn(ao.id, creator)
}

func (ao *actorOwner) exit(message Message) {
	ao.destroyer.destroy(ao.id)
	if message != nil {
		ao.sender.send(ao.parentID, message)
	}
}
