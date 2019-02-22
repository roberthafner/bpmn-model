package model

import "testing"

func TestStartEvent_Methods(t *testing.T) {

	se := StartEvent{
		catchEvent: catchEvent{
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
		},
	}

	if "123" != se.Id() {
		t.Fail()
	}

	if 30 != se.xmlRowNumber {
		t.Fail()
	}

	if 20 != se.xmlColumnNumber {
		t.Fail()
	}

	if 1 != len(se.extensionElements) {
		t.Fail()
	}

	if 1 != len(se.extensionAttributes) {
		t.Fail()
	}

	if "name" != se.Name() {
		t.Fail()
	}

	if "documentation" != se.Documentation() {
		t.Fail()
	}

	if true != se.Asynchronous() {
		t.Fail()
	}

	if true != se.NotExclusive() {
		t.Fail()
	}

	if 1 != len(se.IncomingFlows()) {
		t.Fail()
	}

	if 1 != len(se.OutgoingFlows()) {
		t.Fail()
	}

	if 1 != len(se.EventDefinitions()) {
		t.Fail()
	}
}


// Verify StartEvent implements CatchEvent
var _ CatchEvent = (*StartEvent)(nil)

// Verify StartEvent implements Event
var _ Event = (*StartEvent)(nil)

// Verify StartEvent implements FlowNode
var _ FlowNode = (*StartEvent)(nil)

// Verify StartEvent implements FlowElement
var _ BaseElement = (*StartEvent)(nil)

// Verify StartEvent implements BaseElement
var _ BaseElement = (*StartEvent)(nil)