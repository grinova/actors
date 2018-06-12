package actors

type messageRouter struct {
	handlers map[ActorID]messageHandler
}

func newMessageRouter() messageRouter {
	return messageRouter{handlers: make(map[ActorID]messageHandler)}
}

func (r *messageRouter) register(id ActorID, handler messageHandler) bool {
	_, ok := r.handlers[id]
	r.handlers[id] = handler
	return !ok
}

func (r *messageRouter) unregister(id ActorID) bool {
	_, ok := r.handlers[id]
	delete(r.handlers, id)
	return ok
}

func (r *messageRouter) route(id ActorID, message Message) bool {
	handler, ok := r.handlers[id]
	if ok {
		handler.handle(message)
	}
	return ok
}
