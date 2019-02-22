package model


import "testing"

func TestExtensionElement(t *testing.T) {

	ee := ExtensionElement{
		baseElement : baseElement {
			id:                  "123",
			xmlRowNumber:        30,
			xmlColumnNumber:     20,
			extensionElements:   map[string][]ExtensionElement{"ee": []ExtensionElement{ExtensionElement{}}},
			extensionAttributes: map[string][]ExtensionAttribute{"ea": []ExtensionAttribute{ExtensionAttribute{}}},
		},
		Name: "Name",
		NamespacePrefix: "NamespacePrefix",
		Namespace: "Namespace",
		ElementText: "ElementText",
		ChildElements : []ExtensionElement{},
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
}

// Verify baseElement implements BaseElement
var _ BaseElement = (*baseElement)(nil)