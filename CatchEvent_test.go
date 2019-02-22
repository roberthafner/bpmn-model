package model

import "testing"

func TestCatchEvent_Methods(t *testing.T) {

	ce := catchEvent{
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
		},
	}

	if "123" != ce.Id() {
		t.Fail()
	}

	if 30 != ce.xmlRowNumber {
		t.Fail()
	}

	if 20 != ce.xmlColumnNumber {
		t.Fail()
	}

	if 1 != len(ce.extensionElements) {
		t.Fail()
	}

	if 1 != len(ce.extensionAttributes) {
		t.Fail()
	}

	if "name" != ce.Name() {
		t.Fail()
	}

	if "documentation" != ce.Documentation() {
		t.Fail()
	}

	if true != ce.Asynchronous() {
		t.Fail()
	}

	if true != ce.NotExclusive() {
		t.Fail()
	}

	if 1 != len(ce.IncomingFlows()) {
		t.Fail()
	}

	if 1 != len(ce.OutgoingFlows()) {
		t.Fail()
	}

	if 1 != len(ce.EventDefinitions()) {
		t.Fail()
	}
}

// Verify catchEvent implements CatchEvent
var _ CatchEvent = (*catchEvent)(nil)

// Verify catchEvent implements Event
var _ Event = (*catchEvent)(nil)

// Verify catchEvent implements FlowNode
var _ FlowNode = (*catchEvent)(nil)

// Verify catchEvent implements FlowElement
var _ BaseElement = (*catchEvent)(nil)

// Verify catchEvent implements BaseElement
var _ BaseElement = (*catchEvent)(nil)
