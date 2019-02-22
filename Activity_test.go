package model

import "testing"

func TestActivity_Methods(t *testing.T) {

	a := activity{
		flowNode: flowNode {
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
	}

	if "123" != a.Id() {
		t.Fail()
	}

	if 30 != a.xmlRowNumber {
		t.Fail()
	}

	if 20 != a.xmlColumnNumber {
		t.Fail()
	}

	if 1 != len(a.extensionElements) {
		t.Fail()
	}

	if 1 != len(a.extensionAttributes) {
		t.Fail()
	}
}

// Verify activity implements Activity
var _ Activity = (*activity)(nil)

// Verify activity implements FlowNode
var _ FlowNode = (*activity)(nil)

// Verify activity implements FlowElement
var _ FlowElement = (*activity)(nil)

// Verify activity implements BaseElement
var _ BaseElement = (*activity)(nil)
