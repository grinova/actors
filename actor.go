package actors

// ActorID - идентификатор актора
type ActorID = string

// Message - сообщение
type Message = interface{}

// Spawn - функция создания актора
type Spawn = func(creator ActorCreator) (ActorID, bool)

// Send - функция отправки сообщения
type Send = func(id ActorID, message Message)

// Exit - функция завершения актора и отправка последнего сообщения
type Exit = func(message Message)

// Actor - актор
type Actor struct {
	OnInit    func(selfID ActorID, send Send, spawn Spawn, exit Exit)
	OnMessage func(message Message, send Send, spawn Spawn, exit Exit)
}
