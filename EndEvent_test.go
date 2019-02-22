package model

import "testing"

func TestEndEvent_Methods(t *testing.T) {

	ee := EndEvent{
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

	if "123" != ee.Id() {
		t.Fail()
	}

	if 30 != ee.xmlRowNumber {
		t.Fail()
	}

	if 20 != ee.xmlColumnNumber {
		t.Fail()
	}

	if 1 != len(ee.extensionElements) {
		t.Fail()
	}

	if 1 != len(ee.extensionAttributes) {
		t.Fail()
	}

	if "name" != ee.Name() {
		t.Fail()
	}

	if "documentation" != ee.Documentation() {
		t.Fail()
	}

	if true != ee.Asynchronous() {
		t.Fail()
	}

	if true != ee.NotExclusive() {
		t.Fail()
	}

	if 1 != len(ee.IncomingFlows()) {
		t.Fail()
	}

	if 1 != len(ee.OutgoingFlows()) {
		t.Fail()
	}

	if 1 != len(ee.EventDefinitions()) {
		t.Fail()
	}
}


// Verify EndEvent implements CatchEvent
var _ CatchEvent = (*EndEvent)(nil)

// Verify EndEvent implements Event
var _ Event = (*EndEvent)(nil)

// Verify EndEvent implements FlowNode
var _ FlowNode = (*EndEvent)(nil)

// Verify EndEvent implements FlowElement
var _ BaseElement = (*EndEvent)(nil)

// Verify EndEvent implements BaseElement
var _ BaseElement = (*EndEvent)(nil)
