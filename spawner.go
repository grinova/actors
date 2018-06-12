package actors

type spawnListener interface {
	onSpawn(id ActorID, actor Actor)
}

type spawner struct {
	idGenerator        IDGenerator
	idGeneratorCreator IDGeneratorCreator
	registrator        actorRegistrator
	destroyer          destroyer
	sender             sender
	listener           spawnListener
}

func (s *spawner) spawn(parentID ActorID, actor Actor) ActorID {
	id := s.idGenerator.NewID()
	spawner := spawner{
		idGenerator:        s.idGeneratorCreator(id),
		idGeneratorCreator: s.idGeneratorCreator,
		registrator:        s.registrator,
		destroyer:          s.destroyer,
		sender:             s.sender,
		listener:           s.listener,
	}
	owner := actorOwner{
		parentID:  parentID,
		id:        id,
		actor:     actor,
		spawner:   spawner,
		sender:    s.sender,
		destroyer: s.destroyer,
	}
	s.registrator.register(id, &owner)
	owner.init()
	s.listener.onSpawn(id, actor)
	return id
}
