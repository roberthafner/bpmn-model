package model

type Event interface {
	FlowNode
	EventDefinitions() []EventDefinition

	// TODO: Methods
	// AddEventDefinition(ed EventDefinition)
	// SetEventDefinitions(ed []EventDefinition)
}

type event struct {
	flowNode
	eventDefinitions []EventDefinition
}

func (e event) EventDefinitions() []EventDefinition {
	return e.eventDefinitions
}
