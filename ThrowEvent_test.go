package model

import "testing"

func TestThrowEvent_Methods(t *testing.T) {

	te := throwEvent{
		event: event{
			flowNode: flowNode{
				flowElement: flowElement{
					baseElement: baseElement{
						id:                  "123",
						xmlRowNumber:        30,
						xmlColumnNumber:     20,
						extensionElements:   map[string][]ExtensionElement{"ee": []ExtensionElement{ExtensionElement{}}},
						extensionAttributes: map[string][]ExtensionAttribute{"ea": []ExtensionAttribute{ExtensionAttribute{}}},
					},
					name:          "name",
					documentation: "documentation",
				},
				asynchronous: true,
				notExclusive: true,
				incoming:     []SequenceFlow{SequenceFlow{}},
				outgoing:     []SequenceFlow{SequenceFlow{}},
			},
			eventDefinitions: []EventDefinition{eventDefinition{}},
		},}

	if "123" != te.Id() {
		t.Fail()
	}

	if 30 != te.xmlRowNumber {
		t.Fail()
	}

	if 20 != te.xmlColumnNumber {
		t.Fail()
	}

	if 1 != len(te.extensionElements) {
		t.Fail()
	}

	if 1 != len(te.extensionAttributes) {
		t.Fail()
	}

	if "name" != te.Name() {
		t.Fail()
	}

	if "documentation" != te.Documentation() {
		t.Fail()
	}

	if true != te.Asynchronous() {
		t.Fail()
	}

	if true != te.NotExclusive() {
		t.Fail()
	}

	if 1 != len(te.IncomingFlows()) {
		t.Fail()
	}

	if 1 != len(te.OutgoingFlows()) {
		t.Fail()
	}

	if 1 != len(te.EventDefinitions()) {
		t.Fail()
	}
}

// Verify throwEvent implements ThrowEvent
var _ ThrowEvent = (*throwEvent)(nil)

// Verify throwEvent implements Event
var _ Event = (*throwEvent)(nil)

// Verify throwEvent implements FlowNode
var _ FlowNode = (*throwEvent)(nil)

// Verify throwEvent implements FlowElement
var _ BaseElement = (*throwEvent)(nil)

// Verify throwEvent implements BaseElement
var _ BaseElement = (*throwEvent)(nil)
