package actors

import (
	"strconv"
	"strings"
)

const separator = "/"

// NumericIDGenerator - числовой генератор идентификаторов
type NumericIDGenerator struct {
	prefix string
	index  int
}

// NewID - создаёт новый идентификатор актора
func (g *NumericIDGenerator) NewID() ActorID {
	actorID := strings.Join([]string{g.prefix, strconv.Itoa(g.index)}, separator)
	g.index++
	return actorID
}
