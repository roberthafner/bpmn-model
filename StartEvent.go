package model

type StartEvent struct {
	catchEvent
	Initiator string
	//FormKey string
	Interrupting bool
	//FormProperties []

	// TODO: Fields
	// TODO: Methods
}

func NewStartEvent(id string, name string, documentation string) StartEvent {
	se := StartEvent{}
	se.id = id
	se.name = name
	se.documentation = documentation
	return se
}
