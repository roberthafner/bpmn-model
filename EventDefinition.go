package model

type EventDefinition interface {
	BaseElement

	// TODO: Methods
	// Clone() EventDefinition
}

type eventDefinition struct {
	baseElement
}
