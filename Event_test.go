package model

import "testing"

func TestEvent_Methods(t *testing.T) {

	e := event{
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
		eventDefinitions: []EventDefinition {eventDefinition{}},
	}

	if "123" != e.Id() {
		t.Fail()
	}

	if 30 != e.xmlRowNumber {
		t.Fail()
	}

	if 20 != e.xmlColumnNumber {
		t.Fail()
	}

	if 1 != len(e.extensionElements) {
		t.Fail()
	}

	if 1 != len(e.extensionAttributes) {
		t.Fail()
	}

	if "name" != e.Name() {
		t.Fail()
	}

	if "documentation" != e.Documentation() {
		t.Fail()
	}

	if true != e.Asynchronous() {
		t.Fail()
	}

	if true != e.NotExclusive() {
		t.Fail()
	}

	if 1 != len(e.IncomingFlows()) {
		t.Fail()
	}

	if 1 != len(e.OutgoingFlows()) {
		t.Fail()
	}

	if 1 != len(e.EventDefinitions()) {
		t.Fail()
	}
}

// Verify event implements Event
var _ Event = (*event)(nil)

// Verify event implements FlowNode
var _ FlowNode = (*event)(nil)

// Verify event implements FlowElement
var _ FlowElement = (*event)(nil)

// Verify event implements BaseElement
var _ BaseElement = (*event)(nil)
