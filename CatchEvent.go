package model

type CatchEvent interface {
	Event

	// TODO: Methods
	// Clone() ThrowEvent
	// SetValues(te ThrowEvent)
}

type catchEvent struct {
	event
}
