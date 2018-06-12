package actors

// IDGeneratorCreator - функция создания генератовой идентификаторов
type IDGeneratorCreator func(parentID ActorID) IDGenerator

// Listener - интерфейс обратных вызовов соответствующих событий
type Listener interface {
	onMessage(address ActorID, message Message)
	onSpawn(id ActorID, actor Actor)
}

// Actors - моедль акторов
type Actors struct {
	rootID      ActorID
	router      messageRouter
	rootSpawner spawner
	listener    Listener
}

type props struct {
	idGeneratorCreator IDGeneratorCreator
	rootID             string
	rootIDGenerator    IDGenerator
}

func defaultIDGeneratorCreator(parentID ActorID) IDGenerator {
	return &NumericIDGenerator{}
}

// New создаёт новый экземпляр Actors
func New(props props) Actors {
	rootID := props.rootID
	var idGeneratorCreator IDGeneratorCreator
	if props.idGeneratorCreator != nil {
		idGeneratorCreator = props.idGeneratorCreator
	} else {
		idGeneratorCreator = defaultIDGeneratorCreator
	}
	var rootIDGenerator IDGenerator
	if props.rootIDGenerator != nil {
		rootIDGenerator = props.rootIDGenerator
	} else {
		rootIDGenerator = idGeneratorCreator(rootID)
	}
	router := newMessageRouter()
	rootSpawner := spawner{
		idGenerator:        rootIDGenerator,
		idGeneratorCreator: idGeneratorCreator,
		registrator:        &router,
	}
	actors := Actors{rootID: rootID, router: router, rootSpawner: rootSpawner}
	rootSpawner.destroyer = &actors
	rootSpawner.sender = &actors
	return actors
}

// Send отправляет сообщение соответствующему актору
func (a *Actors) Send(address ActorID, message Message) {
	if a.router.route(address, message) {
		if a.listener != nil {
			a.listener.onMessage(address, message)
		}
	}
}

// SetListener устанавливает листенер событий
func (a *Actors) SetListener(listener Listener) {
	a.listener = listener
}

// Spawn пораджает новый актор
func (a *Actors) Spawn(actor Actor) ActorID {
	id := a.rootSpawner.spawn(a.rootID, actor)
	a.onSpawn(id, actor)
	return id
}

func (a *Actors) destroy(id ActorID) {
	a.router.unregister(id)
}

func (a *Actors) onSpawn(id ActorID, actor Actor) {
	if a.listener != nil {
		a.listener.onSpawn(id, actor)
	}
}

func (a *Actors) send(address ActorID, message Message) {
	a.Send(address, message)
}
