package model

import "testing"

func TestFlowElement_Methods(t *testing.T) {

	fe := flowElement{
		baseElement: baseElement{
			id:                  "123",
			xmlRowNumber:        30,
			xmlColumnNumber:     20,
			extensionElements:   map[string][]ExtensionElement{"ee": []ExtensionElement{ExtensionElement{}}},
			extensionAttributes: map[string][]ExtensionAttribute{"ea": []ExtensionAttribute{ExtensionAttribute{}}},
		},
		name:          "name",
		documentation: "documentation",
	}

	if "123" != fe.Id() {
		t.Fail()
	}

	if 30 != fe.xmlRowNumber {
		t.Fail()
	}

	if 20 != fe.xmlColumnNumber {
		t.Fail()
	}

	if 1 != len(fe.extensionElements) {
		t.Fail()
	}

	if 1 != len(fe.extensionAttributes) {
		t.Fail()
	}

	if "name" != fe.Name() {
		t.Fail()
	}

	if "documentation" != fe.Documentation() {
		t.Fail()
	}
}

// Verify flowElement implements FlowElement
var _ FlowElement = (*flowElement)(nil)

// Verify flowElement implements BaseElement
var _ BaseElement = (*flowElement)(nil)
