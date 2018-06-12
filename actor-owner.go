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
	ao.actor.OnInit(ao.id, ao.sender.send, ao.spawn, ao.exit)
}

func (ao *actorOwner) handle(message Message) {
	ao.actor.OnMessage(message, ao.sender.send, ao.spawn, ao.exit)
}

func (ao *actorOwner) spawn(actor Actor) ActorID {
	return ao.spawner.spawn(ao.id, actor)
}

func (ao *actorOwner) exit(message Message) {
	ao.destroyer.destroy(ao.id)
	if message != nil {
		ao.sender.send(ao.parentID, message)
	}
}
