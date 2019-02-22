package model

import "testing"

func TestFlowNode_Methods(t *testing.T) {

	fn := flowNode{
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
	}

	if "123" != fn.Id() {
		t.Fail()
	}

	if 30 != fn.xmlRowNumber {
		t.Fail()
	}

	if 20 != fn.xmlColumnNumber {
		t.Fail()
	}

	if 1 != len(fn.extensionElements) {
		t.Fail()
	}

	if 1 != len(fn.extensionAttributes) {
		t.Fail()
	}

	if "name" != fn.Name() {
		t.Fail()
	}

	if "documentation" != fn.Documentation() {
		t.Fail()
	}

	if true != fn.Asynchronous() {
		t.Fail()
	}

	if true != fn.NotExclusive() {
		t.Fail()
	}

	if 1 != len(fn.IncomingFlows()) {
		t.Fail()
	}

	if 1 != len(fn.OutgoingFlows()) {
		t.Fail()
	}
}

// Verify flowNode implements FlowNode
var _ FlowNode = (*flowNode)(nil)

// Verify flowNode implements FlowElement
var _ BaseElement = (*flowNode)(nil)

// Verify flowNode implements BaseElement
var _ BaseElement = (*flowNode)(nil)
