package producer

import (

)

type EventType int

const (
	EventTypeUnknown EventType = iota
	EventTypeOne
	EventTypeSentinal
)

func (e EventType) ReflexType() int {
	return int(e)
}
