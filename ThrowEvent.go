package model

type ThrowEvent interface {
	Event

	// TODO: Methods
	// Clone() ThrowEvent
	// SetValues(te ThrowEvent)
}

type throwEvent struct {
	event
}
