package actors

// IDGeneratorCreator - функция создания генератовой идентификаторов
type IDGeneratorCreator func(parentID ActorID) IDGenerator

// Listener - интерфейс обратных вызовов соответствующих событий
type Listener interface {
	OnMessage(address ActorID, message Message)
	OnSpawn(id ActorID, actor Actor)
}

// Actors - моедль акторов
type Actors struct {
	rootID      ActorID
	router      messageRouter
	rootSpawner spawner
	listener    Listener
}

// Props - свойства модели акторов
type Props struct {
	IDGeneratorCreator IDGeneratorCreator
	RootID             string
	RootIDGenerator    IDGenerator
}

func defaultIDGeneratorCreator(parentID ActorID) IDGenerator {
	return &NumericIDGenerator{prefix: parentID}
}

// New создаёт новый экземпляр Actors
func New(props Props) Actors {
	rootID := props.RootID
	var idGeneratorCreator IDGeneratorCreator
	if props.IDGeneratorCreator != nil {
		idGeneratorCreator = props.IDGeneratorCreator
	} else {
		idGeneratorCreator = defaultIDGeneratorCreator
	}
	var rootIDGenerator IDGenerator
	if props.RootIDGenerator != nil {
		rootIDGenerator = props.RootIDGenerator
	} else {
		rootIDGenerator = idGeneratorCreator(rootID)
	}
	router := newMessageRouter()
	actors := Actors{rootID: rootID, router: router}
	rootSpawner := spawner{
		idGenerator:        rootIDGenerator,
		idGeneratorCreator: idGeneratorCreator,
		registrator:        &router,
		sender:             &actors,
		destroyer:          &actors,
		onSpawn:            actors.onSpawn,
	}
	actors.rootSpawner = rootSpawner
	rootSpawner.destroyer = &actors
	rootSpawner.sender = &actors
	return actors
}

// Send отправляет сообщение соответствующему актору
func (a *Actors) Send(address ActorID, message Message) {
	if a.router.route(address, message) {
		if a.listener != nil {
			a.listener.OnMessage(address, message)
		}
	}
}

// SetListener устанавливает листенер событий
func (a *Actors) SetListener(listener Listener) {
	a.listener = listener
}

// Spawn пораджает новый актор
func (a *Actors) Spawn(creator ActorCreator) (ActorID, bool) {
	return a.rootSpawner.spawn(a.rootID, creator)
}

func (a *Actors) destroy(id ActorID) {
	a.router.unregister(id)
}

func (a *Actors) onSpawn(id ActorID, actor Actor) {
	if a.listener != nil {
		a.listener.OnSpawn(id, actor)
	}
}

func (a *Actors) send(address ActorID, message Message) {
	a.Send(address, message)
}
