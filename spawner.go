package actors

// ActorCreator - функция создания актора
type ActorCreator func(id ActorID) (Actor, bool)

type spawner struct {
	idGenerator        IDGenerator
	idGeneratorCreator IDGeneratorCreator
	registrator        actorRegistrator
	destroyer          destroyer
	sender             sender
	onSpawn            func(id ActorID, actor Actor)
}

func (s *spawner) spawn(parentID ActorID, creator ActorCreator) (ActorID, bool) {
	id := s.idGenerator.NewID()
	actor, ok := creator(id)
	if !ok {
		return "", false
	}
	spawner := spawner{
		idGenerator:        s.idGeneratorCreator(id),
		idGeneratorCreator: s.idGeneratorCreator,
		registrator:        s.registrator,
		destroyer:          s.destroyer,
		sender:             s.sender,
		onSpawn:            s.onSpawn,
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
	s.onSpawn(id, actor)
	return id, true
}
