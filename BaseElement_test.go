package model

import "testing"

func TestBaseElement_Methods(t *testing.T) {

	be := baseElement{
		id:                  "123",
		xmlRowNumber:        30,
		xmlColumnNumber:     20,
		extensionElements:   map[string][]ExtensionElement{"ee": []ExtensionElement{ExtensionElement{}}},
		extensionAttributes: map[string][]ExtensionAttribute{"ea": []ExtensionAttribute{ExtensionAttribute{}}},
	}

	if "123" != be.Id() {
		t.Fail()
	}

	if 30 != be.xmlRowNumber {
		t.Fail()
	}

	if 20 != be.xmlColumnNumber {
		t.Fail()
	}

	if 1 != len(be.extensionElements) {
		t.Fail()
	}

	if 1 != len(be.extensionAttributes) {
		t.Fail()
	}
}

// Verify baseElement implements BaseElement
var _ BaseElement = (*baseElement)(nil)
