package model

import "strings"

type FlowElement interface {
	BaseElement
	Name() string
	Documentation() string
	// TODO: Methods
	// ExecutionListeners() []
	// ParentContainer()

	String() string
}

type flowElement struct {
	baseElement
	name          string
	documentation string
	// TODO: Fields?
	// executionListeners []
	// parentContainer []
}

func (fe flowElement) Name() string {
	return fe.name
}

func (fe flowElement) Documentation() string {
	return fe.documentation
}

func (fe flowElement) String() string {
	var s strings.Builder
	s.WriteString("FlowElement: (id=")
	s.WriteString(fe.Id())
	s.WriteString(", name=")
	s.WriteString(fe.Name())
	s.WriteString(")\n")

	return s.String()
}
